package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestFindWords$
func TestFindWords(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		words []string
		found []string
	}{
		{board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, words: []string{"oath", "pea", "eat", "rain"}, found: []string{"oath", "eat"}},
		{board: [][]byte{{'a', 'b'}, {'c', 'd'}}, words: []string{"abcb"}, found: []string{}},
		{board: [][]byte{{'a', 'b'}, {'c', 'd'}}, words: []string{"abcd"}, found: []string{}},
		{board: [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}, words: []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}, found: []string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"}},
		{board: [][]byte{{'a', 'a'}}, words: []string{"aaa"}, found: []string{}},
		{board: [][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}}, words: []string{"oa", "oaa"}, found: []string{"oa", "oaa"}},
		{board: [][]byte{{'a', 'c'}, {'p', 'e'}}, words: []string{"app", "ape", "ace"}, found: []string{"ape", "ace"}},
		{board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, words: []string{"oath", "pea", "eat", "rain", "hklf", "hf"}, found: []string{"oath", "eat", "hf", "hklf"}},
		{board: [][]byte{{'a'}}, words: []string{"a", "a"}, found: []string{"a"}},
	} {
		found := findWords(tc.board, tc.words)
		assert.Equal(t, tc.found, found)
	}
}
