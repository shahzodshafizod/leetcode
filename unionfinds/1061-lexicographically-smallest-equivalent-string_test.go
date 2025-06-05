package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestSmallestEquivalentString$
func TestSmallestEquivalentString(t *testing.T) {
	for _, tc := range []struct {
		s1         string
		s2         string
		baseStr    string
		equivalent string
	}{
		{s1: "parker", s2: "morris", baseStr: "parser", equivalent: "makkek"},
		{s1: "hello", s2: "world", baseStr: "hold", equivalent: "hdld"},
		{s1: "leetcode", s2: "programs", baseStr: "sourcecode", equivalent: "aauaaaaada"},
	} {
		equivalent := smallestEquivalentString(tc.s1, tc.s2, tc.baseStr)
		assert.Equal(t, tc.equivalent, equivalent)
	}
}
