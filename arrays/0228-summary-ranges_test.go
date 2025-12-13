package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSummaryRanges$
func TestSummaryRanges(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output []string
	}{
		{nums: []int{0, 1, 2, 4, 5, 7}, output: []string{"0->2", "4->5", "7"}},
		{nums: []int{0, 2, 3, 4, 6, 8, 9}, output: []string{"0", "2->4", "6", "8->9"}},
		{nums: []int{}, output: []string{}},
		{nums: []int{1}, output: []string{"1"}},
	} {
		output := summaryRanges(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
