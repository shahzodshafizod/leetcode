package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaximumInvitations$
func TestMaximumInvitations(t *testing.T) {
	for _, tc := range []struct {
		favorite []int
		maxnum   int
	}{
		{favorite: []int{2, 2, 1, 2}, maxnum: 3},
		{favorite: []int{1, 2, 0}, maxnum: 3},
		{favorite: []int{3, 0, 1, 4, 1}, maxnum: 4},
	} {
		maxnum := maximumInvitations(tc.favorite)
		assert.Equal(t, tc.maxnum, maxnum)
	}
}
