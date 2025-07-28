package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestTrie$
func TestTrie(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]string
		output   []any
	}{
		{
			commands: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			values:   [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			output:   []any{nil, nil, true, false, true, nil, true},
		},
		{
			commands: []string{"Trie", "insert", "search", "insert", "search", "startsWith", "search", "insert", "search"},
			values:   [][]string{{}, {"apple"}, {"dog"}, {"dog"}, {"dog"}, {"app"}, {"app"}, {"app"}, {"app"}},
			output:   []any{nil, nil, false, nil, true, true, false, nil, true},
		},
		{
			commands: []string{"Trie", "startsWith"},
			values:   [][]string{{}, {"a"}},
			output:   []any{nil, false},
		},
		{
			commands: []string{"Trie", "insert", "search", "startsWith"},
			values:   [][]string{{}, {"a"}, {"a"}, {"a"}},
			output:   []any{nil, nil, true, true},
		},
	} {
		var trie Trie

		for index, command := range tc.commands {
			var output any = nil

			switch command {
			case "Trie":
				trie = NewTrie()
			case "insert":
				trie.Insert(tc.values[index][0])
			case "search":
				output = trie.Search(tc.values[index][0])
			case "startsWith":
				output = trie.StartsWith(tc.values[index][0])
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
