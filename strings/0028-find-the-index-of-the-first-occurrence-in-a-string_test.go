package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestStrStr$
func TestStrStr(t *testing.T) {
	for _, tc := range []struct {
		haystack string
		needle   string
		index    int
	}{
		{haystack: "sadbutsad", needle: "sad", index: 0},
		{haystack: "leetcode", needle: "leeto", index: -1},
		{haystack: "sydbutsad", needle: "sad", index: 6},
		{haystack: "aabaaabaaac", needle: "aabaaac", index: 4},
		{haystack: "ababcaababcaabc", needle: "ababcaabc", index: 6},
		// {haystack: "fourscoreandsevenyearsagoourfathersbroughtforthuponthiscontinentanewnation", needle: "ourfathersbroughtforthuponthiscontinentanewnation", index: 25},
	} {
		index := strStr(tc.haystack, tc.needle)
		assert.Equal(t, tc.index, index)
	}
}
