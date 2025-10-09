package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindAnswer$
func TestFindAnswer(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		answer []bool
	}{
		{n: 6, edges: [][]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2}}, answer: []bool{true, true, true, false, true, true, true, false}},
		{n: 4, edges: [][]int{{2, 0, 1}, {0, 1, 1}, {0, 3, 4}, {3, 2, 2}}, answer: []bool{true, false, false, true}},
	} {
		answer := findAnswer(tc.n, tc.edges)
		assert.Equal(t, tc.answer, answer)
	}
}
