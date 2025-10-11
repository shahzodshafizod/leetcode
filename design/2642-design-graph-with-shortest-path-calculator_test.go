package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestGraph$
func TestGraph(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   []any
		output   []any
	}{
		{
			[]string{"Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"},
			[]any{[]any{4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}}}, []int{3, 2}, []int{0, 3}, [][]int{{1, 3, 4}}, []int{0, 3}},
			[]any{nil, 6, -1, nil, 6},
		},
	} {
		var graph Graph

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "Graph":
				values, ok := tc.values[idx].([]any)
				_ = ok
				n, ok := values[0].(int)
				_ = ok
				edges, ok := values[1].([][]int)
				_ = ok
				graph = NewGraph(n, edges)
			case "addEdge":
				edge, ok := tc.values[idx].([][]int)
				_ = ok

				graph.AddEdge(edge[0])
			case "shortestPath":
				values, ok := tc.values[idx].([]int)
				_ = ok
				output = graph.ShortestPath(values[0], values[1])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
