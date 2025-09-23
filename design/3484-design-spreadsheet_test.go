package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestSpreadsheet$
func TestSpreadsheet(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   []any
		output   []any
	}{
		{
			commands: []string{"Spreadsheet", "getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"},
			values:   []any{[]int{3}, []string{"=5+7"}, []any{"A1", 10}, []string{"=A1+6"}, []any{"B2", 15}, []string{"=A1+B2"}, []string{"A1"}, []string{"=A1+B2"}},
			output:   []any{nil, 12, nil, 16, nil, 25, nil, 15},
		},
	} {
		var sheet Spreadsheet

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "Spreadsheet":
				values, ok := tc.values[idx].([]int)
				_ = ok
				sheet = NewSpreadsheet(values[0])
			case "setCell":
				values, ok := tc.values[idx].([]any)
				_ = ok
				cell, ok := values[0].(string)
				_ = ok
				value, ok := values[1].(int)
				_ = ok

				sheet.SetCell(cell, value)
			case "resetCell":
				values, ok := tc.values[idx].([]string)
				_ = ok

				sheet.ResetCell(values[0])
			case "getValue":
				values, ok := tc.values[idx].([]string)
				_ = ok
				output = sheet.GetValue(values[0])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
