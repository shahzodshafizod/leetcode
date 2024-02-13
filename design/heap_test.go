package design

import (
	"container/heap"
	"testing"
)

// go test -v -count=1 ./design/ -run ^TestMaxHeap$
func TestMaxHeap(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var maxHeap = NewHeap(make([]int, 0), func(x, y int) bool { return x < y })
	for _, elem := range array {
		heap.Push(maxHeap, elem)
	}
	var sorted = make([]int, 0, maxHeap.Len())
	for maxHeap.Len() > 0 {
		sorted = append(sorted, heap.Pop(maxHeap).(int))
	}
	t.Log("array:", array, len(array))
	t.Log("sorted:", sorted, len(sorted))
}

// go test -v -count=1 ./design/ -run ^TestMinHeap$
func TestMinHeap(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var minHeap = NewHeap(make([]int, 0), func(x, y int) bool { return x > y })
	for _, elem := range array {
		heap.Push(minHeap, elem)
	}
	var sorted = make([]int, 0, minHeap.Len())
	for minHeap.Len() > 0 {
		sorted = append(sorted, heap.Pop(minHeap).(int))
	}
	t.Log("array:", array, len(array))
	t.Log("sorted:", sorted, len(sorted))
}
