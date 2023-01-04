package xxhash

import (
	"bytes"
	"hash"
)

const (
	prime32_1 = 0x9e3779b1
	prime32_2 = 0x85ebca77
	prime32_3 = 0xc2b2ae3d
	prime32_4 = 0x27d4eb2f
	prime32_5 = 0x165667b1
)

type digest32 struct {
	seed, v1, v2, v3, v4 uint32
	buf                  *bytes.Buffer
	totalLen             int
	rbuf                 [4]byte
}

func (d *digest32) read32() (v uint32) {
	d.buf.Read(d.rbuf[:4])
	for i := 0; i < 4; i++ {
		v |= uint32(d.rbuf[i]) << (8 * uint(i))
	}
	return

}

func (d *digest32) read8() uint32 {
	d.buf.Read(d.rbuf[:1])
	return uint32(d.rbuf[0])
}

// XXH32 returns new hash.Hash32
func XXH32(seed uint32) hash.Hash32 {
	d := &digest32{seed: seed, buf: new(bytes.Buffer)}
	d.Reset()
	return d
}

func (d *digest32) Write(p []byte) (n int, err error) {
	n, err = d.buf.Write(p)
	if err != nil {
		return 0, err
	}
	d.totalLen += n
	for d.buf.Len() >= 16 {
		d.v1 = round32(d.v1, d.read32())
		d.v2 = round32(d.v2, d.read32())
		d.v3 = round32(d.v3, d.read32())
		d.v4 = round32(d.v4, d.read32())
	}
	return n, nil
}

func (d *digest32) Sum(b []byte) []byte {
	v := d.Sum32()
	for i := d.Size() - 1; i >= 0; i-- {
		b = append(b, byte(v>>uint(8*i)))
	}
	return b
}

func (d *digest32) Reset() {
	d.v1 = d.seed + prime32_1 + prime32_2
	d.v2 = d.seed + prime32_2
	d.v3 = d.seed
	d.v4 = d.seed - prime32_1
	d.buf.Reset()
	d.totalLen = 0
}

func (d *digest32) Size() int {
	return 4
}

func (d *digest32) BlockSize() int {
	return 16
}

func (d *digest32) Sum32() uint32 {
	var sum uint32
	if d.totalLen >= 16 {
		sum = rotl32(d.v1, 1) + rotl32(d.v2, 7) + rotl32(d.v3, 12) + rotl32(d.v4, 18)
	} else {
		sum = d.v3 + prime32_5
	}

	sum += uint32(d.totalLen)

	for d.buf.Len() >= 4 {
		sum += d.read32() * prime32_3
		sum = rotl32(sum, 17) * prime32_4
	}

	for d.buf.Len() >= 1 {
		sum += d.read8() * prime32_5
		sum = rotl32(sum, 11) * prime32_1
	}

	sum ^= sum >> 15
	sum *= prime32_2
	sum ^= sum >> 13
	sum *= prime32_3
	sum ^= sum >> 16

	return sum
}

func round32(acc, val uint32) uint32 {
	acc += val * prime32_2
	return rotl32(acc, 13) * prime32_1
}

func rotl32(x uint32, r uint) uint32 {
	return (x << r) | (x >> (32 - r))
}
