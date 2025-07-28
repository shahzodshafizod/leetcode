package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	maxHeap *pkg.Heap[int] // less that or equal to median
	minHeap *pkg.Heap[int] // greater that median
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		maxHeap: pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x > y }),
		minHeap: pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x < y }),
	}
}

func (m *MedianFinder) AddNum(num int) {
	// heap.Push(m.maxHeap, num)
	// heap.Push(m.minHeap, heap.Pop(m.maxHeap))
	// if m.minHeap.Len() > m.maxHeap.Len() {
	// 	heap.Push(m.maxHeap, heap.Pop(m.minHeap))
	// }
	if float64(num) <= m.FindMedian() {
		heap.Push(m.maxHeap, num)
	} else {
		heap.Push(m.minHeap, num)
	}
	// rebalance
	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		heap.Push(m.minHeap, heap.Pop(m.maxHeap))
	}

	if m.minHeap.Len() > m.maxHeap.Len() {
		heap.Push(m.maxHeap, heap.Pop(m.minHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.maxHeap.Len() > m.minHeap.Len() {
		return float64(m.maxHeap.Peak())
	}

	return float64(m.minHeap.Peak()+m.maxHeap.Peak()) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/*
Follow up:

If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
	1. Represented the stream by an array count where count[k] is the number of times
		that k appears in the sample.
	2. Finding of the median element takes O(100) steps which is O(1) time complexity.

If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
	1. Represented the 99% of the stream by an array count where count[k] is the number of times
		that k appears in the sample. And the 1% can be keeped in a hash map
	2. For finding of the median element, iterate through the map, then the array,
		and again through the map to find element that have keys greater than 100.

Similar Leetcode problem:
https://leetcode.com/problems/statistics-from-a-large-sample/
*/
