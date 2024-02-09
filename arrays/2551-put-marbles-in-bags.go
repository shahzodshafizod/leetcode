package arrays

import (
	"sort"
)

// https://leetcode.com/problems/put-marbles-in-bags/

func putMarbles(weights []int, k int) int64 {
	var n = len(weights)
	for i := 1; i < n; i++ {
		weights[i-1] += weights[i]
	}
	weights = weights[:n-1]
	n--
	sort.Ints(weights)
	var min, max int
	for left, right := 0, n-1; left < k-1; left, right = left+1, right-1 {
		min += weights[left]
		max += weights[right]
	}
	return int64(max - min)
}

// type marbleHeap struct {
// 	data    []int
// 	compare func(data []int, i int, j int) bool
// }

// var _ heap.Interface = &marbleHeap{}

// func (m *marbleHeap) Len() int           { return len(m.data) }
// func (m *marbleHeap) Less(i, j int) bool { return m.compare(m.data, i, j) }
// func (m *marbleHeap) Swap(i, j int)      { m.data[i], m.data[j] = m.data[j], m.data[i] }
// func (m *marbleHeap) Push(x any)         { m.data = append(m.data, x.(int)) }
// func (m *marbleHeap) Pop() any           { last := m.data[m.Len()-1]; m.data = m.data[:m.Len()-1]; return last }

// func putMarbles(weights []int, k int) int64 {
// 	var minHeap = &marbleHeap{compare: func(data []int, i, j int) bool { return data[i] < data[j] }}
// 	var maxHeap = &marbleHeap{compare: func(data []int, i, j int) bool { return data[i] > data[j] }}
// 	for i, len := 1, len(weights); i < len; i++ {
// 		sum := weights[i-1] + weights[i]
// 		heap.Push(minHeap, sum)
// 		heap.Push(maxHeap, sum)
// 	}
// 	var min, max int = 0, 0
// 	for i := 1; i < k; i++ {
// 		min += heap.Pop(minHeap).(int)
// 		max += heap.Pop(maxHeap).(int)
// 	}
// 	return int64(max - min)
// }
