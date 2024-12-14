package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestLFUCache$
func TestLFUCache(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"},
			values:   [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}},
			output:   []any{nil, nil, nil, 1, nil, -1, 3, nil, -1, 3, 4},
		},
	} {
		var lfu LFUCache
		var output any
		for idx, command := range tc.commands {
			output = nil
			switch command {
			case "LFUCache":
				lfu = NewLFUCache(tc.values[idx][0])
			case "get":
				output = lfu.Get(tc.values[idx][0])
			case "put":
				lfu.Put(tc.values[idx][0], tc.values[idx][1])
			}
			assert.Equal(t, tc.output[idx], output)
		}
	}
}
