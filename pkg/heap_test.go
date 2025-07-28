package pkg

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./pkg/ -run ^TestMaxHeap$
func TestMaxHeap(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{0, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 11, 45},
		},
	} {
		maxHeap := NewHeap(tc.array, func(x, y int) bool { return x < y })
		heap.Init(maxHeap)

		sorted := make([]int, 0, maxHeap.Len())
		for maxHeap.Len() > 0 {
			sorted = append(sorted, heap.Pop(maxHeap).(int))
		}

		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./pkg/ -run ^TestMinHeap$
func TestMinHeap(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{45, 11, 7, 7, 6, 6, 6, 6, 5, 5, 5, 4, 4, 4, 4, 3, 2, 2, 1, 1, 0},
		},
	} {
		minHeap := NewHeap(make([]int, 0), func(x, y int) bool { return x > y })
		for _, num := range tc.array {
			heap.Push(minHeap, num)
		}

		sorted := make([]int, 0, minHeap.Len())
		for minHeap.Len() > 0 {
			sorted = append(sorted, heap.Pop(minHeap).(int))
		}

		assert.Equal(t, tc.sorted, sorted)
	}
}
