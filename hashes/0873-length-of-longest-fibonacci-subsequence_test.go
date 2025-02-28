package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestLenLongestFibSubseq$
func TestLenLongestFibSubseq(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		length int
	}{
		{arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, length: 5},
		{arr: []int{1, 3, 7, 11, 12, 14, 18}, length: 3},
	} {
		length := lenLongestFibSubseq(tc.arr)
		assert.Equal(t, tc.length, length)
	}
}
