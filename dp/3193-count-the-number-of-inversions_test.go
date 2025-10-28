package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfPermutations$
func TestNumberOfPermutations(t *testing.T) {
	for _, tc := range []struct {
		n            int
		requirements [][]int
		count        int
	}{
		{n: 3, requirements: [][]int{{2, 2}, {0, 0}}, count: 2},
		{n: 3, requirements: [][]int{{2, 2}, {1, 1}, {0, 0}}, count: 1},
		{n: 2, requirements: [][]int{{0, 0}, {1, 0}}, count: 1},
		{n: 3, requirements: [][]int{{2, 0}}, count: 1},
		{n: 15, requirements: [][]int{{14,58},{0,0},{10,28}}, count: 243296005},
	} {
		count := numberOfPermutations(tc.n, tc.requirements)
		assert.Equal(t, tc.count, count)
	}
}
