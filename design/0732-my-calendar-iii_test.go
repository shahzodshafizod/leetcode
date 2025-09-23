package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyCalendarThree$
func TestMyCalendarThree(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MyCalendarThree", "book", "book", "book", "book", "book", "book"},
			values:   [][]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}},
			output:   []any{nil, 1, 1, 2, 3, 3, 3},
		},
		{
			commands: []string{"MyCalendarThree", "book"},
			values:   [][]int{{}, {24, 40}, {43, 50}, {27, 43}, {5, 21}, {30, 40}, {14, 29}, {3, 19}, {3, 14}, {25, 39}, {6, 19}},
			output:   []any{nil, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4},
		},
	} {
		var calendar MyCalendarThree

		for idx, command := range tc.commands {
			var output any

			switch command {
			case "MyCalendarThree":
				calendar = NewMyCalendarThree()
			case "book":
				output = calendar.Book(tc.values[idx][0], tc.values[idx][1])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
