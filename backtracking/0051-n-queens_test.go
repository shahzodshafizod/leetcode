package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestSolveNQueens$
func TestSolveNQueens(t *testing.T) {
	for _, tc := range []struct {
		n         int
		solutions [][]string
	}{
		{n: 4, solutions: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{n: 1, solutions: [][]string{{"Q"}}},
	} {
		solutions := solveNQueens(tc.n)
		assert.Equal(t, tc.solutions, solutions)
	}
}
