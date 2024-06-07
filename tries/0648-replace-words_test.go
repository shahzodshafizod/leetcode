package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestReplaceWords$
func TestReplaceWords(t *testing.T) {
	for _, tc := range []struct {
		dictionary []string
		sentence   string
		replaced   string
	}{
		{
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			replaced:   "the cat was rat by the bat",
		},
		{
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			replaced:   "a a b c",
		},
		{
			dictionary: []string{"a", "aa", "aaa", "aaaa"},
			sentence:   "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
			replaced:   "a a a a a a a a bbb baba a",
		},
		{
			dictionary: []string{"catt", "cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			replaced:   "the cat was rat by the bat",
		},
	} {
		replaced := replaceWords(tc.dictionary, tc.sentence)
		assert.Equal(t, tc.replaced, replaced)
	}
}
