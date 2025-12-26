package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxProfitIV$
func TestMaxProfitIV(t *testing.T) {
	for _, tc := range []struct {
		k      int
		prices []int
		profit int
	}{
		{k: 2, prices: []int{2, 4, 1}, profit: 2},
		{k: 2, prices: []int{3, 2, 6, 5, 0, 3}, profit: 7},
		{k: 2, prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, profit: 6},
	} {
		profit := maxProfitIV(tc.k, tc.prices)
		assert.Equal(t, tc.profit, profit)
	}
}
