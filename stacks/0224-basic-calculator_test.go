package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestCalculate$
func TestCalculate(t *testing.T) {
	for _, tc := range []struct {
		s      string
		result int
	}{
		{s: "1 + 1", result: 2},
		{s: "-5 + 8", result: 3},
		{s: " 2-1 + 2 ", result: 3},
		{s: "-2+ 1", result: -1},
		{s: "(1+(4+5+2)-3)+(6+8)", result: 23},
		{s: "- (3 + (4 + 5))", result: -12},
		{s: "1-( -2)", result: 3},
		{s: "2147483647", result: 2147483647},
		{s: "()", result: 0},
		{s: "()+()", result: 0},
		{s: "1+ 1", result: 2},
		{s: "1-11", result: -10},
		{s: "(1+7 7)", result: 15},
		{s: "(1 2 + 4) + 1", result: 8},
		{s: "1  2 + 4", result: 7},
	} {
		result := calculate(tc.s)
		assert.Equal(t, tc.result, result)
	}
}
