package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCoinChange$
func TestCoinChange(t *testing.T) {
	for _, tc := range []struct {
		coins  []int
		amount int
		min    int
	}{
		{coins: []int{1, 2, 5}, amount: 11, min: 3},
		{coins: []int{2}, amount: 3, min: -1},
		{coins: []int{1}, amount: 0, min: 0},
		{coins: []int{1, 6, 7, 9, 11}, amount: 13, min: 2},
		{coins: []int{5, 4, 3}, amount: 11, min: 3},
		{coins: []int{6, 5, 1}, amount: 20, min: 4},              // [4]*5
		{coins: []int{186, 419, 83, 408}, amount: 6249, min: 20}, // [5]*419+[8]*408+[3]*186+[4]*83
		{coins: []int{1, 3, 4, 5}, amount: 7, min: 2},
		{coins: []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, amount: 9864, min: 24},
	} {
		minimum := coinChange(tc.coins, tc.amount)
		assert.Equal(t, tc.min, minimum)
	}
}
