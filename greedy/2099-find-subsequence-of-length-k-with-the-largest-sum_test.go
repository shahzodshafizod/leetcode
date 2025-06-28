package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxSubsequence$
func TestMaxSubsequence(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		subseq []int
	}{
		{nums: []int{2, 1, 3, 3}, k: 2, subseq: []int{3, 3}},
		{nums: []int{-1, -2, 3, 4}, k: 3, subseq: []int{-1, 3, 4}},
		{nums: []int{3, 4, 3, 3}, k: 2, subseq: []int{3, 4}},
	} {
		subseq := maxSubsequence(tc.nums, tc.k)
		assert.Equal(t, tc.subseq, subseq)
	}
}
