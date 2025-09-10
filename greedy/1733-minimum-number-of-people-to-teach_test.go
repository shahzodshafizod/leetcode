package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumTeachings$
func TestMinimumTeachings(t *testing.T) {
	for _, tc := range []struct {
		n           int
		languages   [][]int
		friendships [][]int
		teachings   int
	}{
		{n: 2, languages: [][]int{{1}, {2}, {1, 2}}, friendships: [][]int{{1, 2}, {1, 3}, {2, 3}}, teachings: 1},
		{n: 3, languages: [][]int{{2}, {1, 3}, {1, 2}, {3}}, friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}}, teachings: 2},
	} {
		teachings := minimumTeachings(tc.n, tc.languages, tc.friendships)
		assert.Equal(t, tc.teachings, teachings)
	}
}
