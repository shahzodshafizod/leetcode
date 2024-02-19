package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyHashSet$
func TestMyHashSet(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"},
			values:   [][]int{{}, {1}, {2}, {1}, {3}, {2}, {2}, {2}, {2}},
			output:   []any{nil, nil, nil, true, false, nil, true, nil, false},
		},
	} {
		var myHashSet MyHashSet
		for index, command := range tc.commands {
			var output any = nil
			switch command {
			case "MyHashSet":
				myHashSet = NewMyHashSet()
			case "add":
				myHashSet.Add(tc.values[index][0])
			case "remove":
				myHashSet.Remove(tc.values[index][0])
			case "contains":
				output = myHashSet.Contains(tc.values[index][0])
			}
			assert.Equal(t, tc.output[index], output)
		}
	}
}
