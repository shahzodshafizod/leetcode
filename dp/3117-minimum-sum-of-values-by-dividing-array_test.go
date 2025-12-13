package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinimumValueSum$
func TestMinimumValueSum(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		andValues []int
		output    int
	}{
		{nums: []int{1, 4, 3, 3, 2}, andValues: []int{0, 3, 3, 2}, output: 12},
		{nums: []int{2, 3, 5, 7, 7, 7, 5}, andValues: []int{0, 7, 5}, output: 17},
		{nums: []int{1, 2, 3, 4}, andValues: []int{2}, output: -1},
	} {
		output := minimumValueSum(tc.nums, tc.andValues)
		assert.Equal(t, tc.output, output)
	}
}
