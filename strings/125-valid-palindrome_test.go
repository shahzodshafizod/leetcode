package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsPalindrome$
func TestIsPalindrome(t *testing.T) {
	for _, data := range []struct {
		s  string
		is bool
	}{
		{s: "aabaa", is: true},
		{s: "aabbaa", is: true},
		{s: "abc", is: false},
		{s: "a", is: true},
		{s: "", is: true},
		{s: "aba", is: true},
		{s: "race car", is: true},
		{s: "A man, a plan, a canal: Panama", is: true},
		{s: "race a car", is: false},
		{s: " ", is: true},
		{s: ".,", is: true},
		{s: "ab_a", is: true},
	} {
		is := isPalindrome(data.s)
		assert.Equal(t, data.is, is)
	}
}
