package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestShortestPathAllKeys$
func TestShortestPathAllKeys(t *testing.T) {
	for _, tc := range []struct {
		grid   []string
		output int
	}{
		{grid: []string{"@.a..", "###.#", "b.A.B"}, output: 8},
		{grid: []string{"@..aA", "..B#.", "....b"}, output: 6},
		{grid: []string{"@Aa"}, output: -1},
		{grid: []string{"@...a", ".###A", "b.BCc"}, output: 10},
	} {
		output := shortestPathAllKeys(tc.grid)
		assert.Equal(t, tc.output, output)
	}
}
