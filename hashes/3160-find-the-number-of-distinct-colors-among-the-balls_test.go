package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestQueryResults$
func TestQueryResults(t *testing.T) {
	for _, tc := range []struct {
		limit   int
		queries [][]int
		result  []int
	}{
		{limit: 4, queries: [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}, result: []int{1, 2, 2, 3}},
		{limit: 4, queries: [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}, result: []int{1, 2, 2, 3, 4}},
		{limit: 2, queries: [][]int{{0, 1}, {1, 2}, {2, 3}}, result: []int{1, 2, 3}},
		{limit: 1, queries: [][]int{{0, 1}, {0, 4}, {0, 4}, {0, 1}, {1, 2}}, result: []int{1, 1, 1, 1, 2}},
		{limit: 1, queries: [][]int{{0, 2}, {1, 10}, {0, 10}, {0, 3}, {1, 5}}, result: []int{1, 2, 1, 2, 2}},
		{limit: 5, queries: [][]int{{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 4}}, result: []int{1, 1, 2, 3, 4}},
		{limit: 10, queries: [][]int{{1, 5}, {2, 5}, {1, 5}, {3, 7}, {4, 8}}, result: []int{1, 1, 1, 2, 3}},
		{limit: 6, queries: [][]int{{1, 3}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {0, 3}}, result: []int{1, 1, 2, 3, 4, 4}},
		{limit: 1000000000, queries: [][]int{{500000000, 1}, {600000000, 1}}, result: []int{1, 1}},
	} {
		result := queryResults(tc.limit, tc.queries)
		assert.Equal(t, tc.result, result)
	}
}
