package design

import (
	"fmt"
	"testing"
)

// go test -v -count=1 ./design/ -run ^TestMaxHeap$
func TestMaxHeap(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var maxHeap = NewPriorityQueue(func(a, b int) bool { return a < b })
	for _, elem := range array {
		if !maxHeap.Has(elem) {
			maxHeap.Push(elem)
		}
	}
	var sorted = make([]int, 0, maxHeap.Size())
	for !maxHeap.IsEmpty() {
		sorted = append(sorted, maxHeap.Pop())
	}
	fmt.Println("array:", array, len(array))
	fmt.Println("sorted:", sorted, len(sorted))
}

// go test -v -count=1 ./design/ -run ^TestMinHeap$
func TestMinHeap(t *testing.T) {
	var array = []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5}
	var minHeap = NewPriorityQueue(func(a, b int) bool { return a > b })
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
