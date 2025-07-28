package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestBrowserHistory$
func TestBrowserHistory(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]any
		output   []any
	}{
		{
			commands: []string{"BrowserHistory", "visit", "visit", "visit", "back", "back", "forward", "visit", "forward", "back", "back"},
			values:   [][]any{{"leetcode.com"}, {"google.com"}, {"facebook.com"}, {"youtube.com"}, {1}, {1}, {1}, {"linkedin.com"}, {2}, {2}, {7}},
			output:   []any{nil, nil, nil, nil, "facebook.com", "google.com", "facebook.com", nil, "linkedin.com", "google.com", "leetcode.com"},
		},
	} {
		var history BrowserHistory

		for index, command := range tc.commands {
			var output any

			switch command {
			case "BrowserHistory":
				history = NewBrowserHistory(tc.values[index][0].(string))
			case "visit":
				history.Visit(tc.values[index][0].(string))
			case "back":
				output = history.Back(tc.values[index][0].(int))
			case "forward":
				output = history.Forward(tc.values[index][0].(int))
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
