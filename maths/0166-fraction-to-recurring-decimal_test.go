package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestFractionToDecimal$
func TestFractionToDecimal(t *testing.T) {
	for _, tc := range []struct {
		numerator   int
		denominator int
		result      string
	}{
		{numerator: 1, denominator: 2, result: "0.5"},
		{numerator: 2, denominator: 1, result: "2"},
		{numerator: 4, denominator: 333, result: "0.(012)"},
		{numerator: -50, denominator: 8, result: "-6.25"},
		{numerator: 0, denominator: -5, result: "0"},
		{numerator: -1, denominator: -2147483648, result: "0.0000000004656612873077392578125"},
		{numerator: 1, denominator: 6, result: "0.1(6)"},
	} {
		result := fractionToDecimal(tc.numerator, tc.denominator)
		assert.Equal(t, tc.result, result)
	}
}
