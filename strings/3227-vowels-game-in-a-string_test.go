package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestDoesAliceWin$
func TestDoesAliceWin(t *testing.T) {
	for _, tc := range []struct {
		s         string
		aliceWins bool
	}{
		{s: "leetcoder", aliceWins: true},
		{s: "bbcd", aliceWins: false},
	} {
		aliceWins := doesAliceWin(tc.s)
		assert.Equal(t, tc.aliceWins, aliceWins)
	}
}
