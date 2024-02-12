package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestTraversalBFS$
func TestTraversalBFS(t *testing.T) {
	for _, tc := range []struct {
		graph  [][]int
		values []int
	}{
		{
			graph:  [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values: []int{0, 1, 3, 4, 5, 2, 6, 8, 7},
		},
	} {
		values := traversalBFS(tc.graph)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestTraversalDFS$
func TestTraversalDFS(t *testing.T) {
	for _, tc := range []struct {
		graph  [][]int
		values []int
	}{
		{
			graph:  [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values: []int{0, 1, 3, 4, 6, 7, 5, 2, 8},
		},
	} {
		values := traversalDFS(tc.graph)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestUnionFind$
func TestUnionFind(t *testing.T) {
	var uf UnionFind
	var m, n = 4, 3
	for _, tc := range []struct {
		command string
		value   []int
		output  any
	}{
		{command: "UnionFind", value: []int{m + n}, output: nil},

		{command: "Union", value: []int{1, 0 + m}, output: nil},
		{command: "Union", value: []int{1, 2 + m}, output: nil},
		{command: "Union", value: []int{2, 1 + m}, output: nil},
		{command: "Union", value: []int{3, 2 + m}, output: nil},
		{command: "Find", value: []int{1}, output: 3},
		{command: "Find", value: []int{2}, output: 2},
		{command: "Find", value: []int{3}, output: 3},

		{command: "Union", value: []int{0, 0 + m}, output: nil},
		{command: "Find", value: []int{1}, output: 0},
		{command: "Find", value: []int{2}, output: 2},
		{command: "Find", value: []int{3}, output: 0},
		{command: "Find", value: []int{0}, output: 0},

		{command: "Union", value: []int{2, 0 + m}, output: nil},
		{command: "Find", value: []int{1}, output: 2},
		{command: "Find", value: []int{2}, output: 2},
		{command: "Find", value: []int{3}, output: 2},
		{command: "Find", value: []int{0}, output: 2},
	} {
		var output any = nil
		switch tc.command {
		case "UnionFind":
			uf = NewUnionFind(tc.value[0])
		case "Union":
			uf.Union(tc.value[0], tc.value[1])
		case "Find":
			output = uf.Find(tc.value[0])
		}
		assert.Equal(t, tc.output, output)
	}
}
