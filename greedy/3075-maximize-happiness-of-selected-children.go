package greedy

import (
	"sort"
)

// https://leetcode.com/problems/maximize-happiness-of-selected-children/

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	var sum int64

	for turns, idx := 0, len(happiness)-1; turns < k; {
		if happiness[idx]-turns > 0 {
			sum += int64(happiness[idx] - turns)
		}

		idx--
		turns++
	}

	return sum
}

// // time: O(n log n)
// // space: O(1)
// func maximumHappinessSum(happiness []int, k int) int64 {
// 	var heap = design.NewPQ(happiness, func(x, y int) bool { return x < y })
// 	heap.Heapify()
// 	var sum int64 = 0
// 	var turns = 0
// 	var top int
// 	for k > 0 {
// 		k--
// 		top = heap.Pop()
// 		if top-turns > 0 {
// 			sum += int64(top - turns)
// 		}
// 		turns++
// 	}
// 	return sum
// }
