package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxProfitAssignment$
func TestMaxProfitAssignment(t *testing.T) {
	for _, tc := range []struct {
		difficulty []int
		profit     []int
		worker     []int
		maxprofit  int
	}{
		{
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			maxprofit:  100,
		},
		{
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			maxprofit:  0,
		},
		{
			difficulty: []int{68, 35, 52, 47, 86},
			profit:     []int{67, 17, 1, 81, 3},
			worker:     []int{92, 10, 85, 84, 82},
			maxprofit:  324,
		},
		{
			difficulty: []int{2, 5, 2, 66, 56, 23, 90},
			profit:     []int{1, 2, 2, 34, 36, 49, 50},
			worker:     []int{2, 45, 13, 15, 67, 800},
			maxprofit:  154,
		},
		{
			difficulty: []int{71, 42, 36, 83, 22},
			profit:     []int{32, 45, 39, 54, 23},
			worker:     []int{74, 21, 31, 49, 88},
			maxprofit:  167,
		},
	} {
		maxprofit := maxProfitAssignment(tc.difficulty, tc.profit, tc.worker)
		assert.Equal(t, tc.maxprofit, maxprofit)
	}
}
