package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestExist$
func TestExist(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		word  string
		hast  bool
	}{
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCCED", hast: true},
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "SEE", hast: true},
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCB", hast: false},
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCEFSADEESE", hast: true},
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCESEEEFS", hast: true},
	} {
		hast := exist(tc.board, tc.word)
		assert.Equal(t, tc.hast, hast)
	}
}
