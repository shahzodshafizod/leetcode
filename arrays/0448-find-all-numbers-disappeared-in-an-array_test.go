package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindDisappearedNumbers$
func TestFindDisappearedNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output []int
	}{
		{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, output: []int{5, 6}},
		{nums: []int{1, 1}, output: []int{2}},
		{nums: []int{1, 2, 3, 4, 5}, output: nil},
		{nums: []int{2, 2, 2}, output: []int{1, 3}},
	} {
		output := findDisappearedNumbers(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
