package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestIsUgly$
func TestIsUgly(t *testing.T) {
	for _, tc := range []struct {
		n      int
		output bool
	}{
		{n: 6, output: true},   // 6 = 2 * 3
		{n: 1, output: true},   // 1 is ugly
		{n: 14, output: false}, // 14 = 2 * 7
		{n: 0, output: false},  // non-positive
		{n: -6, output: false}, // negative
	} {
		output := isUgly(tc.n)
		assert.Equal(t, tc.output, output)
	}
}
