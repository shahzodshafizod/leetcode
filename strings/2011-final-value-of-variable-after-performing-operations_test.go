package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestFinalValueAfterOperations$
func TestFinalValueAfterOperations(t *testing.T) {
	for _, tc := range []struct {
		operations []string
		final      int
	}{
		{operations: []string{"--X", "X++", "X++"}, final: 1},
		{operations: []string{"++X", "++X", "X++"}, final: 3},
		{operations: []string{"X++", "++X", "--X", "X--"}, final: 0},
	} {
		final := finalValueAfterOperations(tc.operations)
		assert.Equal(t, tc.final, final)
	}
}
