package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinimumCost$
func TestMinimumCost(t *testing.T) {
	for _, tc := range []struct {
		target  string
		words   []string
		costs   []int
		mincost int
	}{
		{target: "abcdef", words: []string{"abdef", "abc", "d", "def", "ef"}, costs: []int{100, 1, 1, 10, 5}, mincost: 7},
		{target: "aaaa", words: []string{"z", "zz", "zzz"}, costs: []int{1, 10, 100}, mincost: -1},
	} {
		mincost := minimumCost(tc.target, tc.words, tc.costs)
		assert.Equal(t, tc.mincost, mincost)
	}
}
