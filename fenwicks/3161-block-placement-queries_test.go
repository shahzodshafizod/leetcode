package fenwicks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./fenwicks/ -run ^TestGetResults$
func TestGetResults(t *testing.T) {
	for _, tc := range []struct {
		queries [][]int
		results []bool
	}{
		{queries: [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}}, results: []bool{false, true, true}},
		{queries: [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}, results: []bool{true, true, false}},
	} {
		results := getResults(tc.queries)
		assert.Equal(t, tc.results, results)
	}
}
