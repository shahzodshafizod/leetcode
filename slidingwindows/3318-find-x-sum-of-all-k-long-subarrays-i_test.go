package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestFindXSum$
func TestFindXSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		x      int
		answer []int
	}{
		{nums: []int{1, 1, 2, 2, 3, 4, 2, 3}, k: 6, x: 2, answer: []int{6, 10, 12}},
		{nums: []int{3, 8, 7, 8, 7, 5}, k: 2, x: 2, answer: []int{11, 15, 15, 15, 12}},
	} {
		answer := findXSum(tc.nums, tc.k, tc.x)
		assert.Equal(t, tc.answer, answer)
	}
}
