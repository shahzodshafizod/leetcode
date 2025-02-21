package design

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestFindElements$
func TestFindElements(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]any
		output   []any
	}{
		{
			commands: []string{"FindElements", "find", "find"},
			values:   [][]any{{[]any{-1, nil, -1}}, {1}, {2}},
			output:   []any{nil, false, true},
		},
		{
			commands: []string{"FindElements", "find", "find", "find"},
			values:   [][]any{{[]any{-1, -1, -1, -1, -1}}, {1}, {3}, {5}},
			output:   []any{nil, true, true, false},
		},
		{
			commands: []string{"FindElements", "find", "find", "find", "find"},
			values:   [][]any{{[]any{-1, nil, -1, -1, nil, -1}}, {2}, {3}, {4}, {5}},
			output:   []any{nil, true, false, false, true},
		},
	} {
		var finder FindElements
		var output any
		for idx, command := range tc.commands {
			output = nil
			switch command {
			case "FindElements":
				finder = NewFindElements(pkg.MakeTree2(tc.values[idx][0].([]any)...))
			case "find":
				output = finder.Find(tc.values[idx][0].(int))
			}
			assert.Equal(t, tc.output[idx], output)
		}
	}
}
