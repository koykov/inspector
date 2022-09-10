package inspector

// AccumulativeBuffer describes buffer that accumulates bytes data.
// Collects data during inspector functions work.
type AccumulativeBuffer interface {
	// AcquireBytes returns more space to use.
	AcquireBytes() []byte
	// ReleaseBytes returns space to the buffer.
	ReleaseBytes([]byte)
	// Reset all accumulated data.
	Reset()
}

type ByteBuffer struct {
	b []byte
}

func (ab ByteBuffer) AcquireBytes() []byte {
	return ab.b
}

func (ab *ByteBuffer) ReleaseBytes(p []byte) {
	if len(p) == 0 {
		return
	}
	ab.b = p
}

func (ab *ByteBuffer) Reset() {
	ab.b = ab.b[:0]
}
