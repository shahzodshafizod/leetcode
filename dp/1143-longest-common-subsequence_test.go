package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLongestCommonSubsequence$
func TestLongestCommonSubsequence(t *testing.T) {
	for _, tc := range []struct {
		text1  string
		text2  string
		length int
	}{
		{text1: "abcde", text2: "ace", length: 3},
		{text1: "abc", text2: "abc", length: 3},
		{text1: "abc", text2: "def", length: 0},
		{text1: "shahzod", text2: "shafizod", length: 6},
		{text1: "mhunuzqrkzsnidwbun", text2: "szulspmhwpazoxijwbq", length: 6},
		{text1: "abcba", text2: "abcbcba", length: 5},
	} {
		length := longestCommonSubsequence(tc.text1, tc.text2)
		assert.Equal(t, tc.length, length)
	}
}
