package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxProfitIII$
func TestMaxProfitIII(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		profit int
	}{
		{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, profit: 6},
		{prices: []int{1, 2, 3, 4, 5}, profit: 4},
		{prices: []int{7, 6, 4, 3, 1}, profit: 0},
		{prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, profit: 13},
		{prices: []int{3}, profit: 0},
		{prices: []int{2, 1, 2, 0, 1}, profit: 2},
		{prices: []int{14, 9, 10, 12, 4, 8, 1, 16}, profit: 19},
		{prices: []int{1, 2, 0, 4}, profit: 5},
		{prices: []int{0, 3, 1, 4}, profit: 6},
	} {
		profit := maxProfitIII(tc.prices)
		assert.Equal(t, tc.profit, profit)
	}
}
