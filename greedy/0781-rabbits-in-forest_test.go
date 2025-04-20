package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestNumRabbits$
func TestNumRabbits(t *testing.T) {
	for _, tc := range []struct {
		answers []int
		total   int
	}{
		{answers: []int{1, 1, 2}, total: 5},
		{answers: []int{10, 10, 10}, total: 11},
		{answers: []int{0, 0, 1, 1, 1}, total: 6},
		{answers: []int{24, 24, 24, 24, 2, 2, 1, 2, 8, 8, 2, 8, 8}, total: 42},
		{answers: []int{2, 1, 0, 0, 0, 0, 0}, total: 10},
		{answers: []int{5, 1, 5, 3, 5, 2, 12, 4, 20, 20}, total: 54},
		{answers: []int{33, 25, 22, 2, 39, 15, 50, 7, 5, 12, 12, 12, 12, 3, 4}, total: 229},
	} {
		total := numRabbits(tc.answers)
		assert.Equal(t, tc.total, total)
	}
}
