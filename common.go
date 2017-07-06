package xxhash

import (
	"io"
)

func read64(r io.Reader) (v uint64) {
	var buf [8]byte
	r.Read(buf[:])
	for i := 0; i < len(buf); i++ {
		v |= uint64(buf[i]) << (8 * uint(i))
	}
	return v
}

func read32(r io.Reader) (v uint32) {
	var buf [4]byte
	r.Read(buf[:])
	for i := 0; i < len(buf); i++ {
		v |= uint32(buf[i]) << (8 * uint(i))
	}
	return v
}

func read8(r io.Reader) uint8 {
	var buf [1]byte
	r.Read(buf[:])
	return uint8(buf[0])
}
