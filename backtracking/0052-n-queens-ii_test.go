package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestTotalNQueens$
func TestTotalNQueens(t *testing.T) {
	for _, tc := range []struct {
		n         int
		solutions int
	}{
		{n: 4, solutions: 2},
		{n: 1, solutions: 1},
	} {
		solutions := totalNQueens(tc.n)
		assert.Equal(t, tc.solutions, solutions)
	}
}
