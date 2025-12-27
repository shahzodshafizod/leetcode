package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestMaxCount$
func TestMaxCount(t *testing.T) {
	for _, tc := range []struct {
		m      int
		n      int
		ops    [][]int
		output int
	}{
		{m: 3, n: 3, ops: [][]int{{2, 2}, {3, 3}}, output: 4},
		{m: 3, n: 3, ops: [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}, output: 4},
		{m: 3, n: 3, ops: [][]int{}, output: 9},
		{m: 40000, n: 40000, ops: [][]int{}, output: 1600000000},
		{m: 3, n: 3, ops: [][]int{{1, 1}, {2, 2}, {3, 3}}, output: 1},
	} {
		output := maxCount(tc.m, tc.n, tc.ops)
		assert.Equal(t, tc.output, output)
	}
}
