package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestFindTheWinner$
func TestFindTheWinner(t *testing.T) {
	for _, tc := range []struct {
		n      int
		k      int
		winner int
	}{
		{n: 5, k: 2, winner: 3},
		{n: 6, k: 5, winner: 1},
	} {
		winner := findTheWinner(tc.n, tc.k)
		assert.Equal(t, tc.winner, winner)
	}
}
