package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestNumberOfArrays$
func TestNumberOfArrays(t *testing.T) {
	for _, tc := range []struct {
		differences []int
		lower       int
		upper       int
		count       int
	}{
		{differences: []int{1, -3, 4}, lower: 1, upper: 6, count: 2},
		{differences: []int{3, -4, 5, 1, -2}, lower: -4, upper: 5, count: 4},
		{differences: []int{4, -7, 2}, lower: 3, upper: 6, count: 0},
		{differences: []int{2, -2, 3, -3, 4, -4, 5, -5}, lower: 0, upper: 10, count: 6},
		{differences: []int{1, -1, 2, -2, 3, -3, 4, -4, 5, -5}, lower: -3, upper: 7, count: 6},
		{differences: []int{3, -2, 1, -3, 2, -1, 3, -3}, lower: 1, upper: 8, count: 4},
		{differences: []int{1, 1, 1, 1, 1, -5}, lower: 0, upper: 5, count: 1},
		{differences: []int{5, -4, 3, -2, 1, -1, -2, 2, -3, 3}, lower: -2, upper: 6, count: 3},
		{differences: []int{1, 2, -1, -2, 3, -3, 4, -4, 5, -5}, lower: 0, upper: 10, count: 6},
		{differences: []int{-40}, lower: -46, upper: 53, count: 60},
		{differences: []int{83702, -5216}, lower: -82788, upper: 14602, count: 13689},
	} {
		count := numberOfArrays(tc.differences, tc.lower, tc.upper)
		assert.Equal(t, tc.count, count)
	}
}
