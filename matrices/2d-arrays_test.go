package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestTraversalDFS$
func TestTraversalDFS(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		values []int
	}{
		{matrix: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}},
			values: []int{1, 2, 3, 4, 5, 10, 15, 20, 19, 14, 9, 8, 13, 18, 17, 12, 7, 6, 11, 16}},
		{matrix: nil, values: nil},
		{matrix: [][]int{}, values: nil},
		{matrix: [][]int{{-1, 2, 3}, {0, 9, 8}, {1, 0, 1}}, values: []int{-1, 2, 3, 8, 1, 0, 9, 0, 1}},
		{matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}, values: []int{1, 2, 3, 7, 11, 10, 6, 5, 9}},
	} {
		values := traversalDFS(tc.matrix)
		assert.Equal(t, tc.values, values)

		values = traversalDFSiter(tc.matrix)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./matrices/ -run ^TestTraversalBFS$
func TestTraversalBFS(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		values []int
	}{
		{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, values: []int{1, 2, 5, 3, 6, 9, 4, 7, 10, 13, 8, 11, 14, 12, 15, 16}},
		{matrix: [][]int{{-1, 0, 0, 1}, {-1, -1, -2, -1}, {-1, -1, -1, -1}, {0, 0, 0, 0}}, values: []int{-1, 0, -1, 0, -1, -1, 1, -2, -1, 0, -1, -1, 0, -1, 0, 0}},
	} {
		values := traversalBFS(tc.matrix)
		assert.Equal(t, tc.values, values)

		// values = traversalBFSiter(tc.matrix)
		// assert.Equal(t, tc.values, values)
	}
}
