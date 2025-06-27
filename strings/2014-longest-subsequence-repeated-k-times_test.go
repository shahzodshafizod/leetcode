package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLongestSubsequenceRepeatedK$
func TestLongestSubsequenceRepeatedK(t *testing.T) {
	for _, tc := range []struct {
		s      string
		k      int
		subseq string
	}{
		{s: "bb", k: 2, subseq: "b"},
		{s: "ab", k: 2, subseq: ""},
		{s: "gbjbjigjbji", k: 2, subseq: "gjbji"},
		{s: "letsleetcode", k: 2, subseq: "let"},
		{s: "nhsbbeonhsbbeonhsbbeo", k: 3, subseq: "nhsbbeo"},
	} {
		subseq := longestSubsequenceRepeatedK(tc.s, tc.k)
		assert.Equal(t, tc.subseq, subseq)
	}
}
