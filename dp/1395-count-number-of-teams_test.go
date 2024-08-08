package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumTeams$
func TestNumTeams(t *testing.T) {
	for _, tc := range []struct {
		rating []int
		teams  int
	}{
		{rating: []int{2, 5, 3, 4, 1}, teams: 3},
		{rating: []int{2, 1, 3}, teams: 0},
		{rating: []int{1, 2, 3, 4}, teams: 4},
		{rating: []int{1, 3, 6, 8, 4}, teams: 5},
	} {
		teams := numTeams(tc.rating)
		assert.Equal(t, tc.teams, teams)
	}
}
