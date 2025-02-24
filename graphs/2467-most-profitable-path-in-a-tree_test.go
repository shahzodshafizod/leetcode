package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMostProfitablePath$
func TestMostProfitablePath(t *testing.T) {
	for _, tc := range []struct {
		edges    [][]int
		bob      int
		amount   []int
		aliceNet int
	}{
		{edges: [][]int{{0, 1}}, bob: 1, amount: []int{-7280, 2350}, aliceNet: -7280},
		{edges: [][]int{{0, 1}, {0, 2}}, bob: 2, amount: []int{-3360, -5394, -1146}, aliceNet: -3360},
		{edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, bob: 3, amount: []int{-2, 4, 2, -4, 6}, aliceNet: 6},
		{edges: [][]int{{0, 1}, {1, 2}, {2, 3}}, bob: 3, amount: []int{-5644, -6018, 1188, -8502}, aliceNet: -11662},
		{edges: [][]int{{0, 2}, {0, 4}, {1, 3}, {1, 2}}, bob: 1, amount: []int{3958, -9854, -8334, -9388, 3410}, aliceNet: 7368},
		{edges: [][]int{{0, 2}, {0, 5}, {1, 3}, {1, 5}, {2, 4}}, bob: 4, amount: []int{5018, 8388, 6224, 3466, 3808, 3456}, aliceNet: 20328},
		{edges: [][]int{{0, 2}, {0, 6}, {1, 3}, {1, 5}, {2, 5}, {4, 6}}, bob: 6, amount: []int{8838, -6396, -5940, 2694, -1366, 4616, 2966}, aliceNet: 7472},
		{edges: [][]int{{0, 2}, {1, 4}, {1, 6}, {2, 4}, {3, 6}, {3, 7}, {5, 7}}, bob: 4, amount: []int{-6896, -1216, -1208, -1108, 1606, -7704, -9212, -8258}, aliceNet: -34998},
		{edges: [][]int{{0, 2}, {1, 4}, {1, 6}, {2, 3}, {2, 8}, {3, 7}, {4, 5}, {6, 7}}, bob: 1, amount: []int{-1410, -9440, 5536, -774, 3044, 7924, -2122, -1192, 9166}, aliceNet: 14320},
	} {
		aliceNet := mostProfitablePath(tc.edges, tc.bob, tc.amount)
		assert.Equal(t, tc.aliceNet, aliceNet)
	}
}
