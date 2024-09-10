package slidingwindows

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/sliding-window-median/

// Approach: Two Heaps (Two Priority Queues)
// Time: O(n log k)
// Space: O(2k) = O(k)
func medianSlidingWindow(nums []int, k int) []float64 {
	var maxHeap = design.NewPQ(make([]int, 0), func(x, y int) bool { return x < y })
	var minHeap = design.NewPQ(make([]int, 0), func(x, y int) bool { return x > y })
	for idx := 0; idx < k; idx++ {
		maxHeap.Push(nums[idx])
		minHeap.Push(maxHeap.Pop())
		if minHeap.Len() > maxHeap.Len() {
			maxHeap.Push(minHeap.Pop())
		}
	}
	var n = len(nums)
	var medians = make([]float64, n-k+1)
	var median float64
	if k&1 == 1 {
		median = float64(maxHeap.Peek())
	} else {
		median = float64(maxHeap.Peek()+minHeap.Peek()) / 2
	}
	medians[0] = median
	var outbounds = make(map[int]int)
	var outbound, balance int
	for idx := k; idx < n; idx++ {
		outbound = nums[idx-k]
		outbounds[outbound]++

		if float64(outbound) <= median {
			balance = -1
		} else {
			balance = 1
		}

		if float64(nums[idx]) <= median {
			balance++
			maxHeap.Push(nums[idx])
		} else {
			balance--
			minHeap.Push(nums[idx])
		}

		if balance < 0 {
			maxHeap.Push(minHeap.Pop())
		} else if balance > 0 {
			minHeap.Push(maxHeap.Pop())
		}

		for maxHeap.Len() > 0 && outbounds[maxHeap.Peek()] > 0 {
			outbounds[maxHeap.Pop()]--
		}
		for minHeap.Len() > 0 && outbounds[minHeap.Peek()] > 0 {
			outbounds[minHeap.Pop()]--
		}

		if k&1 == 1 {
			median = float64(maxHeap.Peek())
		} else {
			median = float64(maxHeap.Peek()+minHeap.Peek()) / 2
		}
		medians[idx-k+1] = median
	}
	return medians
}

// // Approach: Brute-Force
// func medianSlidingWindow(nums []int, k int) []float64 {
// 	var n = len(nums)
// 	var window = make([]int, k)
// 	var medians = make([]float64, n-k+1)
// 	for end := k; end <= n; end++ {
// 		copy(window, nums[end-k:end])
// 		sort.Ints(window)
// 		medians[end-k] = float64(window[k/2-1+k&1]+window[k/2]) / 2
// 	}
// 	return medians
// }
