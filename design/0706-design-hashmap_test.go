package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyHashMap$
func TestMyHashMap(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"},
			values:   [][]int{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}},
			output:   []any{nil, nil, nil, 1, -1, nil, 1, nil, -1},
		},
	} {
		var myHashMap MyHashMap

		for index, command := range tc.commands {
			var output any = nil

			switch command {
			case "MyHashMap":
				myHashMap = NewMyHashMap()
			case "put":
				myHashMap.Put(tc.values[index][0], tc.values[index][1])
			case "get":
				output = myHashMap.Get(tc.values[index][0])
			case "remove":
				myHashMap.Remove(tc.values[index][0])
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
