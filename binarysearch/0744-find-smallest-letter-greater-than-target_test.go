package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestNextGreatestLetter$
func TestNextGreatestLetter(t *testing.T) {
	for _, tc := range []struct {
		letters  []byte
		target   byte
		expected byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
		{[]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'e', 'n'},
		{[]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'n', 'e'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
	} {
		output := nextGreatestLetter(tc.letters, tc.target)
		assert.Equal(t, tc.expected, output)
	}
}
