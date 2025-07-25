package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./pkg/ -run ^TestMaxPriorityQueue$
func TestMaxPriorityQueue(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{45, 11, 7, 7, 6, 6, 6, 6, 5, 5, 5, 4, 4, 4, 4, 3, 2, 2, 1, 1, 0},
		},
	} {
		maxHeap := newPriorityQueue(tc.array, func(x, y int) bool { return x < y })
		maxHeap.Heapify()
		sorted := make([]int, 0, maxHeap.Len())
		for maxHeap.Len() > 0 {
			sorted = append(sorted, maxHeap.Pop())
		}
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./pkg/ -run ^TestMinPriorityQueue$
func TestMinPriorityQueue(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{0, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 11, 45},
		},
	} {
		minHeap := newPriorityQueue(tc.array, func(x, y int) bool { return x > y })
		minHeap.Heapify()
		sorted := make([]int, 0, minHeap.Len())
		for minHeap.Len() > 0 {
			sorted = append(sorted, minHeap.Pop())
		}
		assert.Equal(t, tc.sorted, sorted)
	}
}
