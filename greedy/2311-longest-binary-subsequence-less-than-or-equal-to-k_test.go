package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestLongestSubsequence$
func TestLongestSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s      string
		k      int
		length int
	}{
		{s: "0", k: 1, length: 1},
		{s: "1001010", k: 5, length: 5},
		{s: "00101001", k: 1, length: 6},
		{s: "101011010", k: 15, length: 6},
		{s: "1111111111", k: 1, length: 1},
		{s: "0001101010", k: 5, length: 7},
		{s: "11001101000111010111", k: 15, length: 11},
		{s: "00101000001101011111111110110111101011010011100100", k: 5, length: 22},
	} {
		length := longestSubsequence(tc.s, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
