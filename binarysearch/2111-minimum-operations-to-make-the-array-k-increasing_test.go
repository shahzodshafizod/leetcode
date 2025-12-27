package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestKIncreasing$
func TestKIncreasing(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		k      int
		output int
	}{
		{arr: []int{5, 4, 3, 2, 1}, k: 1, output: 4},
		{arr: []int{4, 1, 5, 2, 6, 2}, k: 2, output: 0},
		{arr: []int{4, 1, 5, 2, 6, 2}, k: 3, output: 2},
		{arr: []int{12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3}, k: 1, output: 12},
	} {
		output := kIncreasing(tc.arr, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
