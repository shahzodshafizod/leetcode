package design

import "container/heap"

// https://leetcode.com/problems/find-median-from-data-stream/

type medianHeap struct {
	data    []int
	compare func(data []int, i int, j int) bool
}

var _ heap.Interface = &medianHeap{}

func (m *medianHeap) Len() int           { return len(m.data) }
func (m *medianHeap) Less(i, j int) bool { return m.compare(m.data, i, j) }
func (m *medianHeap) Swap(i, j int)      { m.data[i], m.data[j] = m.data[j], m.data[i] }
func (m *medianHeap) Push(x any)         { m.data = append(m.data, x.(int)) }
func (m *medianHeap) Pop() any           { last := m.data[m.Len()-1]; m.data = m.data[:m.Len()-1]; return last }
func (m *medianHeap) Top() int           { return m.data[0] }

type MedianFinder struct {
	minHeap *medianHeap // right from median
	maxHeap *medianHeap // left from median
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		minHeap: &medianHeap{
			data:    make([]int, 0),
			compare: func(data []int, i, j int) bool { return data[i] < data[j] },
		},
		maxHeap: &medianHeap{
			data:    make([]int, 0),
			compare: func(data []int, i, j int) bool { return data[i] > data[j] },
		},
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
	for m.minHeap.Len() > 0 && (m.minHeap.Len() > m.maxHeap.Len()+1 || float64(m.minHeap.Top()) < m.FindMedian()) {
		heap.Push(m.maxHeap, heap.Pop(m.minHeap))
	}
	for m.maxHeap.Len() > 0 && (m.maxHeap.Len() > m.minHeap.Len()+1 || float64(m.maxHeap.Top()) > m.FindMedian()) {
		heap.Push(m.minHeap, heap.Pop(m.maxHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	var minLen, maxLen = m.minHeap.Len(), m.maxHeap.Len()
	if minLen == 0 && maxLen == 0 {
		return 0
	}
	if minLen > maxLen {
		return float64(m.minHeap.Top())
	}
	if maxLen > minLen {
		return float64(m.maxHeap.Top())
	}
	return float64(m.minHeap.Top()+m.maxHeap.Top()) / 2
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
