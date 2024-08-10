package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./math/ -run ^TestIsPalindrome$
func TestIsPalindrome(t *testing.T) {
	for _, tc := range []struct {
		x          int
		palindrome bool
	}{
		{x: 121, palindrome: true},
		{x: -121, palindrome: false},
		{x: 10, palindrome: false},
		{x: 100, palindrome: false},
		{x: 1, palindrome: true},
	} {
		palindrome := isPalindrome(tc.x)
		assert.Equal(t, tc.palindrome, palindrome)
	}
}
