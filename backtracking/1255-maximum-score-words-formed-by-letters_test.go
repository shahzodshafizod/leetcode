package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestMaxScoreWords$
func TestMaxScoreWords(t *testing.T) {
	for _, tc := range []struct {
		words    []string
		letters  []byte
		score    []int
		maxscore int
	}{
		{
			words:    []string{"dog", "cat", "dad", "good"},
			letters:  []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
			score:    []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			maxscore: 23,
		},
		{
			words:    []string{"xxxz", "ax", "bx", "cx"},
			letters:  []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
			score:    []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
			maxscore: 27,
		},
		{
			words:    []string{"leetcode"},
			letters:  []byte{'l', 'e', 't', 'c', 'o', 'd'},
			score:    []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			maxscore: 0,
		},
	} {
		maxscore := maxScoreWords(tc.words, tc.letters, tc.score)
		assert.Equal(t, tc.maxscore, maxscore)
	}
}
