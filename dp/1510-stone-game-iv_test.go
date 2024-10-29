package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestWinnerSquareGame$
func TestWinnerSquareGame(t *testing.T) {
	for _, tc := range []struct {
		n         int
		aliceWins bool
	}{
		{n: 1, aliceWins: true},
		{n: 2, aliceWins: false},
		{n: 4, aliceWins: true},
	} {
		aliceWins := winnerSquareGame(tc.n)
		assert.Equal(t, tc.aliceWins, aliceWins)
	}
}
