package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestIsAnagram$
func TestIsAnagram(t *testing.T) {
	for _, tc := range []struct {
		s  string
		t  string
		is bool
	}{
		{s: "anagram", t: "nagaram", is: true},
		{s: "rat", t: "car", is: false},
		{s: "a", t: "abb", is: false}, // a clear example why bit manipulation doesn't work
	} {
		is := isAnagram(tc.s, tc.t)
		assert.Equal(t, tc.is, is)
	}
}
