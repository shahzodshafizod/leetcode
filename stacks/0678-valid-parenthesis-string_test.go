package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestCheckValidString$
func TestCheckValidString(t *testing.T) {
	for _, tc := range []struct {
		s     string
		valid bool
	}{
		{s: "()", valid: true},
		{s: "(*)", valid: true},
		{s: "(*))", valid: true},
		{s: "(()()(())()()(())()()(())()()(())()()(())()()(())()()(())()", valid: false},
		{s: "(()()(())()()(())()()(())()()(())()()(())()()(())()()(())())", valid: true},
		{s: "((((()(()()()()((((()()(**(())))))(())()())(((())())())))))))(((((()))))()))(()((()()))()()", valid: false},
		{s: "*", valid: true},
		{s: "(((((*)))**", valid: true},
		{s: "((((()(()()()()((((()()(**(())))))(())()())(((())())())))))))(((((()))))()))(()(()()))()()", valid: true},
		{s: "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()", valid: true},
		{s: "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())", valid: false},
		{s: "(((()))())))*))())()(**(((())(()(*()((((())))*())(())*(*(()(*)))()*())**((()(()))())(*(*))*))())", valid: false},
	} {
		valid := checkValidString(tc.s)
		assert.Equal(t, tc.valid, valid)
	}
}
