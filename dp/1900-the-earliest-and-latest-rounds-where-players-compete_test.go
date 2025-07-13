package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestEarliestAndLatest$
func TestEarliestAndLatest(t *testing.T) {
	for _, tc := range []struct {
		n int
		firstPlayer int
		secondPlayer int
		rounds []int
	}{
		{n: 11, firstPlayer: 2, secondPlayer: 4, rounds: []int{3, 4}},
		{n: 5, firstPlayer: 1, secondPlayer: 5, rounds: []int{1, 1}},
		{n: 11, firstPlayer: 4, secondPlayer: 11, rounds: []int{2, 4}},
	} {
		rounds := earliestAndLatest(tc.n, tc.firstPlayer, tc.secondPlayer)
		assert.Equal(t, tc.rounds, rounds)
	}
}
