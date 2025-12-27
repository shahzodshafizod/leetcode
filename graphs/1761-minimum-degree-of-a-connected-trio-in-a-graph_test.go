package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinTrioDegree$
func TestMinTrioDegree(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		output int
	}{
		{n: 6, edges: [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}, output: 3},
		{n: 7, edges: [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}, output: 0},
		{n: 4, edges: [][]int{{1, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 4}}, output: 2},
	} {
		output := minTrioDegree(tc.n, tc.edges)
		assert.Equal(t, tc.output, output)
	}
}
