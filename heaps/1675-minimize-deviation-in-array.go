package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimize-deviation-in-array/

// Approach: Heap (Priority Queue) (Optimized)
// Time: O(n + logn x logm), n=len(nums), m=max(nums[i])
// Space: O(n)
func minimumDeviation(nums []int) int {
	n := len(nums)
	maxHeap := make([]int, 0, n)
	var smallest int = 1e9
	for _, num := range nums {
		if num&1 == 1 {
			num *= 2
		}
		maxHeap = append(maxHeap, num)
		smallest = min(smallest, num)
	}
	pq := pkg.NewHeap(maxHeap,
		func(x, y int) bool {
			return x > y
		},
	)
	heap.Init(pq)
	var num int
	var deviation int = 1e9
	for pq.Len() == n {
		num = maxHeap[0]
		deviation = min(deviation, num-smallest)
		if num&1 == 0 {
			num /= 2
			smallest = min(smallest, num)
			maxHeap[0] = num
			heap.Fix(pq, 0)
		} else {
			heap.Pop(pq)
		}
	}
	return deviation
}

// // Approach: Heap (Priority Queue)
// // Time: O(n logm logn), n=len(nums), m=max(nums[i])
// // Space: O(n)
// func minimumDeviation(nums []int) int {
// 	var n = len(nums)
// 	var maxHeap = make([][2]int, n)
// 	var mn, mx int
// 	var largest = 0
// 	for idx := range nums {
// 		mn, mx = nums[idx], nums[idx]
// 		for mn&1 == 0 {
// 			mn /= 2
// 		}
// 		if mx&1 == 1 {
// 			mx *= 2
// 		}
// 		maxHeap[idx][0] = mn
// 		maxHeap[idx][1] = mx
// 		largest = max(largest, mn)
// 	}
// 	var pq = pkg.NewHeap(maxHeap,
// 		func(x, y [2]int) bool {
// 			return x[0] < y[0]
// 		},
// 	)
// 	heap.Init(pq)
// 	var deviation int = 1e9
// 	for pq.Len() == n {
// 		mn, mx = pq.Peek()[0], pq.Peek()[1]
// 		heap.Pop(pq)
// 		deviation = min(deviation, largest-mn)
// 		if mn < mx {
// 			heap.Push(pq, [2]int{2 * mn, mx})
// 			largest = max(largest, mn*2)
// 		}
// 	}
// 	return deviation
// }
