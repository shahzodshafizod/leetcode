package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestProfitableSchemes$
func TestProfitableSchemes(t *testing.T) {
	for _, tc := range []struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		count     int
	}{
		{n: 5, minProfit: 3, group: []int{2, 2}, profit: []int{2, 3}, count: 2},
		{n: 10, minProfit: 5, group: []int{2, 3, 5}, profit: []int{6, 7, 8}, count: 7},
		{n: 64, minProfit: 0, group: []int{80, 40}, profit: []int{88, 88}, count: 2},
		// {n: 100, minProfit: 100, group: []int{24, 23, 7, 4, 26, 3, 7, 11, 1, 7, 1, 3, 5, 26, 26, 1, 13, 12, 2, 1, 7, 4, 1, 27, 13, 16, 26, 18, 6, 1, 1, 7, 16, 1, 6, 2, 5, 9, 19, 28, 1, 23, 2, 1, 3, 4, 4, 3, 22, 1, 1, 3, 5, 34, 2, 1, 22, 16, 8, 5, 3, 21, 1, 8, 14, 2, 1, 3, 8, 12, 40, 6, 4, 2, 2, 14, 1, 11, 9, 1, 7, 1, 1, 1, 6, 6, 4, 1, 1, 7, 8, 10, 20, 2, 14, 31, 1, 13, 1, 9}, profit: []int{5, 2, 38, 25, 4, 17, 5, 1, 4, 0, 0, 8, 13, 0, 20, 0, 28, 1, 22, 7, 10, 32, 6, 37, 0, 11, 6, 11, 23, 20, 13, 13, 6, 2, 36, 1, 0, 9, 4, 5, 6, 14, 20, 1, 13, 6, 33, 0, 22, 1, 17, 12, 10, 1, 19, 13, 8, 1, 0, 17, 20, 9, 8, 6, 2, 2, 1, 4, 22, 11, 3, 2, 6, 0, 40, 0, 0, 7, 1, 0, 25, 5, 12, 7, 19, 4, 12, 7, 4, 4, 1, 15, 33, 14, 2, 1, 1, 61, 4, 5}, count: 653387801},
	} {
		count := profitableSchemes(tc.n, tc.minProfit, tc.group, tc.profit)
		assert.Equal(t, tc.count, count)
	}
}
