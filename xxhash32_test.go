package xxhash

import (
	"fmt"
	"testing"
)

func TestXXHash32(t *testing.T) {
	testCases := []struct {
		sequence []byte
		seed     uint32
		result   uint32
	}{
		{nil, 0, 0x02CC5D05},
		{nil, prime, 0x36B78AE7},
		{sanityBuf[:1], 0, 0xB85CBEE5},
		{sanityBuf[:1], prime, 0xD5845D64},
		{sanityBuf[:14], 0, 0xE5AA0AB4},
		{sanityBuf[:14], prime, 0x4481951D},
		{sanityBuf, 0, 0x1F1AA412},
		{sanityBuf, prime, 0x498EC8E2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("len=%v,seed=%v", len(tc.sequence), tc.seed), func(t *testing.T) {
			h := New32(tc.seed)
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

func BenchmarkXXHash32(b *testing.B) {
	h := New32(prime)
	for i := 0; i < b.N; i++ {
		h.Write(sanityBuf)
		h.Sum32()
	}
}
