package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestSpellchecker$
func TestSpellchecker(t *testing.T) {
	for _, tc := range []struct {
		wordlist []string
		queries  []string
		answer   []string
	}{
		{wordlist: []string{"KiTe", "kite", "hare", "Hare"}, queries: []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}, answer: []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}},
		{wordlist: []string{"yellow"}, queries: []string{"YellOw"}, answer: []string{"yellow"}},
	} {
		answer := spellchecker(tc.wordlist, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
