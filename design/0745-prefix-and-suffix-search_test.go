package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestWordFilter$
func TestWordFilter(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]string
		output   []any
	}{
		{
			commands: []string{"WordFilter", "f", "f"},
			values:   [][]string{{"apple", "app", "ape"}, {"a", "e"}, {"appl", "pple"}},
			output:   []any{nil, 2, 0},
		},
		{
			commands: []string{"WordFilter", "f"},
			values:   [][]string{{"abbba", "abba"}, {"ab", "ba"}},
			output:   []any{nil, 1},
		},
	} {
		var filter WordFilter

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "WordFilter":
				filter = NewWordFilter(tc.values[idx])
			case "f":
				output = filter.F(tc.values[idx][0], tc.values[idx][1])
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
