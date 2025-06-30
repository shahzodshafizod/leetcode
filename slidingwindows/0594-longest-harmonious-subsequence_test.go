package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestFindLHS$
func TestFindLHS(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{1, 3, 2, 2, 5, 2, 3, 7}, length: 5},
		{nums: []int{1, 2, 3, 4}, length: 2},
		{nums: []int{1, 1, 1, 1}, length: 0},
	} {
		length := findLHS(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
