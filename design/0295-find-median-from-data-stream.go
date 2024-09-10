package design

// https://leetcode.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	maxHeap PQ[int] // less that or equal to median
	minHeap PQ[int] // greater that median
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		maxHeap: NewPQ(make([]int, 0), func(x, y int) bool { return x < y }),
		minHeap: NewPQ(make([]int, 0), func(x, y int) bool { return x > y }),
	}
}

func (m *MedianFinder) AddNum(num int) {
	// m.maxHeap.Push(num)
	// m.minHeap.Push(m.maxHeap.Pop())
	// if m.minHeap.Len() > m.maxHeap.Len() {
	// 	m.maxHeap.Push(m.minHeap.Pop())
	// }
	if float64(num) <= m.FindMedian() {
		m.maxHeap.Push(num)
	} else {
		m.minHeap.Push(num)
	}
	// rebalance
	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		m.minHeap.Push(m.maxHeap.Pop())
	}
	if m.minHeap.Len() > m.maxHeap.Len() {
		m.maxHeap.Push(m.minHeap.Pop())
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.maxHeap.Len() > m.minHeap.Len() {
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
