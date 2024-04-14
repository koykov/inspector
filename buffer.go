package inspector

import "github.com/koykov/byteconv"

// AccumulativeBuffer describes buffer that accumulates bytes data.
// Collects data during inspector functions work.
type AccumulativeBuffer interface {
	// AcquireBytes returns more space to use.
	AcquireBytes() []byte
	// ReleaseBytes returns space to the buffer.
	ReleaseBytes([]byte)
	// Bufferize makes a copy of p to buffer and returns pointer to copy.
	Bufferize(p []byte) []byte
	// BufferizeString makes a copy of s to buffer and returns pointer to copy.
	BufferizeString(s string) string
	// Reset all accumulated data.
	Reset()
}

type ByteBuffer struct {
	b []byte
}

func NewByteBuffer(size int) *ByteBuffer {
	b := ByteBuffer{}
	if size > 0 {
		b.b = make([]byte, 0, size)
	}
	return &b
}

func (b *ByteBuffer) AcquireBytes() []byte {
	return b.b
}

func (b *ByteBuffer) ReleaseBytes(p []byte) {
	if len(p) == 0 {
		return
	}
	b.b = p
}

func (b *ByteBuffer) Bufferize(p []byte) []byte {
	off := len(b.b)
	b.b = append(b.b, p...)
	return b.b[off:]
}

func (b *ByteBuffer) BufferizeString(s string) string {
	off := len(b.b)
	b.b = append(b.b, s...)
	return byteconv.B2S(b.b[off:])
}

func (b *ByteBuffer) Reset() {
	b.b = b.b[:0]
}
