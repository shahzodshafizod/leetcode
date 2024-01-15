package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMinRemoveToMakeValid$
func TestMinRemoveToMakeValid(t *testing.T) {
	for _, data := range []struct {
		s      string
		validS string
	}{
		{s: "(ab(c)d", validS: "ab(c)d"},
		{s: "lee(t(c)o)de)", validS: "lee(t(c)o)de"},
		{s: "))((", validS: ""},
		{s: "a)bc(d)", validS: "abc(d)"},
		{s: "a)b(c)d", validS: "ab(c)d"},
		{s: "((()))", validS: "((()))"},
		{s: ")))(((", validS: ""},
		{s: "leetcode beats", validS: "leetcode beats"},
		{s: "", validS: ""},
		{s: "())()(((", validS: "()()"},
		{s: "(t(e)y))d(()(e(", validS: "(t(e)y)d()e"},
		{s: "())fwk)))())(())))())()()((())", validS: "()fwk()(())()()()(())"},
	} {
		validS := minRemoveToMakeValid(data.s)
		assert.Equal(t, data.validS, validS)
	}
}
