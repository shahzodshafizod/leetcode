package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestFindKthSmallest$
func TestFindKthSmallest(t *testing.T) {
	for _, tc := range []struct {
		coins  []int
		k      int
		output int64
	}{
		{coins: []int{3, 6, 9}, k: 3, output: 9},
		{coins: []int{5, 2}, k: 7, output: 12},
		{coins: []int{1}, k: 5, output: 5},
	} {
		output := findKthSmallest(tc.coins, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
