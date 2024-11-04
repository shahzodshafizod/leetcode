package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsCircularSentence$
func TestIsCircularSentence(t *testing.T) {
	for _, tc := range []struct {
		sentence string
		is       bool
	}{
		{sentence: "L", is: true},
		{sentence: "sus", is: true},
		{sentence: "eetcode", is: true},
		{sentence: "leetcode", is: false},
		{sentence: "a b c d e a", is: false},
		{sentence: "Leetcode is cool", is: false},
		{sentence: "MuFoevIXCZzrpXeRmTssj lYSW U jM", is: false},
		{sentence: "leetcode Exercises sound delightfu", is: false},
		{sentence: "leetcode exercises sound delightful", is: true},
		{sentence: "Leetcode Exercises sound delightful", is: false},
		{sentence: "IuTiUtGGsNydmacGduehPPGksKQyT TmOraUbCcQdnZUCpGCYtGp p pG GCcRvZDRawqGKOiBSLwjIDOjdhnHiisfddYoeHqxOqkUvOEyI", is: true},
	} {
		is := isCircularSentence(tc.sentence)
		assert.Equal(t, tc.is, is)
	}
}
