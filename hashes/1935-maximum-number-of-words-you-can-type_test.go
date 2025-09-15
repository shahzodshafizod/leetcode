package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCanBeTypedWords$
func TestCanBeTypedWords(t *testing.T) {
	for _, tc := range []struct {
		text          string
		brokenLetters string
		count         int
	}{
		{text: "hello world", brokenLetters: "ad", count: 1},
		{text: "leet code", brokenLetters: "lt", count: 1},
		{text: "leet code", brokenLetters: "e", count: 0},
	} {
		count := canBeTypedWords(tc.text, tc.brokenLetters)
		assert.Equal(t, tc.count, count)
	}
}
