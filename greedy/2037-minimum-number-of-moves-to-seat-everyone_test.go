package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinMovesToSeat$
func TestMinMovesToSeat(t *testing.T) {
	for _, tc := range []struct {
		seats    []int
		students []int
		moves    int
	}{
		{seats: []int{3, 1, 5}, students: []int{2, 7, 4}, moves: 4},
		{seats: []int{4, 1, 5, 9}, students: []int{1, 3, 2, 6}, moves: 7},
		{seats: []int{2, 2, 6, 6}, students: []int{1, 3, 2, 6}, moves: 4},
	} {
		moves := minMovesToSeat(tc.seats, tc.students)
		assert.Equal(t, tc.moves, moves)
	}
}
