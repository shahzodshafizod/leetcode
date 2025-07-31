package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestGuessNumber$
func TestGuessNumber(t *testing.T) {
	for _, tc := range []struct {
		n       int
		picked  int
		guessed int
	}{
		{n: 10, picked: 6, guessed: 6},
		{n: 1, picked: 1, guessed: 1},
		{n: 2, picked: 1, guessed: 1},
	} {
		GuessedNumber = tc.picked
		guessed := guessNumber(tc.n)
		assert.Equal(t, tc.guessed, guessed)
	}
}
