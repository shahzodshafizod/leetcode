package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestHasSameDigits$
func TestHasSameDigits(t *testing.T) {
	for _, tc := range []struct {
		s      string
		output bool
	}{
		{s: "3902", output: true},
		{s: "34789", output: false},
		{s: "111", output: true},
		{s: "12", output: false},
	} {
		output := hasSameDigits(tc.s)
		assert.Equal(t, tc.output, output)
	}
}
