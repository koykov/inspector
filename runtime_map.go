package inspector

import "unsafe"

const (
	// Maximum number of key/value pairs a bucket can hold.
	BucketCntBits = 3
	BucketCnt     = 1 << BucketCntBits

	EmptyRest  = 0 // this cell is empty, and there are no more non-empty cells at higher indexes or overflows.
	EmptyOne   = 1 // this cell is empty
	MinTopHash = 5 // minimum tophash for a normal filled cell.

	HashWriting  = 4 // a goroutine is writing to the map
	SameSizeGrow = 8 // the current map growth is to a new map of the same size

	// data offset should be the size of the bmap struct, but needs to be
	// aligned correctly. For amd64p32 this means 64-bit alignment
	// even though pointers are 32 bit.
	DataOffset = unsafe.Offsetof(struct {
		b Bmap
		v int64
	}{}.v)

	PtrSize = 4 << (^uintptr(0) >> 63)
)

// A header for a Go map.
type Hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	Count     int // # live cells == size of map.  Must be first (used by len() builtin)
	Flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	Noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	Hash0     uint32 // hash seed

	Buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	Oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	Nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	Extra *Mapextra // optional fields
}

// mapextra holds fields that are not present on all maps.
type Mapextra struct {
	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and value do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	Overflow    *[]*Bmap
	Oldoverflow *[]*Bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	NextOverflow *Bmap
}

// A bucket for a Go map.
type Bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	Tophash [BucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

func Mapaccess2(t *Maptype, h *Hmap, key unsafe.Pointer) (unsafe.Pointer, bool) {
	// if raceenabled && h != nil {
	// 	callerpc := getcallerpc()
	// 	pc := funcPC(mapaccess2)
	// 	racereadpc(unsafe.Pointer(h), callerpc, pc)
	// 	raceReadObjectPC(t.key, key, callerpc, pc)
	// }
	// if msanenabled && h != nil {
	// 	msanread(key, t.key.size)
	// }
	if h == nil || h.Count == 0 {
		if t.HashMightPanic() {
			t.Key.Alg.Hash(key, 0) // see issue 23734
		}
		return unsafe.Pointer(&zeroVal[0]), false
	}
	if h.Flags&HashWriting != 0 {
		// Replaces panic with false result.
		return unsafe.Pointer(&zeroVal[0]), false
	}
	alg := t.Key.Alg
	hash := alg.Hash(key, uintptr(h.Hash0))
	m := BucketMask(h.B)
	b := (*Bmap)(unsafe.Pointer(uintptr(h.Buckets) + (hash&m)*uintptr(t.Bucketsize)))
	if c := h.Oldbuckets; c != nil {
		if !h.SameSizeGrow() {
			// There used to be half as many buckets; mask down one more power of two.
			m >>= 1
		}
		oldb := (*Bmap)(unsafe.Pointer(uintptr(c) + (hash&m)*uintptr(t.Bucketsize)))
		if !Evacuated(oldb) {
			b = oldb
		}
	}
	top := Tophash(hash)
bucketloop:
	for ; b != nil; b = b.Overflow(t) {
		for i := uintptr(0); i < BucketCnt; i++ {
			if b.Tophash[i] != top {
				if b.Tophash[i] == EmptyRest {
					break bucketloop
				}
				continue
			}
			k := Add(unsafe.Pointer(b), DataOffset+i*uintptr(t.Keysize))
			if t.Indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
			if alg.Equal(key, k) {
				v := Add(unsafe.Pointer(b), DataOffset+BucketCnt*uintptr(t.Keysize)+i*uintptr(t.Valuesize))
				if t.Indirectvalue() {
					v = *((*unsafe.Pointer)(v))
				}
				return v, true
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0]), false
}

// bucketShift returns 1<<b, optimized for code generation.
func BucketShift(b uint8) uintptr {
	// if sys.GoarchAmd64|sys.GoarchAmd64p32|sys.Goarch386 != 0 {
	b &= PtrSize*8 - 1 // help x86 archs remove shift overflow checks
	// }
	return uintptr(1) << b
}

// bucketMask returns 1<<b - 1, optimized for code generation.
func BucketMask(b uint8) uintptr {
	return BucketShift(b) - 1
}

// sameSizeGrow reports whether the current growth is to a map of the same size.
func (h *Hmap) SameSizeGrow() bool {
	return h.Flags&SameSizeGrow != 0
}

func Evacuated(b *Bmap) bool {
	h := b.Tophash[0]
	return h > EmptyOne && h < MinTopHash
}

func (b *Bmap) Overflow(t *Maptype) *Bmap {
	return *(**Bmap)(Add(unsafe.Pointer(b), uintptr(t.Bucketsize)-PtrSize))
}

// Tophash calculates the tophash value for hash.
func Tophash(hash uintptr) uint8 {
	top := uint8(hash >> (PtrSize*8 - 8))
	if top < MinTopHash {
		top += MinTopHash
	}
	return top
}

func Add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

const maxZero = 1024 // must match value in cmd/compile/internal/gc/walk.go
var zeroVal [maxZero]byte
