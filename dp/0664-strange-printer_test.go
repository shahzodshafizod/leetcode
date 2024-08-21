package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestStrangePrinter$
func TestStrangePrinter(t *testing.T) {
	for _, tc := range []struct {
		s     string
		turns int
	}{
		{s: "aaabbb", turns: 2},
		{s: "aba", turns: 2},
		{s: "abcab", turns: 4},
		{s: "abcabc", turns: 5},
		{s: "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", turns: 1},
		{s: "abcabccb", turns: 5},
		{s: "abcdbca", turns: 5},
		{s: "abbbccbbbbccccbbbbbccccc", turns: 5},
		{s: "cabad", turns: 4},
		{s: "ababa", turns: 3},
	} {
		turns := strangePrinter(tc.s)
		assert.Equal(t, tc.turns, turns)
	}
}
