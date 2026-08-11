package inspector

import (
	"github.com/koykov/byteconv"
	"github.com/koykov/x2bytes"
)

// AccumulativeBuffer describes buffer that accumulates bytes data.
// Collects data during inspector functions work.
type AccumulativeBuffer interface {
	// AcquireBytes returns more space to use.
	AcquireBytes() []byte
	// ReleaseBytes returns space to the buffer.
	ReleaseBytes([]byte)
	// Bufferize makes a copy of p to buffer and returns pointer to copy.
	Bufferize(p []byte) []byte
	// BufferizeString makes a copy of s to buffer and returns  result as a string.
	BufferizeString(s string) string
	// BufferizeAny makes a copy of x to buffer and returns pointer to copy.
	BufferizeAny(x any) ([]byte, error)
	// BufferizeAnyString makes a copy of x to buffer and returns result as a string.
	BufferizeAnyString(x any) (string, error)
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

func (b *ByteBuffer) BufferizeAny(x any) ([]byte, error) {
	bb, err := x2bytes.ToBytes(b.b, x)
	if err != nil {
		return nil, err
	}
	return bb, nil
}

func (b *ByteBuffer) BufferizeAnyString(x any) (string, error) {
	bb, err := x2bytes.ToBytes(b.b, x)
	if err != nil {
		return "", err
	}
	return byteconv.B2S(bb), nil
}

func (b *ByteBuffer) Reset() {
	b.b = b.b[:0]
}
