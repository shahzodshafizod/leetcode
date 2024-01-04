package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./internal/arrays/ -run ^TestTwoSum$
func TestTwoSum(t *testing.T) {
	for _, data := range []struct {
		nums   []int
		target int
		result []int
	}{
		{nums: []int{1, 3, 7, 9, 2}, target: 11, result: []int{3, 4}},
		{nums: []int{1, 3, 7, 9, 2}, target: 25, result: nil},
		{nums: []int{}, target: 1, result: nil},
		{nums: []int{5}, target: 5, result: nil},
		{nums: []int{1, 6}, target: 7, result: []int{0, 1}},
		{nums: nil, target: 1, result: nil},
		{nums: []int{2, 7, 11, 15}, target: 9, result: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, result: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, result: []int{0, 1}},
	} {
		result := twoSum(data.nums, data.target)
		assert.Equal(t, data.result, result)
	}
}
