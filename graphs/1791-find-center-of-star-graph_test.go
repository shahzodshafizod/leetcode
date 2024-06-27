package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindCenter$
func TestFindCenter(t *testing.T) {
	for _, tc := range []struct {
		edges  [][]int
		center int
	}{
		{edges: [][]int{{1, 2}, {2, 3}, {4, 2}}, center: 2},
		{edges: [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, center: 1},
	} {
		center := findCenter(tc.edges)
		assert.Equal(t, tc.center, center)
	}
}
