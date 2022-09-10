package inspector

import "github.com/koykov/fastconv"

func Bufferize(buf, p []byte) ([]byte, []byte) {
	off := len(buf)
	buf = append(buf, p...)
	return buf, buf[off:]
}

func BufferizeString(buf []byte, s string) ([]byte, string) {
	off := len(buf)
	buf = append(buf, s...)
	return buf, fastconv.B2S(buf[off:])
}
