package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaxFrequencyElements$
func TestMaxFrequencyElements(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		totalMax int
	}{
		{nums: []int{1, 2, 2, 3, 1, 4}, totalMax: 4},
		{nums: []int{1, 2, 3, 4, 5}, totalMax: 5},
		{nums: []int{93}, totalMax: 1},
		{nums: []int{46, 64}, totalMax: 2},
		{nums: []int{61, 68, 75}, totalMax: 3},
	} {
		totalMax := maxFrequencyElements(tc.nums)
		assert.Equal(t, tc.totalMax, totalMax)
	}
}
