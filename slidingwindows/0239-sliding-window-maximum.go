package slidingwindows

// https://leetcode.com/problems/sliding-window-maximum/

// Approach #4: Monotonic Decreasing Queue
// Time: O(n)
// Space: O(n)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	queue := make([]int, n)
	head, tail := 0, -1
	maxes := make([]int, n-k+1)

	for idx := range nums {
		if head <= tail && idx-queue[head] >= k {
			head++
		}

		for head <= tail && nums[queue[tail]] <= nums[idx] {
			tail--
		}

		tail++

		queue[tail] = idx
		if idx+1 >= k {
			maxes[idx-k+1] = nums[queue[head]]
		}
	}

	return maxes
}

// // Approach #3: Monotonic Decreasing Queue with container/list
// // Time: O(n)
// // Space: O(k)
// func maxSlidingWindow(nums []int, k int) []int {
// 	var list = list.New()
// 	var n = len(nums)
// 	var max = make([]int, n-k+1)
// 	for idx := range nums {
// 		if list.Len() > 0 && idx-list.Front().Value.(int) >= k {
// 			list.Remove(list.Front())
// 		}
// 		for list.Len() > 0 && nums[list.Back().Value.(int)] <= nums[idx] {
// 			list.Remove(list.Back())
// 		}
// 		list.PushBack(idx)
// 		if idx+1 >= k {
// 			max[idx-k+1] = nums[list.Front().Value.(int)]
// 		}
// 	}
// 	return max
// }

// // Approach #2: Priority Queue
// // Time: O(n log n)
// // Space: O(n)
// func maxSlidingWindow(nums []int, k int) []int {
// 	var maxHeap = design.NewPQ(
// 		make([]int, 0),
// 		func(i, j int) bool {
// 			return nums[i] < nums[j]
// 		},
// 	)
// 	var maxes = make([]int, len(nums)-k+1)
// 	for idx := range nums {
// 		for maxHeap.Len() > 0 && idx-maxHeap.Peek() >= k {
// 			maxHeap.Pop()
// 		}
// 		maxHeap.Push(idx)
// 		if idx+1 >= k {
// 			maxes[idx-k+1] = nums[maxHeap.Peek()]
// 		}
// 	}
// 	return maxes
// }

// // Approach #1: Brute-Force
// // Time: O(n x k)
// // Space: O(1)
// func maxSlidingWindow(nums []int, k int) []int {
// 	var n = len(nums)
// 	var maxes = make([]int, n-k+1)
// 	for end := k; end <= n; end++ {
// 		maxes[end-k] = -1e4
// 		for idx := end - k; idx < end; idx++ {
// 			maxes[end-k] = max(maxes[end-k], nums[idx])
// 		}
// 	}
// 	return maxes
// }
