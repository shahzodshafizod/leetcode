package fenwicks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./fenwicks/ -run ^TestResultArray$
func TestResultArray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		result []int
	}{
		{nums: []int{2, 1, 3, 3}, result: []int{2, 3, 1, 3}},
		{nums: []int{5, 14, 3, 1, 2}, result: []int{5, 3, 1, 2, 14}},
		{nums: []int{3, 3, 3, 3}, result: []int{3, 3, 3, 3}},
	} {
		result := resultArray(tc.nums)
		assert.Equal(t, tc.result, result)
	}
}
