package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumMusicPlaylists$
func TestNumMusicPlaylists(t *testing.T) {
	for _, tc := range []struct {
		n     int
		goal  int
		k     int
		count int
	}{
		{n: 3, goal: 3, k: 1, count: 6},
		{n: 2, goal: 3, k: 0, count: 6},
		{n: 2, goal: 3, k: 1, count: 2},
		{n: 2, goal: 4, k: 0, count: 14},
		{n: 16, goal: 16, k: 4, count: 789741546},
	} {
		count := numMusicPlaylists(tc.n, tc.goal, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
