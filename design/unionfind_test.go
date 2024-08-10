package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestUnionFind$
func TestUnionFind(t *testing.T) {
	var ds DisjointSet
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
			ds = NewDisjointSet(tc.value[0])
		case "Union":
			ds.Union(tc.value[0], tc.value[1])
		case "Find":
			output = ds.Find(tc.value[0])
		}
		assert.Equal(t, tc.output, output)
	}
}

// go test -v -count=1 ./design/ -run ^TestDSQuickFind$
func TestDSQuickFind(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []bool
	}{
		{
			commands: []string{"NewQuickFind", "Union", "Union", "Union", "Union", "Union", "Union", "Connected", "Connected", "Connected", "Union", "Connected"},
			values:   [][]int{{10}, {1, 2}, {2, 5}, {5, 6}, {6, 7}, {3, 8}, {8, 9}, {1, 5}, {5, 7}, {4, 9}, {9, 4}, {4, 9}},
			output:   []bool{false, false, false, false, false, false, false, true, true, false, false, true},
		},
	} {
		var ds DisjointSet
		var output bool
		for idx, command := range tc.commands {
			output = false
			switch command {
			case "NewQuickFind":
				ds = NewOptimizedDS(tc.values[idx][0])
			case "Union":
				ds.Union(tc.values[idx][0], tc.values[idx][1])
			case "Connected":
				output = ds.Connected(tc.values[idx][0], tc.values[idx][1])
			}
			assert.Equal(t, tc.output[idx], output)
		}
	}
}
