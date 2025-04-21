package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/

// Approach: Heap (Priority Queue)
// Time: O(k log n)
// Space: O(n)
func getFinalState(nums []int, k int, multiplier int) []int {
	var n = len(nums)
	var indices = make([]int, n)
	for idx := range indices {
		indices[idx] = idx
	}
	var pq = pkg.NewHeap(indices, func(x, y int) bool {
		return nums[x] == nums[y] && x < y || nums[x] < nums[y]
	})
	heap.Init(pq)
	for ; k > 0; k-- {
		nums[pq.Peak()] *= multiplier
		heap.Fix(pq, 0)
	}
	return nums
}
