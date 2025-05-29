package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaxTargetNodes3373$
func TestMaxTargetNodes3373(t *testing.T) {
	for _, tc := range []struct {
		edges1 [][]int
		edges2 [][]int
		answer []int
	}{
		{edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}, answer: []int{8, 7, 7, 8, 8}},
		{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, edges2: [][]int{{0, 1}, {1, 2}, {2, 3}}, answer: []int{3, 6, 6, 6, 6}},
	} {
		answer := maxTargetNodes3373(tc.edges1, tc.edges2)
		assert.Equal(t, tc.answer, answer)
	}
}
