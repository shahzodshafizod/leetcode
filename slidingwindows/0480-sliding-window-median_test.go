package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMedianSlidingWindow$
func TestMedianSlidingWindow(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		k       int
		medians []float64
	}{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, medians: []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000}},
		{nums: []int{1, 2, 3, 4, 2, 3, 1, 4, 2}, k: 3, medians: []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}},
	} {
		medians := medianSlidingWindow(tc.nums, tc.k)
		assert.Equal(t, tc.medians, medians)
	}
}
