package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestParseBoolExpr$
func TestParseBoolExpr(t *testing.T) {
	for _, tc := range []struct {
		expression string
		result     bool
	}{
		{expression: "&(|(f))", result: false},
		{expression: "|(f,f,f,t)", result: true},
		{expression: "!(&(f,t))", result: true},
		{expression: "!(&(!(t),&(f),|(f)))", result: true},
		{expression: "!(&(!(&(f)),&(t),|(f,f,t)))", result: false},
		{expression: "|(&(t,f,f,t,t,t,t,f,f,t),t,f)", result: true},
		{expression: "|(&(t,f,t,!(&(|(f,f,f,f)))),t)", result: true},
		{expression: "!(&(&(!(&(f)),&(t),|(f,f,t)),&(t),&(t,t,f)))", result: true},
	} {
		result := parseBoolExpr(tc.expression)
		assert.Equal(t, tc.result, result)
	}
}
