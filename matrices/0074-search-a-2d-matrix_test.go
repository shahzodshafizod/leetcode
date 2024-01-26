package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSearchMatrix$
func TestSearchMatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		exists bool
	}{
		{matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3, exists: true},
		{matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 13, exists: false},
		{matrix: [][]int{{1}, {3}, {5}}, target: 3, exists: true},
		{matrix: [][]int{{1, 3}}, target: 3, exists: true},
		{matrix: [][]int{{1, 3}}, target: 2, exists: false},
		{matrix: [][]int{{1}}, target: 1, exists: true},
		{matrix: [][]int{{1, 2, 2, 2, 3, 5, 7, 9, 10}, {13, 15, 15, 16, 19, 19, 21, 21, 23}, {24, 26, 26, 27, 29, 29, 31, 31, 33}, {34, 35, 35, 35, 36, 36, 36, 37, 38}, {40, 41, 44, 45, 45, 48, 48, 51, 52}, {53, 53, 57, 57, 58, 59, 59, 59, 61}, {63, 64, 65, 67, 68, 69, 70, 70, 70}}, target: 31, exists: true},
	} {
		exists := searchMatrix(tc.matrix, tc.target)
		assert.Equal(t, tc.exists, exists)
	}
}
