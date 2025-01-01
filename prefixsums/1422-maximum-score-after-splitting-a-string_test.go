package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMaxScore$
func TestMaxScore(t *testing.T) {
	for _, tc := range []struct {
		s     string
		score int
	}{
		{s: "00", score: 1},
		{s: "01", score: 2},
		{s: "010", score: 2},
		{s: "1111", score: 3},
		{s: "0010", score: 3},
		{s: "00111", score: 5},
		{s: "01101", score: 4},
		{s: "011101", score: 5},
		{s: "1111000", score: 3},
	} {
		score := maxScore(tc.s)
		assert.Equal(t, tc.score, score)
	}
}
