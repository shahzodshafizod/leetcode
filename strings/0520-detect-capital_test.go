package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestDetectCapitalUse$
func TestDetectCapitalUse(t *testing.T) {
	for _, tc := range []struct {
		word    string
		correct bool
	}{
		{word: "USA", correct: true},
		{word: "uSA", correct: false},
		{word: "FlaG", correct: false},
		{word: "LeetCode", correct: false},
		{word: "g", correct: true},
	} {
		correct := detectCapitalUse(tc.word)
		assert.Equal(t, tc.correct, correct)
	}
}
