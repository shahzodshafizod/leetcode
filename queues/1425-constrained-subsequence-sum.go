package queues

// https://leetcode.com/problems/constrained-subsequence-sum/

// Approach: Sliding Window + Monotonic Decreasing Queue
// Time: O(n)
// Space: O(n)
func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	queue := make([]int, n)
	head, tail := 0, 0
	result := nums[0]
	for idx := 1; idx < n; idx++ {
		if queue[head]+k < idx {
			head++ // popleft
		}
		nums[idx] += max(0, nums[queue[head]])
		result = max(result, nums[idx])
		for tail >= head && nums[queue[tail]] < nums[idx] {
			tail-- // pop
		}
		tail++
		queue[tail] = idx // append
	}
	return result
}

// // Approach: Heap/Priority Queue
// // Time: O(n Log n)
// // Space: O(n)
// func constrainedSubsetSum(nums []int, k int) int {
// 	var result = nums[0]
// 	var maxHeap = design.NewPQ(
// 		[]int{0},
// 		func(x, y int) bool {
// 			return nums[x] < nums[y]
// 		},
// 	)
// 	for idx, n := 1, len(nums); idx < n; idx++ {
// 		for idx-maxHeap.Peek() > k {
// 			maxHeap.Pop()
// 		}
// 		nums[idx] += max(0, nums[maxHeap.Peek()])
// 		result = max(result, nums[idx])
// 		maxHeap.Push(idx)
// 	}
// 	return result
// }

// // Approach: Dynamic Programming (Bottom-Up)
// // Time: O(n*k)
// // Space: O(1)
// func constrainedSubsetSum(nums []int, k int) int {
// 	var n = len(nums)
// 	var result = nums[0]
// 	var original int
// 	for idx := 1; idx < n; idx++ {
// 		original = nums[idx]
// 		for next := max(0, idx-k); next < idx; next++ {
// 			nums[idx] = max(nums[idx], original+nums[next])
// 		}
// 		result = max(result, nums[idx])
// 	}
// 	return result
// }

// // Approach: Dynamic Programming (Top-Down)
// // Time: O(n*k)
// // Space: O(n)
// func constrainedSubsetSum(nums []int, k int) int {
// 	var n = len(nums)
// 	var cache = make([]*int, n)
// 	var result = nums[0]
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if cache[idx] != nil {
// 			return *cache[idx]
// 		}
// 		var maxsum = nums[idx]
// 		for j := 1; j <= k && idx+j < n; j++ {
// 			maxsum = max(maxsum, nums[idx]+dp(idx+j))
// 		}
// 		cache[idx] = &maxsum
// 		result = max(result, maxsum)
// 		return maxsum
// 	}
// 	dp(0)
// 	return result
// }
