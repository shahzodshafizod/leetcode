package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestBeautifulSubsets$
func TestBeautifulSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		k       int
		subsets int
	}{
		{nums: []int{2, 4, 6}, k: 2, subsets: 4},
		{nums: []int{1}, k: 1, subsets: 1},
		{nums: []int{1, 5, 9, 13, 17, 21, 25, 29}, k: 2, subsets: 255},
		{nums: []int{4, 2, 5, 9, 10, 3}, k: 1, subsets: 23},
	} {
		subsets := beautifulSubsets(tc.nums, tc.k)
		assert.Equal(t, tc.subsets, subsets)
	}
}
