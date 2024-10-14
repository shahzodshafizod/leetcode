package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaximumScore$
func TestMaximumScore(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		score int
	}{
		{nums: []int{1, 4, 3, 7, 4, 5}, k: 3, score: 15},
		{nums: []int{5, 5, 4, 5, 4, 1, 1, 1}, k: 0, score: 20},
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2}, k: 28, score: 30},
	} {
		score := maximumScore(tc.nums, tc.k)
		assert.Equal(t, tc.score, score)
	}
}
