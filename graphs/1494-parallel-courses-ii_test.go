package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinNumberOfSemesters$
func TestMinNumberOfSemesters(t *testing.T) {
	for _, tc := range []struct {
		n         int
		relations [][]int
		k         int
		mincount  int
	}{
		{n: 11, relations: [][]int{}, k: 2, mincount: 6},
		{n: 4, relations: [][]int{{2, 1}, {3, 1}, {1, 4}}, k: 2, mincount: 3},
		{n: 5, relations: [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}, k: 2, mincount: 4},
		{n: 12, relations: [][]int{{1, 2}, {1, 3}, {7, 5}, {7, 6}, {4, 8}, {8, 9}, {9, 10}, {10, 11}, {11, 12}}, k: 2, mincount: 6},
		{n: 13, relations: [][]int{{12, 8}, {2, 4}, {3, 7}, {6, 8}, {11, 8}, {9, 4}, {9, 7}, {12, 4}, {11, 4}, {6, 4}, {1, 4}, {10, 7}, {10, 4}, {1, 7}, {1, 8}, {2, 7}, {8, 4}, {10, 8}, {12, 7}, {5, 4}, {3, 4}, {11, 7}, {7, 4}, {13, 4}, {9, 8}, {13, 8}}, k: 9, mincount: 3},
	} {
		mincount := minNumberOfSemesters(tc.n, tc.relations, tc.k)
		assert.Equal(t, tc.mincount, mincount)
	}
}
