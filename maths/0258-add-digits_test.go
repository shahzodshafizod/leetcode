package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestAddDigits$
func TestAddDigits(t *testing.T) {
	for _, tc := range []struct {
		num   int
		digit int
	}{
		{num: 38, digit: 2},
		{num: 0, digit: 0},
		{num: 121, digit: 4},
		{num: 10, digit: 1},
	} {
		digit := addDigits(tc.num)
		assert.Equal(t, tc.digit, digit)
	}
}
