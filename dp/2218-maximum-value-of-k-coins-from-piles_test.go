package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxValueOfCoins$
func TestMaxValueOfCoins(t *testing.T) {
	for _, tc := range []struct {
		piles [][]int
		k     int
		value int
	}{
		{piles: [][]int{{1, 100, 3}, {7, 8, 9}}, k: 2, value: 101},
		{piles: [][]int{{1}, {100}, {100}, {100}, {100}, {100}, {10, 1, 1, 1, 1, 1, 700}}, k: 5, value: 500},
		{piles: [][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, k: 7, value: 706},
		{piles: [][]int{{80, 62, 78, 78, 40, 59, 98, 35}, {79, 19, 100, 15}, {79, 2, 27, 73, 12, 13, 11, 37, 27, 55, 54, 55, 87, 10, 97, 26, 78, 20, 75, 23, 46, 94, 56, 32, 14, 70, 70, 37, 60, 46, 1, 53}}, k: 25, value: 1332},
	} {
		value := maxValueOfCoins(tc.piles, tc.k)
		assert.Equal(t, tc.value, value)
	}
}
