package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestJudgeCircle$
func TestJudgeCircle(t *testing.T) {
	for _, tc := range []struct {
		moves   string
		returns bool
	}{
		{moves: "UD", returns: true},
		{moves: "LL", returns: false},
	} {
		returns := judgeCircle(tc.moves)
		assert.Equal(t, tc.returns, returns)
	}
}
