package xxhash

import (
	"bytes"
	"hash"
)

const (
	prime64_1 = 11400714785074694791
	prime64_2 = 14029467366897019727
	prime64_3 = 1609587929392839161
	prime64_4 = 9650029242287828579
	prime64_5 = 2870177450012600261
)

type digest64 struct {
	seed, v1, v2, v3, v4 uint64
	buf                  *bytes.Buffer
	totalLen             int
}

// New64 returns new hash.Hash64
func New64(seed uint64) hash.Hash64 {
	d := &digest64{seed: seed, buf: new(bytes.Buffer)}
	d.Reset()
	return d
}

func (d *digest64) Write(p []byte) (n int, err error) {
	n, err = d.buf.Write(p)
	if err != nil {
		return 0, err
	}
	d.totalLen += n
	d.update()
	return n, nil
}

func (d *digest64) Sum(b []byte) []byte {
	v := d.Sum64()
	for i := d.Size() - 1; i >= 0; i-- {
		b = append(b, byte(v>>uint(8*i)))
	}
	return b
}

func (d *digest64) Reset() {
	d.v1 = d.seed + prime64_1 + prime64_2
	d.v2 = d.seed + prime64_2
	d.v3 = d.seed
	d.v4 = d.seed - prime64_1
	d.buf.Reset()
	d.totalLen = 0
}

func (d *digest64) Size() int {
	return 8
}

func (d *digest64) BlockSize() int {
	return 1
}

func (d *digest64) update() {
	for d.buf.Len() > 32 {
		d.v1 = round64(d.v1, read64(d.buf))
		d.v2 = round64(d.v2, read64(d.buf))
		d.v3 = round64(d.v3, read64(d.buf))
		d.v4 = round64(d.v4, read64(d.buf))
	}
}

func (d *digest64) digest() (sum uint64) {
	if d.totalLen >= 32 {
		sum = rotl64(d.v1, 1) + rotl64(d.v2, 7) + rotl64(d.v3, 12) + rotl64(d.v4, 18)
		sum = merge64(sum, d.v1)
		sum = merge64(sum, d.v2)
		sum = merge64(sum, d.v3)
		sum = merge64(sum, d.v4)
	} else {
		sum = d.v3 + prime64_5
	}

	sum += uint64(d.totalLen)

	for d.buf.Len() >= 8 {
		sum ^= round64(0, read64(d.buf))
		sum = rotl64(sum, 27)*prime64_1 + prime64_4
	}

	if d.buf.Len() >= 4 {
		sum ^= uint64(read32(d.buf)) * prime64_1
		sum = rotl64(sum, 23)*prime64_2 + prime64_3
	}

	for d.buf.Len() >= 1 {
		sum ^= uint64(read8(d.buf)) * prime64_5
		sum = rotl64(sum, 11) * prime64_1
	}

	sum ^= sum >> 33
	sum *= prime64_2
	sum ^= sum >> 29
	sum *= prime64_3
	sum ^= sum >> 32

	return sum
}

func (d *digest64) Sum64() uint64 {
	d.update()
	return d.digest()
}

func round64(acc, val uint64) uint64 {
	acc += val * prime64_2
	acc = rotl64(acc, 31)
	acc *= prime64_1
	return acc
}

func merge64(acc, val uint64) uint64 {
	val = round64(0, val)
	acc ^= val
	acc = acc*prime64_1 + prime64_4
	return acc
}

func rotl64(x uint64, r uint) uint64 {
	return (x << r) | (x >> (64 - r))
}
