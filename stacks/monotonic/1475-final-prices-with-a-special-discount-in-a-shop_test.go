package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestFinalPrices$
func TestFinalPrices(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		answer []int
	}{
		{prices: []int{10, 1, 1, 6}, answer: []int{9, 0, 1, 6}},
		{prices: []int{8, 4, 6, 2, 3}, answer: []int{4, 2, 4, 2, 3}},
		{prices: []int{1, 2, 3, 4, 5}, answer: []int{1, 2, 3, 4, 5}},
		{prices: []int{10, 2, 5, 2, 8}, answer: []int{8, 0, 3, 2, 8}},
		{prices: []int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}, answer: []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1}},
	} {
		answer := finalPrices(tc.prices)
		assert.Equal(t, tc.answer, answer)
	}
}
