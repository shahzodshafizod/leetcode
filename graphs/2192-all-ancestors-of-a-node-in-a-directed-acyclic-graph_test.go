package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestGetAncestors$
func TestGetAncestors(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		answer [][]int
	}{
		{
			n:      8,
			edges:  [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},
			answer: [][]int{nil, nil, nil, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
		},
		{
			n:      5,
			edges:  [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			answer: [][]int{nil, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
		},
	} {
		answer := getAncestors(tc.n, tc.edges)
		assert.Equal(t, tc.answer, answer)
	}
}
