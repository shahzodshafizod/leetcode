package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestReverseParentheses$
func TestReverseParentheses(t *testing.T) {
	for _, tc := range []struct {
		s      string
		result string
	}{
		{s: "(abcd)", result: "dcba"},
		{s: "(u(love)i)", result: "iloveu"},
		{s: "(ed(et(oc))el)", result: "leetcode"},
	} {
		result := reverseParentheses(tc.s)
		assert.Equal(t, tc.result, result)
	}
}
