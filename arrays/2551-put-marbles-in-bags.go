package arrays

import "sort"

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

// func putMarbles(weights []int, k int) int64 {
// 	var minHeap = design.NewHeap(make([]int, 0), func(x int, y int) bool { return x < y })
// 	var maxHeap = design.NewHeap(make([]int, 0), func(x int, y int) bool { return x > y })
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
