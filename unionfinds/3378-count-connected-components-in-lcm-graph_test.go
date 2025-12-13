package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestCountComponents$
func TestCountComponents(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		threshold int
		output    int
	}{
		{nums: []int{2, 4, 8, 3, 9}, threshold: 5, output: 4},
		{nums: []int{2, 4, 8, 3, 9, 12}, threshold: 10, output: 2},
		{nums: []int{1}, threshold: 1, output: 1},
		{nums: []int{5, 10, 15}, threshold: 20, output: 1},
	} {
		output := countComponents(tc.nums, tc.threshold)
		assert.Equal(t, tc.output, output)
	}
}
