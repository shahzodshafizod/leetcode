package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFloodFill$
func TestFloodFill(t *testing.T) {
	for _, tc := range []struct {
		image    [][]int
		sr       int
		sc       int
		color    int
		modified [][]int
	}{
		{image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, sr: 1, sc: 1, color: 2, modified: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{image: [][]int{{0, 0, 0}, {0, 0, 0}}, sr: 0, sc: 0, color: 0, modified: [][]int{{0, 0, 0}, {0, 0, 0}}},
		{image: [][]int{{0, 0, 0}, {0, 0, 0}}, sr: 0, sc: 0, color: 2, modified: [][]int{{2, 2, 2}, {2, 2, 2}}},
		{image: [][]int{{0, 0, 0}, {0, 0, 0}}, sr: 1, sc: 0, color: 2, modified: [][]int{{2, 2, 2}, {2, 2, 2}}},
		{image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, sr: 1, sc: 1, color: 2, modified: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{image: [][]int{{0, 0, 0}, {0, 0, 1}}, sr: 1, sc: 0, color: 2, modified: [][]int{{2, 2, 2}, {2, 2, 1}}},
	} {
		modified := floodFill(tc.image, tc.sr, tc.sc, tc.color)
		assert.Equal(t, tc.modified, modified)
	}
}
