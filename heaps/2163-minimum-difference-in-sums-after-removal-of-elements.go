package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/

// Time: O(n log n)
// Space: O(n)
func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	mnhp := pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x < y })
	sumsec := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		heap.Push(mnhp, nums[2*n+i])
		sumsec[n] += int64(nums[2*n+i])
	}

	for i := n - 1; i >= 0; i-- {
		heap.Push(mnhp, nums[n+i])
		sumsec[i] = sumsec[i+1] + int64(nums[n+i]) - int64(heap.Pop(mnhp).(int))
	}

	mxhp := pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x > y })

	var sumfir int64 = 0

	for i := 0; i < n; i++ {
		heap.Push(mxhp, nums[i])
		sumfir += int64(nums[i])
	}

	var diff int64 = sumfir - sumsec[0]

	for i := 0; i < n; i++ {
		heap.Push(mxhp, nums[n+i])
		sumfir += int64(nums[n+i]) - int64(heap.Pop(mxhp).(int))
		diff = min(diff, sumfir-sumsec[i+1])
	}

	return diff
}
