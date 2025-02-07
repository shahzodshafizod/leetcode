package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestTupleSameProduct$
func TestTupleSameProduct(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{2, 3, 4, 6}, count: 8},
		{nums: []int{1, 2, 4, 5, 10}, count: 16},
		{nums: []int{2, 3, 5, 7}, count: 0},
		{nums: []int{2, 3, 4, 6, 8, 12}, count: 40},
		{nums: []int{1, 2, 4, 8, 16, 32}, count: 56},
		{nums: []int{10, 5, 15, 8, 6, 12, 20, 4}, count: 72},
		{nums: []int{20, 10, 6, 24, 15, 5, 4, 30}, count: 48},
		{nums: []int{19, 32, 39, 38, 43, 47, 9, 5, 35, 22}, count: 0},
		{nums: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}, count: 1288},
	} {
		count := tupleSameProduct(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
