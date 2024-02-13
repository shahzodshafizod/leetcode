package design

import "container/heap"

// https://leetcode.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	minHeap *Heap[int] // right from median
	maxHeap *Heap[int] // left from median
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(make([]int, 0), func(x, y int) bool { return x < y }),
		maxHeap: NewHeap(make([]int, 0), func(x, y int) bool { return x > y }),
	}
}

func (m *MedianFinder) AddNum(num int) {
	if float64(num) > m.FindMedian() {
		heap.Push(m.minHeap, num)
	} else {
		heap.Push(m.maxHeap, num)
	}
	m.rebalance()
}

func (m *MedianFinder) rebalance() {
	// Every element in the min-heap is greater or equal to the median,
	// and every element in the max-heap is less or equal to the median.
	// THAT'S: maxHeap is in left side, minHeap is in right side.
	for m.minHeap.Len() > 0 && (m.minHeap.Len() > m.maxHeap.Len()+1 || float64(m.minHeap.Peek()) < m.FindMedian()) {
		heap.Push(m.maxHeap, heap.Pop(m.minHeap))
	}
	for m.maxHeap.Len() > 0 && (m.maxHeap.Len() > m.minHeap.Len()+1 || float64(m.maxHeap.Peek()) > m.FindMedian()) {
		heap.Push(m.minHeap, heap.Pop(m.maxHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	var minLen, maxLen = m.minHeap.Len(), m.maxHeap.Len()
	if minLen == 0 && maxLen == 0 {
		return 0
	}
	if minLen > maxLen {
		return float64(m.minHeap.Peek())
	}
	if maxLen > minLen {
		return float64(m.maxHeap.Peek())
	}
	return float64(m.minHeap.Peek()+m.maxHeap.Peek()) / 2
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
