package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestUniquePaths$
func TestUniquePaths(t *testing.T) {
	for _, tc := range []struct {
		m         int
		n         int
		pathCount int
	}{
		{m: 3, n: 7, pathCount: 28},
		{m: 4, n: 7, pathCount: 84},
		{m: 3, n: 2, pathCount: 3},
		{m: 1, n: 1, pathCount: 1},
	} {
		pathCount := uniquePaths(tc.m, tc.n)
		assert.Equal(t, tc.pathCount, pathCount)
	}
}
