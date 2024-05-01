package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestReversePrefix$
func TestReversePrefix(t *testing.T) {
	for _, tc := range []struct {
		word   string
		ch     byte
		result string
	}{
		{word: "abcdefd", ch: 'd', result: "dcbaefd"},
		{word: "xyxzxe", ch: 'z', result: "zxyxxe"},
		{word: "abcd", ch: 'z', result: "abcd"},
		{word: "a", ch: 'a', result: "a"},
	} {
		result := reversePrefix(tc.word, tc.ch)
		assert.Equal(t, tc.result, result)
	}
}
