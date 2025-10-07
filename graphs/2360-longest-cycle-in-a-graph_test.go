package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLongestCycle$
func TestLongestCycle(t *testing.T) {
	for _, tc := range []struct {
		edges  []int
		length int
	}{
		{edges: []int{3, 3, 4, 2, 3}, length: 3},
		{edges: []int{2, -1, 3, 1}, length: -1},
		{edges: []int{-1, 4, -1, 2, 0, 4}, length: -1},
		{edges: []int{1, 0}, length: 2},
		{edges: []int{1, 19, 30, 87, 53, 91, 36, 6, 95, 14, 73, 2, 59, 76, 49, 41, 29, 28, 8, 9, 96, 80, 68, 10, 31, 24, 0, 42, 39, 4, 51, 64, 25, 90, 35, 71, 97, 32, 16, 18, 62, 22, 40, 78, 55, 13, 99, 93, 66, 26, 98, 5, 88, 74, 89, 81, 43, 12, 44, 57, 75, 47, 34, 72, 85, 77, 3, 65, 46, 20, 60, 33, 48, 94, 84, 21, 69, 54, 56, 11, 70, 83, 86, 79, 61, 37, 67, 15, 7, 38, 23, 52, 58, 27, 50, 63, 92, 45, 17, 82}, length: 50},
		{edges: []int{49, 29, 24, 24, -1, -1, -1, 2, 23, -1, 44, 47, 52, 49, 9, 31, 40, 34, -1, 53, 33, -1, 2, -1, 18, 31, 0, 9, 47, 35, -1, -1, -1, -1, 4, 12, 14, 19, -1, 4, 4, 43, 25, 11, 54, -1, 25, 39, 17, -1, 22, 44, -1, 44, 29, 50, -1, -1}, length: -1},
	} {
		length := longestCycle(tc.edges)
		assert.Equal(t, tc.length, length)
	}
}
