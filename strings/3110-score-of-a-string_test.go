package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestScoreOfString$
func TestScoreOfString(t *testing.T) {
	for _, tc := range []struct {
		s     string
		score int
	}{
		{s: "hello", score: 13},
		{s: "zaz", score: 50},
	} {
		score := scoreOfString(tc.s)
		assert.Equal(t, tc.score, score)
	}
}
