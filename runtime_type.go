package inspector

import "unsafe"

// tflag is documented in reflect/type.go.
//
// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	reflect/type.go
type Tflag uint8

const (
	TflagUncommon  Tflag = 1 << 0
	TflagExtraStar Tflag = 1 << 1
	TflagNamed     Tflag = 1 << 2
)

type Maptype struct {
	Typ        Typ
	Key        *Typ
	Elem       *Typ
	Bucket     *Typ   // internal type representing a hash bucket
	Keysize    uint8  // size of key slot
	Valuesize  uint8  // size of value slot
	Bucketsize uint16 // size of bucket
	Flags      uint32
}

// Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
type Typ struct {
	Size       uintptr
	Ptrdata    uintptr // size of memory prefix holding all pointers
	Hash       uint32
	Tflag      Tflag
	Align      uint8
	Fieldalign uint8
	Kind       uint8
	Alg        *TypeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	Gcdata    *byte
	Str       NameOff
	PtrToThis TypeOff
}

// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.
type TypeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	Hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type NameOff int32
type TypeOff int32

func (mt *Maptype) HashMightPanic() bool { // true if hash function might panic
	return mt.Flags&16 != 0
}

// Note: flag values must match those used in the TMAP case
// in ../cmd/compile/internal/gc/reflect.go:dtypesym.
func (mt *Maptype) Indirectkey() bool { // store ptr to key instead of key itself
	return mt.Flags&1 != 0
}
func (mt *Maptype) Indirectvalue() bool { // store ptr to value instead of value itself
	return mt.Flags&2 != 0
}
