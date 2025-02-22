package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestFlipAndInvertImage$
func TestFlipAndInvertImage(t *testing.T) {
	for _, tc := range []struct {
		image   [][]int
		flipped [][]int
	}{
		{image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, flipped: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{image: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, flipped: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	} {
		flipped := flipAndInvertImage(tc.image)
		assert.Equal(t, tc.flipped, flipped)
	}
}
