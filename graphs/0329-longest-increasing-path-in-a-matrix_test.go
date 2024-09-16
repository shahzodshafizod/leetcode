package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLongestIncreasingPath$
func TestLongestIncreasingPath(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		length int
	}{
		{matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, length: 4},
		{matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, length: 4},
		{matrix: [][]int{{1}}, length: 1},
	} {
		length := longestIncreasingPath(tc.matrix)
		assert.Equal(t, tc.length, length)
	}
}
