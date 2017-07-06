package xxhash

import (
	"encoding/binary"
	"io"
)

func read64(r io.Reader) (v uint64) {
	binary.Read(r, binary.LittleEndian, &v)
	return
}

func read32(r io.Reader) (v uint32) {
	binary.Read(r, binary.LittleEndian, &v)
	return
}

func read8(r io.Reader) (v uint8) {
	binary.Read(r, binary.LittleEndian, &v)
	return
}
