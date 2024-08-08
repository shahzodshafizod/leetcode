package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums        []int
		increasings int
	}{
		{nums: []int{1, 1, 1}, increasings: 3},
		{nums: []int{1, 5, 2, 4, 1}, increasings: 14},
		{nums: []int{8}, increasings: 0},
		{nums: []int{7, 7, 10, 2, 9}, increasings: 13},
	} {
		increasings := minOperations(tc.nums)
		assert.Equal(t, tc.increasings, increasings)
	}
}
