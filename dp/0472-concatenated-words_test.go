package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestFindAllConcatenatedWordsInADict$
func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	for _, tc := range []struct {
		words        []string
		concatenated []string
	}{
		{words: []string{"cat", "dog", "catdog"}, concatenated: []string{"catdog"}},
		{words: []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, concatenated: []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
	} {
		concatenated := findAllConcatenatedWordsInADict(tc.words)
		assert.ElementsMatch(t, tc.concatenated, concatenated)
	}
}
