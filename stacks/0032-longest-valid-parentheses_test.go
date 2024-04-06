package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestLongestValidParentheses$
func TestLongestValidParentheses(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "(()", length: 2},
		{s: "(())", length: 4},
		{s: "()()", length: 4},
		{s: ")()())", length: 4},
		{s: "", length: 0},
		{s: "()(()", length: 2},
		{s: "(()())", length: 6},
		{s: "((()(())))", length: 10},
		{s: "()(())", length: 6},
		{s: "))", length: 0},
		{s: "()))", length: 2},
		{s: "()(()))()()()", length: 6},
		{s: ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())", length: 132},
	} {
		length := longestValidParentheses(tc.s)
		assert.Equal(t, tc.length, length)
	}
}
