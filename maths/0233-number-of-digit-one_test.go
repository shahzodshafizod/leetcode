package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountDigitOne$
func TestCountDigitOne(t *testing.T) {
	for _, tc := range []struct {
		n      int
		output int
	}{
		{n: 13, output: 6},    // 1,10,11,11,12,13 -> 6 ones
		{n: 0, output: 0},     // No numbers with 1
		{n: 1, output: 1},     // Just 1 -> 1 one
		{n: 10, output: 2},    // 1,10 -> 2 ones
		{n: 100, output: 21},  // Multiple positions with ones
		{n: 824, output: 273}, // Larger number
	} {
		output := countDigitOne(tc.n)
		assert.Equal(t, tc.output, output)
	}
}
