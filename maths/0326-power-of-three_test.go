package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestIsPowerOfThree$
func TestIsPowerOfThree(t *testing.T) {
	for _, tc := range []struct {
		n  int
		is bool
	}{
		{n: 27, is: true},
		{n: 0, is: false},
		{n: -1, is: false},
	} {
		is := isPowerOfThree(tc.n)
		assert.Equal(t, tc.is, is)
	}
}
