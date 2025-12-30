package backtracking

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestRemoveInvalidParentheses$
func TestRemoveInvalidParentheses(t *testing.T) {
	for _, tc := range []struct {
		s        string
		expected []string
	}{
		{"()())()", []string{"(())()", "()()()"}},
		{"(a)())()", []string{"(a())()", "(a)()()"}},
		{")(", []string{""}},
		{"", []string{""}},
		{"()", []string{"()"}},
		{"()())", []string{"(())", "()()"}},
		{"())", []string{"()"}},
		{"(()", []string{"()"}},
	} {
		output := removeInvalidParentheses(tc.s)
		sort.Strings(output)
		sort.Strings(tc.expected)
		assert.Equal(t, tc.expected, output)
	}
}
