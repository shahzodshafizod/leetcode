package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinCost1578$
func TestMinCost1578(t *testing.T) {
	for _, tc := range []struct {
		colors     string
		neededTime []int
		time       int
	}{
		{colors: "abaac", neededTime: []int{1, 2, 3, 4, 5}, time: 3},
		{colors: "abc", neededTime: []int{1, 2, 3}, time: 0},
		{colors: "aabaa", neededTime: []int{1, 2, 3, 4, 1}, time: 2},
	} {
		time := minCost1578(tc.colors, tc.neededTime)
		assert.Equal(t, tc.time, time)
	}
}
