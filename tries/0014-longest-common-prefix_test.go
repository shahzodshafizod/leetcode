package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestLongestCommonPrefix$
func TestLongestCommonPrefix(t *testing.T) {
	for _, tc := range []struct {
		strs   []string
		prefix string
	}{
		{strs: []string{"flower", "flow", "flight"}, prefix: "fl"},
		{strs: []string{"dog", "racecar", "car"}, prefix: ""},
		{strs: []string{"reflower", "flow", "flight"}, prefix: ""},
		{strs: []string{"", ""}, prefix: ""},
		{strs: []string{"aaaa", "aa", "aaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa"}, prefix: "aa"},
		{strs: []string{"flower", "fkow"}, prefix: "f"},
		{strs: []string{"spoiler", "flower", "flow", "flight"}, prefix: ""},
		{strs: []string{"c", "acc", "ccc"}, prefix: ""},
	} {
		prefix := longestCommonPrefix(tc.strs)
		assert.Equal(t, tc.prefix, prefix)
	}
}
