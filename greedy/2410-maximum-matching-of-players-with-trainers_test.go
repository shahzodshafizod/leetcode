package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMatchPlayersAndTrainers$
func TestMatchPlayersAndTrainers(t *testing.T) {
	for _, tc := range []struct {
		players []int
		trainers []int
		count int
	}{
		{players: []int{4, 7, 9}, trainers: []int{8, 2, 5, 8}, count: 2},
		{players: []int{1, 1, 1}, trainers: []int{10}, count: 1},
	} {
		count := matchPlayersAndTrainers(tc.players, tc.trainers)
		assert.Equal(t, tc.count, count)
	}
}
