package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindTheCity$
func TestFindTheCity(t *testing.T) {
	for _, tc := range []struct {
		n                 int
		edges             [][]int
		distanceThreshold int
		city              int
	}{
		{n: 4, edges: [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, distanceThreshold: 4, city: 3},
		{n: 5, edges: [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, distanceThreshold: 2, city: 0},
		{n: 6, edges: [][]int{{0, 1, 10}, {0, 2, 1}, {2, 3, 1}, {1, 3, 1}, {1, 4, 1}, {4, 5, 10}}, distanceThreshold: 20, city: 5},
	} {
		city := findTheCity(tc.n, tc.edges, tc.distanceThreshold)
		assert.Equal(t, tc.city, city)
	}
}
