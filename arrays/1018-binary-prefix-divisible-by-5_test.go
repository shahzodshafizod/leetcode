package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestPrefixesDivBy5$
func TestPrefixesDivBy5(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output []bool
	}{
		{nums: []int{0, 1, 1}, output: []bool{true, false, false}},
		{nums: []int{1, 1, 1}, output: []bool{false, false, false}},
		{nums: []int{0, 1, 1, 1, 1, 1}, output: []bool{true, false, false, false, true, false}},
	} {
		output := prefixesDivBy5(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
