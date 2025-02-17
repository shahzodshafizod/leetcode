package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestNumTilePossibilities$
func TestNumTilePossibilities(t *testing.T) {
	for _, tc := range []struct {
		tiles string
		count int
	}{
		{tiles: "AAB", count: 8},
		{tiles: "AAABBC", count: 188},
		{tiles: "V", count: 1},
	} {
		count := numTilePossibilities(tc.tiles)
		assert.Equal(t, tc.count, count)
	}
}
