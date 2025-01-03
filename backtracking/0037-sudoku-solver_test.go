package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestSolveSudoku$
func TestSolveSudoku(t *testing.T) {
	for _, tc := range []struct {
		board  [][]byte
		solved [][]byte
	}{
		{
			board:  [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			solved: [][]byte{{'5', '3', '4', '6', '7', '8', '9', '1', '2'}, {'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8', '3', '4', '2', '5', '6', '7'}, {'8', '5', '9', '7', '6', '1', '4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'}, {'7', '1', '3', '9', '2', '4', '8', '5', '6'}, {'9', '6', '1', '5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9', '6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
		{
			board:  [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '8', '.'}, {'.', '.', '.', '.', '.', '.', '1', '2', '.'}, {'.', '.', '.', '6', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '9', '.', '.', '.', '.'}, {'.', '8', '9', '.', '.', '.', '.', '4', '.'}, {'.', '.', '.', '.', '.', '6', '.', '.', '2'}, {'.', '9', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}},
			solved: [][]byte{{'1', '2', '3', '4', '5', '8', '6', '7', '9'}, {'4', '6', '7', '1', '2', '9', '3', '8', '5'}, {'9', '5', '8', '3', '6', '7', '1', '2', '4'}, {'2', '1', '4', '6', '3', '5', '7', '9', '8'}, {'3', '7', '5', '8', '9', '4', '2', '1', '6'}, {'6', '8', '9', '2', '7', '1', '5', '4', '3'}, {'5', '4', '1', '7', '8', '6', '9', '3', '2'}, {'7', '9', '2', '5', '4', '3', '8', '6', '1'}, {'8', '3', '6', '9', '1', '2', '4', '5', '7'}},
		},
		{
			board:  [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}},
			solved: [][]byte{{'1', '2', '3', '4', '5', '6', '7', '8', '9'}, {'4', '5', '6', '7', '8', '9', '1', '2', '3'}, {'7', '8', '9', '1', '2', '3', '4', '5', '6'}, {'2', '1', '4', '3', '6', '5', '8', '9', '7'}, {'3', '6', '5', '8', '9', '7', '2', '1', '4'}, {'8', '9', '7', '2', '1', '4', '3', '6', '5'}, {'5', '3', '1', '6', '4', '2', '9', '7', '8'}, {'6', '4', '2', '9', '7', '8', '5', '3', '1'}, {'9', '7', '8', '5', '3', '1', '6', '4', '2'}},
		},
	} {
		solveSudoku(tc.board)
		assert.Equal(t, tc.solved, tc.board)
	}
}
