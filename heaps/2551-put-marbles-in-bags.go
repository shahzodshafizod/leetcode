package heaps

import "sort"

// https://leetcode.com/problems/put-marbles-in-bags/

// Approach #3: Sorting (Space-Optimized)
// Time: O(nlogn)
// Space: O(1)
func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	for idx := 1; idx < n; idx++ {
		weights[idx-1] += weights[idx]
	}
	weights = weights[:n-1]
	sort.Ints(weights)
	var min, max int64 = 0, 0
	for idx := 0; idx < k-1; idx++ {
		min += int64(weights[idx])
		max += int64(weights[n-2-idx])
	}
	return max - min
}

// // Approach #2: Sorting
// // Time: O(nlogn)
// // Space: O(n)
// func putMarbles(weights []int, k int) int64 {
// 	var n = len(weights) - 1
// 	var borders = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		borders[idx] = weights[idx] + weights[idx+1]
// 	}
// 	sort.Ints(borders)
// 	var min, max int64
// 	for left, right := 0, n-1; left < k-1; {
// 		min += int64(borders[left])
// 		max += int64(borders[right])
// 		left++
// 		right--
// 	}
// 	return int64(max - min)
// }

// // Approach #1: Heap (Priority Queue)
// // Time: O(nlogn)
// // Space: O(2n) = O(n)
// func putMarbles(weights []int, k int) int64 {
// 	var n = len(weights) - 1
// 	var mins, maxs = make([]int, n), make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		mins[idx] = weights[idx] + weights[idx+1]
// 		maxs[idx] = mins[idx]
// 	}
// 	var minHeap = pkg.NewHeap(mins, func(x int, y int) bool { return x < y })
// 	var maxHeap = pkg.NewHeap(maxs, func(x int, y int) bool { return x > y })
// 	heap.Init(minHeap)
// 	heap.Init(maxHeap)
// 	var min, max int64 = 0, 0
// 	for i := 1; i < k; i++ {
// 		min += int64(heap.Pop(minHeap).(int))
// 		max += int64(heap.Pop(maxHeap).(int))
// 	}
// 	return max - min
// }
