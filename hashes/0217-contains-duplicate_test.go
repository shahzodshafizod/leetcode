package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestContainsDuplicate$
func TestContainsDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		contains bool
	}{
		{nums: []int{1, 2, 3, 1}, contains: true},
		{nums: []int{1, 2, 3, 4}, contains: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, contains: true},
		{nums: []int{3, 3}, contains: true},
	} {
		contains := containsDuplicate(tc.nums)
		assert.Equal(t, tc.contains, contains)
	}
}
