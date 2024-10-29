package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestStoneGameIII$
func TestStoneGameIII(t *testing.T) {
	for _, tc := range []struct {
		stoneValue []int
		winner     string
	}{
		{stoneValue: []int{1, 2, 3, 7}, winner: "Bob"},
		{stoneValue: []int{1, 2, 3, -9}, winner: "Alice"},
		{stoneValue: []int{1, 2, 3, 6}, winner: "Tie"},
	} {
		winner := stoneGameIII(tc.stoneValue)
		assert.Equal(t, tc.winner, winner)
	}
}
