package xxhash

import (
	"fmt"
	"testing"
)

func TestXXH32(t *testing.T) {
	sanityBuf := newSanityBuf(t)
	testCases := []struct {
		sequence []byte
		seed     uint32
		result   uint32
	}{
		{nil, 0, 0x02CC5D05},
		{nil, sanityPrime, 0x36B78AE7},
		{sanityBuf[:1], 0, 0xB85CBEE5},
		{sanityBuf[:1], sanityPrime, 0xD5845D64},
		{sanityBuf[:14], 0, 0xE5AA0AB4},
		{sanityBuf[:14], sanityPrime, 0x4481951D},
		{sanityBuf, 0, 0x1F1AA412},
		{sanityBuf, sanityPrime, 0x498EC8E2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("len=%v,seed=%v", len(tc.sequence), tc.seed), func(t *testing.T) {
			h := XXH32(tc.seed)
			n, err := h.Write(tc.sequence)
			if err != nil {
				t.Error("write:", err)
			}
			if n != len(tc.sequence) {
				t.Errorf("wrote=%v, want=%v", n, len(tc.sequence))
			}
			sum := h.Sum32()
			if sum != tc.result {
				t.Errorf("got=%x, want=%x", sum, tc.result)
			}
		})
	}
}

func BenchmarkXXH32(b *testing.B) {
	sanityBuf := newSanityBuf(b)
	h := XXH32(sanityPrime)
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(sanityBuf)
		h.Sum32()
	}
}
