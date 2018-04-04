package xxhash

import (
	"fmt"
	"testing"
)

func TestXXHash64(t *testing.T) {
	sanityBuf := newSanityBuf(t)
	testCases := []struct {
		sequence []byte
		seed     uint64
		result   uint64
	}{
		{nil, 0, 0xEF46DB3751D8E999},
		{nil, sanityPrime, 0xAC75FDA2929B17EF},
		{sanityBuf[:1], 0, 0x4FCE394CC88952D8},
		{sanityBuf[:1], sanityPrime, 0x739840CB819FA723},
		{sanityBuf[:14], 0, 0xCFFA8DB881BC3A3D},
		{sanityBuf[:14], sanityPrime, 0x5B9611585EFCC9CB},
		{sanityBuf, 0, 0x0EAB543384F878AD},
		{sanityBuf, sanityPrime, 0xCAA65939306F1E21},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("len=%v,seed=%v", len(tc.sequence), tc.seed), func(t *testing.T) {
			h := New64(tc.seed)
			n, err := h.Write(tc.sequence)
			if err != nil {
				t.Error("write:", err)
			}
			if n != len(tc.sequence) {
				t.Errorf("wrote=%v, want=%v", n, len(tc.sequence))
			}
			sum := h.Sum64()
			if sum != tc.result {
				t.Errorf("got=%x, want=%x", sum, tc.result)
			}
		})
	}
}

func BenchmarkXXHash64(b *testing.B) {
	sanityBuf := newSanityBuf(b)
	h := New64(sanityPrime)
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(sanityBuf)
		h.Sum64()
	}
}
