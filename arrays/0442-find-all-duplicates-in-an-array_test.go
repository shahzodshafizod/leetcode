package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindDuplicates$
func TestFindDuplicates(t *testing.T) {
	for _, tc := range []struct {
		nums       []int
		duplicates []int
	}{
		{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, duplicates: []int{2, 3}},
		{nums: []int{1, 1, 2}, duplicates: []int{1}},
		{nums: []int{1}, duplicates: []int{}},
		{nums: []int{2, 5, 2, 1, 1, 4}, duplicates: []int{2, 1}},
	} {
		duplicates := findDuplicates(tc.nums)
		assert.Equal(t, tc.duplicates, duplicates)
	}
}
