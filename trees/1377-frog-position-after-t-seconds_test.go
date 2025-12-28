package trees

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFrogPosition$
func TestFrogPosition(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		t      int
		target int
		result float64
	}{
		{
			n:      7,
			edges:  [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
			t:      2,
			target: 4,
			result: 0.16666666666666666,
		},
		{
			n:      7,
			edges:  [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
			t:      1,
			target: 7,
			result: 0.3333333333333333,
		},
		{
			n:      7,
			edges:  [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
			t:      20,
			target: 6,
			result: 0.16666666666666666,
		},
		{
			n:      3,
			edges:  [][]int{{2, 1}, {3, 2}},
			t:      1,
			target: 2,
			result: 1.0,
		},
		{
			n:      1,
			edges:  [][]int{},
			t:      1,
			target: 1,
			result: 1.0,
		},
		{
			n:      8,
			edges:  [][]int{{2, 1}, {3, 2}, {4, 1}, {5, 1}, {6, 4}, {7, 1}, {8, 7}},
			t:      7,
			target: 7,
			result: 0.0,
		},
	} {
		result := frogPosition(tc.n, tc.edges, tc.t, tc.target)
		assert.Less(t, math.Abs(result-tc.result), 1e-5, "expected: %v, got: %v", tc.result, result)
	}
}
