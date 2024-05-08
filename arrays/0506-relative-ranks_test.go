package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindRelativeRanks$
func TestFindRelativeRanks(t *testing.T) {
	for _, tc := range []struct {
		score []int
		ranks []string
	}{
		{score: []int{5, 4, 3, 2, 1}, ranks: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{score: []int{10, 3, 8, 9, 4}, ranks: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	} {
		ranks := findRelativeRanks(tc.score)
		assert.Equal(t, tc.ranks, ranks)
	}
}
