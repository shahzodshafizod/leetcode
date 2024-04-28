package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestSumOfDistancesInTree$
func TestSumOfDistancesInTree(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		answer []int
	}{
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, answer: []int{8, 12, 6, 10, 10, 10}},
		{n: 1, edges: [][]int{}, answer: []int{0}},
		{n: 2, edges: [][]int{{1, 0}}, answer: []int{1, 1}},
	} {
		answer := sumOfDistancesInTree(tc.n, tc.edges)
		assert.Equal(t, tc.answer, answer)
	}
}
