package slidingwindows

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestFindSubstring$
func TestFindSubstring(t *testing.T) {
	for _, tc := range []struct {
		s          string
		words      []string
		substrings []int
	}{
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}, substrings: []int{0, 9}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}, substrings: []int{}},
		{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}, substrings: []int{6, 9, 12}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "good"}, substrings: []int{8}},
		{s: "lingmindraboofooowingdingbarrwingmonkeypoundcake", words: []string{"fooo", "barr", "wing", "ding", "wing"}, substrings: []int{13}},
		{s: "arrarra", words: []string{"ar", "rr", "ra"}, substrings: []int{0, 1}},
		{s: "aaaaaaaaaaaaaa", words: []string{"aa", "aa"}, substrings: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{s: "ooooooooooooooooooooooooooooooo", words: []string{"oo", "oo"}, substrings: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}},
		{s: "aaaaa", words: []string{"aa", "aa"}, substrings: []int{0, 1}},
		{s: "aaa", words: []string{"a", "a"}, substrings: []int{0, 1}},
		{s: "Ftimerunsfast", words: []string{"time", "runs", "fast"}, substrings: []int{1}},
		{s: "barfoofoxobarfhefoobarman", words: []string{"arf", "rfo", "ofo", "xob"}, substrings: []int{2}},
		{s: "abababab", words: []string{"ab", "ab", "ab"}, substrings: []int{0, 2}},
		{s: "ababaab", words: []string{"ab", "ba", "ba"}, substrings: []int{1}},
	} {
		substrings := findSubstring(tc.s, tc.words)
		sort.Ints(substrings)
		assert.Equal(t, tc.substrings, substrings)
	}
}
