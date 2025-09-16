package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestReplaceNonCoprimes$
func TestReplaceNonCoprimes(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		final []int
	}{
		{nums: []int{6, 4, 3, 2, 7, 6, 2}, final: []int{12, 7, 6}},
		{nums: []int{2, 2, 1, 1, 3, 3, 3}, final: []int{2, 1, 1, 3}},
	} {
		final := replaceNonCoprimes(tc.nums)
		assert.Equal(t, tc.final, final)
	}
}
