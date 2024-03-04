package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestBagOfTokensScore$
func TestBagOfTokensScore(t *testing.T) {
	for _, tc := range []struct {
		tokens []int
		power  int
		score  int
	}{
		{tokens: []int{100}, power: 50, score: 0},
		{tokens: []int{200, 100}, power: 150, score: 1},
		{tokens: []int{100, 200, 300, 400}, power: 200, score: 2},
		{tokens: []int{1, 4, 1, 1}, power: 1, score: 2},
	} {
		score := bagOfTokensScore(tc.tokens, tc.power)
		assert.Equal(t, tc.score, score)
	}
}
