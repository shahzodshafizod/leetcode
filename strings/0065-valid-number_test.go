package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsNumber$
func TestIsNumber(t *testing.T) {
	for _, tc := range []struct {
		s      string
		result bool
	}{
		// Valid numbers
		{s: "0", result: true},
		{s: "2", result: true},
		{s: "0089", result: true},
		{s: "-0.1", result: true},
		{s: "+3.14", result: true},
		{s: "4.", result: true},
		{s: "-.9", result: true},
		{s: "2e10", result: true},
		{s: "-90E3", result: true},
		{s: "3e+7", result: true},
		{s: "+6e-1", result: true},
		{s: "53.5e93", result: true},
		{s: "-123.456e789", result: true},
		// Invalid numbers
		{s: "abc", result: false},
		{s: "1a", result: false},
		{s: "1e", result: false},
		{s: "e3", result: false},
		{s: "99e2.5", result: false},
		{s: "--6", result: false},
		{s: "-+3", result: false},
		{s: "95a54e53", result: false},
		{s: "e", result: false},
		{s: ".", result: false},
		{s: "+", result: false},
		{s: "6e6.5", result: false},
		{s: ".e1", result: false},
		{s: "3.", result: true},
		{s: ".1", result: true},
	} {
		result := isNumber(tc.s)
		assert.Equal(t, tc.result, result)
	}
}
