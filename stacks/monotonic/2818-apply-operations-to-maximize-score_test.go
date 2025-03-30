package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestMaximumScore$
func TestMaximumScore(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		score int
	}{
		{nums: []int{8, 3, 9, 3, 8}, k: 2, score: 81},
		{nums: []int{19, 12, 14, 6, 10, 18}, k: 3, score: 4788},
		{nums: []int{6, 8, 16, 14, 9}, k: 3, score: 3136},
		{nums: []int{1}, k: 1, score: 1},
		{nums: []int{1, 7, 11, 1, 5}, k: 14, score: 823751938},
		{nums: []int{3289, 2832, 14858, 22011}, k: 6, score: 256720975},
		{nums: []int{6, 1, 13, 10, 1, 17, 6}, k: 27, score: 630596200},
		{nums: []int{3289, 2832, 14858, 22011}, k: 6, score: 256720975},
		{nums: []int{12, 5, 1, 6, 9, 1, 17, 14}, k: 12, score: 62996359},
	} {
		score := maximumScore(tc.nums, tc.k)
		assert.Equal(t, tc.score, score)
	}
}
