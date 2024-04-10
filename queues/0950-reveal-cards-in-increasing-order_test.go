package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestDeckRevealedIncreasing$
func TestDeckRevealedIncreasing(t *testing.T) {
	for _, tc := range []struct {
		deck     []int
		revealed []int
	}{
		{deck: []int{17, 13, 11, 2, 3, 5, 7}, revealed: []int{2, 13, 3, 11, 5, 17, 7}},
		{deck: []int{1, 1000}, revealed: []int{1, 1000}},
		{deck: []int{4728}, revealed: []int{4728}},
	} {
		revealed := deckRevealedIncreasing(tc.deck)
		assert.Equal(t, tc.revealed, revealed)
	}
}
