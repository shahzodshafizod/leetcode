package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestNumRescueBoats$
func TestNumRescueBoats(t *testing.T) {
	for _, tc := range []struct {
		people []int
		limit  int
		boats  int
	}{
		{people: []int{1, 2}, limit: 3, boats: 1},
		{people: []int{3, 2, 2, 1}, limit: 3, boats: 3},
		{people: []int{3, 5, 3, 4}, limit: 5, boats: 4},
		{people: []int{3, 5, 3, 4}, limit: 5, boats: 4},
		{people: []int{2, 4}, limit: 5, boats: 2},
	} {
		boats := numRescueBoats(tc.people, tc.limit)
		assert.Equal(t, tc.boats, boats)
	}
}
