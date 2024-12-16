package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestIsPowerOfFour$
func TestIsPowerOfFour(t *testing.T) {
	for _, tc := range []struct {
		n  int
		is bool
	}{
		{n: 16, is: true},
		{n: 5, is: false},
		{n: 1, is: true},
		{n: 512, is: false},
		{n: 1024, is: true},
		{n: 2048, is: false},
		// {n: 1073741824, is: true},
	} {
		is := isPowerOfFour(tc.n)
		assert.Equal(t, tc.is, is)
	}
}
