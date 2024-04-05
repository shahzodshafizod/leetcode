package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMaxDepth$
func TestMaxDepth(t *testing.T) {
	for _, tc := range []struct {
		s     string
		depth int
	}{
		{s: "(1+(2*3)+((8)/4))+1", depth: 3},
		{s: "(1)+((2))+(((3)))", depth: 3},
		{s: "", depth: 0},
		{s: "()()", depth: 1},
		{s: "()(()())", depth: 2},
	} {
		depth := maxDepth(tc.s)
		assert.Equal(t, tc.depth, depth)
	}
}
