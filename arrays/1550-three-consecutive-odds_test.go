package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestThreeConsecutiveOdds$
func TestThreeConsecutiveOdds(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		exists bool
	}{
		{arr: []int{2, 6, 4, 1}, exists: false},
		{arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}, exists: true},
		{arr: []int{1, 3, 4, 6, 7, 45}, exists: false},
	} {
		exists := threeConsecutiveOdds(tc.arr)
		assert.Equal(t, tc.exists, exists)
	}
}
