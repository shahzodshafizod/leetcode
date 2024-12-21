package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaxKDivisibleComponents$
func TestMaxKDivisibleComponents(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		values []int
		k      int
		count  int
	}{
		{n: 1, edges: [][]int{}, values: []int{0}, k: 1, count: 1},
		{n: 1, edges: [][]int{}, values: []int{100}, k: 100, count: 1},
		{n: 5, edges: [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, values: []int{1, 8, 1, 4, 4}, k: 6, count: 2},
		{n: 4, edges: [][]int{{0, 1}, {1, 2}, {1, 3}}, values: []int{176968959, 450655404, 922326524, 897145068}, k: 5, count: 1},
		{n: 4, edges: [][]int{{3, 0}, {0, 2}, {0, 1}}, values: []int{742593846, 612800589, 813572140, 813559096}, k: 47, count: 2},
		{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, values: []int{3, 0, 6, 1, 5, 2, 1}, k: 3, count: 3},
		{n: 5, edges: [][]int{{1, 2}, {1, 3}, {0, 2}, {2, 4}}, values: []int{999999999, 999999999, 999999999, 999999999, 999999999}, k: 5, count: 1},
		{n: 7, edges: [][]int{{2, 0}, {0, 4}, {2, 5}, {1, 2}, {5, 3}, {4, 6}}, values: []int{12154284, 649536765, 974051464, 821507385, 392654193, 770357917, 37707285}, k: 11, count: 2},
		// {n: 20, edges: [][]int{{17, 4}, {4, 2}, {3, 15}, {15, 19}, {19, 7}, {7, 12}, {15, 16}, {19, 18}, {18, 13}, {19, 9}, {19, 1}, {12, 8}, {1, 11}, {7, 6}, {9, 0}, {15, 10}, {13, 14}, {18, 5}, {4, 5}}, values: []int{900160, 891774, 2283737, 414736, 265741, 212260, 2983538, 280834, 17654, 221822, 996740, 1517078, 374434, 43488, 610048, 241798, 611382, 3234714, 1156844, 93794}, k: 5784192, count: 2},
	} {
		count := maxKDivisibleComponents(tc.n, tc.edges, tc.values, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
