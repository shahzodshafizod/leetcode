package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestReverseWords$
func TestReverseWords(t *testing.T) {
	for _, tc := range []struct {
		s        string
		reversed string
	}{
		{s: "Let's take LeetCode contest", reversed: "s'teL ekat edoCteeL tsetnoc"},
		{s: "Mr Ding", reversed: "rM gniD"},
	} {
		reversed := reverseWords(tc.s)
		assert.Equal(t, tc.reversed, reversed)
	}
}
