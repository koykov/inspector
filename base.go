package inspector

import "github.com/koykov/fastconv"

// BaseInspector describes base struct.
type BaseInspector struct{}

func (b BaseInspector) Bufferize(buf, p []byte) ([]byte, []byte) {
	off := len(buf)
	buf = append(buf, p...)
	return buf, buf[off:]
}

func (b BaseInspector) BufferizeString(buf []byte, s string) ([]byte, string) {
	off := len(buf)
	buf = append(buf, s...)
	return buf, fastconv.B2S(buf[off:])
}
