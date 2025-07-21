package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxEvents$
func TestMaxEvents(t *testing.T) {
	for _, tc := range []struct {
		events [][]int
		count  int
	}{
		{events: [][]int{{1, 10}, {1, 10}, {1, 10}, {1, 10}}, count: 4},
		{events: [][]int{{1, 100}, {50, 100}, {60, 70}, {90, 95}}, count: 4},
		{events: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}, count: 4},
		{events: [][]int{{1, 2}, {2, 2}, {1, 2}, {1, 2}}, count: 2},
		{events: [][]int{{1, 2}, {3, 4}, {5, 6}}, count: 3},
		{events: [][]int{{1, 2}, {4, 5}, {7, 8}, {10, 11}}, count: 4},
		{events: [][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}, count: 1},
		{events: [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}}, count: 5},
		{events: [][]int{{1, 3}, {1, 3}, {1, 3}, {1, 3}}, count: 3},
		{events: [][]int{{1, 3}, {5, 8}, {6, 10}, {9, 11}, {12, 15}}, count: 5},
		{events: [][]int{{1, 2}, {2, 3}, {3, 4}}, count: 3},
		{events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}, count: 4},
		{events: [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}, count: 4},
		{events: [][]int{{1, 100000}}, count: 1},
		{events: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}, count: 7},
		{events: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, count: 5},
		{events: [][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}, count: 5},
		{events: [][]int{{1, 10}, {2, 10}, {3, 10}, {4, 10}}, count: 4},
		{events: [][]int{{1, 2}, {1, 2}, {1, 6}, {1, 2}, {1, 2}}, count: 3},
		{events: [][]int{{1, 10}, {2, 2}, {2, 2}, {2, 2}, {2, 2}}, count: 2},
		{events: [][]int{{1, 2}, {2, 2}, {3, 3}, {3, 4}, {3, 4}}, count: 4},
		{events: [][]int{{1, 4}, {1, 1}, {2, 2}, {3, 4}, {4, 4}}, count: 4},
		{events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}, {2, 5}}, count: 5},
		{events: [][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}, count: 5},
	} {
		count := maxEvents(tc.events)
		assert.Equal(t, tc.count, count)
	}
}
