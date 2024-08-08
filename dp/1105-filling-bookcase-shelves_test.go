package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinHeightShelves$
func TestMinHeightShelves(t *testing.T) {
	for _, tc := range []struct {
		books      [][]int
		shelfWidth int
		minHeight  int
	}{
		{books: [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, shelfWidth: 4, minHeight: 6},
		{books: [][]int{{1, 3}, {2, 4}, {3, 2}}, shelfWidth: 6, minHeight: 4},
		{books: [][]int{{7, 3}, {8, 7}, {2, 7}, {2, 5}}, shelfWidth: 10, minHeight: 15},
	} {
		minHeight := minHeightShelves(tc.books, tc.shelfWidth)
		assert.Equal(t, tc.minHeight, minHeight)
	}
}
