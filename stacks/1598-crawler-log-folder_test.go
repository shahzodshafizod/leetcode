package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		logs  []string
		depth int
	}{
		{logs: []string{"d1/", "d2/", "../", "d21/", "./"}, depth: 2},
		{logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, depth: 3},
		{logs: []string{"d1/", "../", "../", "../"}, depth: 0},
	} {
		depth := minOperations(tc.logs)
		assert.Equal(t, tc.depth, depth)
	}
}
