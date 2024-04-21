package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestValidPath$
func TestValidPath(t *testing.T) {
	for _, tc := range []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		valid       bool
	}{
		{n: 3, edges: [][]int{{0, 1}, {1, 2}, {2, 0}}, source: 0, destination: 2, valid: true},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, source: 0, destination: 5, valid: false},
	} {
		valid := validPath(tc.n, tc.edges, tc.source, tc.destination)
		assert.Equal(t, tc.valid, valid)
	}
}
