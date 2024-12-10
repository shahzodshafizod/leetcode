package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestReverseStr$
func TestReverseStr(t *testing.T) {
	for _, tc := range []struct {
		s        string
		k        int
		modified string
	}{
		{s: "abcd", k: 2, modified: "bacd"},
		{s: "abcdefg", k: 2, modified: "bacdfeg"},
		{s: "abcdefg", k: 8, modified: "gfedcba"},
	} {
		modified := reverseStr(tc.s, tc.k)
		assert.Equal(t, tc.modified, modified)
	}
}
