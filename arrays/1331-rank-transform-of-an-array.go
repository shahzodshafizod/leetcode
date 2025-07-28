package arrays

import (
	"sort"
)

// https://leetcode.com/problems/rank-transform-of-an-array/

// Approach: Sorting + Hash Table
// Time: O(NLogN)
// Space: O(N)
func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	ranks := make(map[int]int)

	prev, rank := 0, 0
	for _, num := range sorted {
		if rank > 0 && num == prev {
			continue
		}

		rank++
		ranks[num] = rank
		prev = num
	}

	for idx, num := range arr {
		arr[idx] = ranks[num]
	}

	return arr
}

// // Approach: Heap (Priority Queue)
// // Time: O(NLogN)
// // Space: O(N)
// func arrayRankTransform(arr []int) []int {
// 	var minHeap = design.NewHeap(
// 		make([][2]int, 0),
// 		func(x, y [2]int) bool { return x[0] < y[0] },
// 	)
// 	for idx, num := range arr {
// 		heap.Push(minHeap, [2]int{num, idx})
// 	}
// 	var rank = 0
// 	var prev, num, idx int
// 	for minHeap.Len() > 0 {
// 		item := heap.Pop(minHeap).([2]int)
// 		num, idx = item[0], item[1]
// 		if rank == 0 || num != prev {
// 			rank++
// 		}
// 		arr[idx] = rank
// 		prev = num
// 	}
// 	return arr
// }
