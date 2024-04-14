package inspector

import "github.com/koykov/byteconv"

// Bufferize appends p to the buffer and returns both pointers to buffer and buffered data.
func Bufferize(buf, p []byte) ([]byte, []byte) {
	off := len(buf)
	buf = append(buf, p...)
	return buf, buf[off:]
}

// BufferizeString appends s to the buffer and returns both pointers to buffer and buffered data.
func BufferizeString(buf []byte, s string) ([]byte, string) {
	off := len(buf)
	buf = append(buf, s...)
	return buf, byteconv.B2S(buf[off:])
}
