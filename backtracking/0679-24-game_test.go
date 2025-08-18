package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestJudgePoint24$
func TestJudgePoint24(t *testing.T) {
	for _, tc := range []struct {
		cards []int
		can   bool
	}{
		{cards: []int{4, 1, 8, 7}, can: true},
		{cards: []int{1, 2, 1, 2}, can: false},
	} {
		can := judgePoint24(tc.cards)
		assert.Equal(t, tc.can, can)
	}
}
