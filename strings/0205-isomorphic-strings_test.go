package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsIsomorphic$
func TestIsIsomorphic(t *testing.T) {
	for _, tc := range []struct {
		s  string
		t  string
		is bool
	}{
		{s: "egg", t: "add", is: true},
		{s: "foo", t: "bar", is: false},
		{s: "paper", t: "title", is: true},
		{s: "bbbaaaba", t: "aaabbbba", is: false},
		{s: "badc", t: "baba", is: false},
		{s: "12", t: "\u0067\u0067", is: false},
		{s: "12", t: "\u0067\u0068", is: true},
		{s: "a", t: "a", is: true},
		{s: "12", t: "gg", is: false},
	} {
		is := isIsomorphic(tc.s, tc.t)
		assert.Equal(t, tc.is, is)
	}
}
