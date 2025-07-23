package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMaximumWealth$
func TestMaximumWealth(t *testing.T) {
	for _, tc := range []struct {
		accounts [][]int
		wealth   int
	}{
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, wealth: 6},
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, wealth: 10},
		{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, wealth: 17},
	} {
		wealth := maximumWealth(tc.accounts)
		assert.Equal(t, tc.wealth, wealth)
	}
}
