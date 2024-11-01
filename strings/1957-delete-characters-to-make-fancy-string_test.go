package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMakeFancyString$
func TestMakeFancyString(t *testing.T) {
	for _, tc := range []struct {
		s     string
		fancy string
	}{
		{s: "a", fancy: "a"},
		{s: "aab", fancy: "aab"},
		{s: "aaabaaaa", fancy: "aabaa"},
		{s: "leeetcode", fancy: "leetcode"},
		{s: "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww", fancy: "ww"},
		{s: "leeetcodeeeeeeeeeeeeeeeeeeeeeedddddddddddddddddddddd", fancy: "leetcodeedd"},
	} {
		fancy := makeFancyString(tc.s)
		assert.Equal(t, tc.fancy, fancy)
	}
}
