package design

import (
	"fmt"
	"testing"
)

// go test -v -count=1 ./design/ -run ^TestMaxPriorityQueue$
func TestMaxPriorityQueue(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var maxHeap = NewPriorityQueue(func(item1, item2 int) bool { return item1 < item2 })
	for _, elem := range array {
		maxHeap.Push(elem)
	}
	var sorted = make([]int, 0, maxHeap.Size())
	for !maxHeap.IsEmpty() {
		sorted = append(sorted, maxHeap.Pop())
	}
	fmt.Println("array:", array, len(array))
	fmt.Println("sorted:", sorted, len(sorted))
}

// go test -v -count=1 ./design/ -run ^TestMinPriorityQueue$
func TestMinPriorityQueue(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var minHeap = NewPriorityQueue(func(item1, item2 int) bool { return item1 > item2 })
	for _, elem := range array {
		minHeap.Push(elem)
	}
	var sorted = make([]int, 0, minHeap.Size())
	for !minHeap.IsEmpty() {
		sorted = append(sorted, minHeap.Pop())
	}
	fmt.Println("array:", array, len(array))
	fmt.Println("sorted:", sorted, len(sorted))
}
