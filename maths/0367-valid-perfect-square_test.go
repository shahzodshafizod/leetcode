package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestIsPerfectSquare$
func TestIsPerfectSquare(t *testing.T) {
	for _, tc := range []struct {
		num      int
		isSquare bool
	}{
		{num: 16, isSquare: true},
		{num: 14, isSquare: false},
		{num: 2147483647, isSquare: false},
		{num: 1, isSquare: true},
		{num: 2, isSquare: false},
		{num: 3, isSquare: false},
		{num: 4, isSquare: true},
	} {
		isSquare := isPerfectSquare(tc.num)
		assert.Equal(t, tc.isSquare, isSquare)
	}
}
