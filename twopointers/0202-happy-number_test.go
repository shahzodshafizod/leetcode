package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestIsHappy$
func TestIsHappy(t *testing.T) {
	for _, tc := range []struct {
		n int
		is bool
	}{
		{n: 19, is: true},
		{n: 2, is: false},
	} {
		is := isHappy(tc.n)
		assert.Equal(t, tc.is, is)
	}
}
