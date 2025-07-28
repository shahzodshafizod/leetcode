package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestLRUCache$
func TestLRUCache(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			values:   [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			output:   []any{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
		{
			commands: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			values:   [][]int{{2}, {1, 0}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			output:   []any{nil, nil, nil, 0, nil, -1, nil, -1, 3, 4},
		},
		{
			commands: []string{"LRUCache", "put", "put", "put", "put", "get", "get"},
			values:   [][]int{{2}, {2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}},
			output:   []any{nil, nil, nil, nil, nil, -1, 3},
		},
		{
			commands: []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"},
			values:   [][]int{{2}, {2}, {2, 6}, {1}, {1, 5}, {1, 2}, {1}, {2}},
			output:   []any{nil, -1, nil, -1, nil, nil, 2, 6},
		},
	} {
		var cache LRUCache

		for index, command := range tc.commands {
			var output any = nil

			switch command {
			case "LRUCache":
				cache = NewLRUCache(tc.values[index][0])
			case "get":
				output = cache.Get(tc.values[index][0])
			case "put":
				cache.Put(tc.values[index][0], tc.values[index][1])
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
