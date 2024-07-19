package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestLuckyNumbers$
func TestLuckyNumbers(t *testing.T) {
	for _, tc := range []struct {
		matrix  [][]int
		luckies []int
	}{
		{matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, luckies: []int{15}},
		{matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, luckies: []int{12}},
		{matrix: [][]int{{7, 8}, {1, 2}}, luckies: []int{7}},
	} {
		luckies := luckyNumbers(tc.matrix)
		assert.Equal(t, tc.luckies, luckies)
	}
}
