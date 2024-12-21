package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestIsPowerOfTwo$
func TestIsPowerOfTwo(t *testing.T) {
	for _, tc := range []struct {
		n  int
		is bool
	}{
		{n: 1, is: true},
		{n: 16, is: true},
		{n: 3, is: false},
		{n: 4, is: true},
		{n: 0, is: false},
		{n: -16, is: false},
	} {
		is := isPowerOfTwo(tc.n)
		assert.Equal(t, tc.is, is)
	}
}
