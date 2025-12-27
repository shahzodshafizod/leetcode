package graphs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCriticalConnections$
func TestCriticalConnections(t *testing.T) {
	for _, tc := range []struct {
		n           int
		connections [][]int
		critical    [][]int
	}{
		{n: 4, connections: [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, critical: [][]int{{1, 3}}},
		{n: 2, connections: [][]int{{0, 1}}, critical: [][]int{{0, 1}}},
	} {
		critical := criticalConnections(tc.n, tc.connections)
		sort.Slice(critical, func(i, j int) bool {
			if critical[i][0] != critical[j][0] {
				return critical[i][0] < critical[j][0]
			}

			return critical[i][1] < critical[j][1]
		})
		assert.Equal(t, tc.critical, critical)
	}
}
