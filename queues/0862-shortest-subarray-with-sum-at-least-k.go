package queues

// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

// Approach#3: Deque (Monotonic Queue)
// Time: O(n)
// Space: O(n)
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for idx := range nums {
		presum[idx+1] = presum[idx] + nums[idx]
	}
	queue := make([]int, n+1)
	left, right, length := 0, 0, n+1
	for end := 1; end <= n; end++ {
		// Find the minimum valid window ending at 'end'
		for left <= right && presum[end]-presum[queue[left]] >= k {
			length = min(length, end-queue[left])
			left++
		}
		// Validate the Monotonic Increasing Queue
		for left <= right && presum[end] <= presum[queue[right]] {
			right--
		}
		right++
		queue[right] = end
	}
	if length > n {
		return -1
	}
	return length
}

// // Approach#2: Heap (Priority Queue) + Prefix Sum
// // Time: O(nlogn)
// // Space: O(n)
// func shortestSubarray(nums []int, k int) int {
// 	var pq = pkg.NewHeap(
// 		make([][2]int, 0),
// 		func(x, y [2]int) bool { return x[0] < y[0] },
// 	)
// 	var n = len(nums)
// 	var presum, length = 0, n + 1
// 	for idx := range nums {
// 		presum += nums[idx]
// 		if presum >= k {
// 			length = min(length, idx+1)
// 		}
// 		for pq.Len() > 0 && presum-pq.Peek()[0] >= k {
// 			length = min(length, idx-heap.Pop(pq).([2]int)[1])
// 		}
// 		heap.Push(pq, [2]int{presum, idx})
// 	}
// 	if length > n {
// 		return -1
// 	}
// 	return length
// }

// // Approach#1: Sliding Window + Prefix Sum (Time Limit Exceeded)
// // Time: O(nn)
// // Space: O(n)
// func shortestSubarray(nums []int, k int) int {
// 	var n = len(nums)
// 	var presum = make([]int, n+1)
// 	for idx := 1; idx <= n; idx++ {
// 		presum[idx] = nums[idx-1] + presum[idx-1]
// 	}
// 	for size := 1; size <= n; size++ {
// 		for end := size; end <= n; end++ {
// 			if presum[end]-presum[end-size] >= k {
// 				return size
// 			}
// 		}
// 	}
// 	return -1
// }
