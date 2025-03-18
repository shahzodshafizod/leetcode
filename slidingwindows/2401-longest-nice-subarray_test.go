package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestLongestNiceSubarray$
func TestLongestNiceSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{1, 3, 8, 48, 10}, length: 3},
		{nums: []int{3, 1, 5, 11, 13}, length: 1},
	} {
		length := longestNiceSubarray(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
