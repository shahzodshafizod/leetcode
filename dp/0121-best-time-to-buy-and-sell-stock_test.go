package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxProfit$
func TestMaxProfit(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		max    int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, max: 5},
		{prices: []int{7, 6, 4, 3, 1}, max: 0},
		{prices: []int{2, 4, 1}, max: 2},
	} {
		max := maxProfit(tc.prices)
		assert.Equal(t, tc.max, max)
	}
}
