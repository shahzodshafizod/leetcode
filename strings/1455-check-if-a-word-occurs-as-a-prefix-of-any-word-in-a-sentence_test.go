package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsPrefixOfWord$
func TestIsPrefixOfWord(t *testing.T) {
	for _, tc := range []struct {
		sentence   string
		searchWord string
		index      int
	}{
		{sentence: "love", searchWord: "lov", index: 1},
		{sentence: "i am tired", searchWord: "you", index: -1},
		{sentence: "helloh hellohe", searchWord: "hellohe", index: 2},
		{sentence: "i love eating burger", searchWord: "burg", index: 4},
		{sentence: "hello from the other side", searchWord: "they", index: -1},
		{sentence: "hellohello hellohellohello", searchWord: "ello", index: -1},
		{sentence: "i love eating broadburgers", searchWord: "burg", index: -1},
		{sentence: "this problem is an easy problem", searchWord: "pro", index: 2},
		{sentence: "love i love eating bunburger burger", searchWord: "i", index: 2},
		{sentence: "i love eating bunburger burger", searchWord: "burg", index: 5},
		// {sentence: "winstontang love they i pillow jonathan you winstontang pillow dream you duck", searchWord: "p", index: 5},
	} {
		index := isPrefixOfWord(tc.sentence, tc.searchWord)
		assert.Equal(t, tc.index, index)
	}
}
