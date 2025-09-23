package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestNumArray$
func TestNumArray(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"NumArray", "sumRange", "sumRange", "sumRange"},
			values:   [][]int{{-2, 0, 3, -5, 2, -1}, {0, 2}, {2, 5}, {0, 5}},
			output:   []any{nil, 1, -1, -3},
		},
	} {
		var numArray NumArray

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "NumArray":
				numArray = NewNumArray(tc.values[idx])
			case "sumRange":
				output = numArray.SumRange(tc.values[idx][0], tc.values[idx][1])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
