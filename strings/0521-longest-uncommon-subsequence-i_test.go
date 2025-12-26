package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestFindLUSlength$
func TestFindLUSlength(t *testing.T) {
	for _, tc := range []struct {
		a      string
		b      string
		length int
	}{
		{a: "aba", b: "cdc", length: 3},
		{a: "aaa", b: "bbb", length: 3},
		{a: "aaa", b: "aaa", length: -1},
		{a: "abc", b: "ab", length: 3},
		{a: "a", b: "b", length: 1},
	} {
		length := findLUSlength(tc.a, tc.b)
		assert.Equal(t, tc.length, length)
	}
}
