package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaxTargetNodes$
func TestMaxTargetNodes(t *testing.T) {
	for _, tc := range []struct {
		edges1 [][]int
		edges2 [][]int
		k      int
		answer []int
	}{
		{edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}, k: 2, answer: []int{9, 7, 9, 8, 8}},
		{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, edges2: [][]int{{0, 1}, {1, 2}, {2, 3}}, k: 1, answer: []int{6, 3, 3, 3, 3}},
	} {
		answer := maxTargetNodes(tc.edges1, tc.edges2, tc.k)
		assert.Equal(t, tc.answer, answer)
	}
}
