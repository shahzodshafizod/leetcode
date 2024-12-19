package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestIsPrintable$
func TestIsPrintable(t *testing.T) {
	for _, tc := range []struct {
		targetGrid [][]int
		possible   bool
	}{
		{targetGrid: [][]int{{1, 1, 1, 1}, {1, 2, 2, 1}, {1, 2, 2, 1}, {1, 1, 1, 1}}, possible: true},
		{targetGrid: [][]int{{1, 1, 1, 1}, {1, 1, 3, 3}, {1, 1, 3, 4}, {5, 5, 1, 4}}, possible: true},
		{targetGrid: [][]int{{1, 2, 1}, {2, 1, 2}, {1, 2, 1}}, possible: false},
	} {
		possible := isPrintable(tc.targetGrid)
		assert.Equal(t, tc.possible, possible)
	}
}
