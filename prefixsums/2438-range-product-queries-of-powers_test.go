package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestProductQueries$
func TestProductQueries(t *testing.T) {
	for _, tc := range []struct {
		n int
		queries [][]int
		answers []int
	}{
		{n: 15, queries: [][]int{{0, 1}, {2, 2}, {0, 3}}, answers: []int{2, 4, 64}},
		{n: 2, queries: [][]int{{0, 0}}, answers: []int{2}},
	} {
		answers := productQueries(tc.n, tc.queries)
		assert.Equal(t, tc.answers, answers)
	}
}
