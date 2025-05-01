package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaxTaskAssign$
func TestMaxTaskAssign(t *testing.T) {
	for _, tc := range []struct {
		tasks    []int
		workers  []int
		pills    int
		strength int
		count    int
	}{
		{tasks: []int{3, 2, 1}, workers: []int{0, 3, 3}, pills: 1, strength: 1, count: 3},
		{tasks: []int{5, 4}, workers: []int{0, 0, 0}, pills: 1, strength: 5, count: 1},
		{tasks: []int{10, 15, 30}, workers: []int{0, 10, 10, 10, 10}, pills: 3, strength: 10, count: 2},
		{tasks: []int{2, 5}, workers: []int{1, 3, 4}, pills: 1, strength: 1, count: 2},
		{tasks: []int{5, 9, 8, 5, 9}, workers: []int{1, 6, 4, 2, 6}, pills: 1, strength: 5, count: 3},
		{tasks: []int{10, 10, 10}, workers: []int{0, 0, 0}, pills: 3, strength: 10, count: 3},
		{tasks: []int{8, 10}, workers: []int{4}, pills: 0, strength: 6, count: 0},
	} {
		count := maxTaskAssign(tc.tasks, tc.workers, tc.pills, tc.strength)
		assert.Equal(t, tc.count, count)
	}
}
