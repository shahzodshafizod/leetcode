package binarysearch

// https://leetcode.com/problems/split-array-largest-sum/

// Approach: Binary Search
// Time: O(NLogS), N=len(nums), S=sum(nums)
// Space: O(1)
func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, num := range nums {
		left = max(left, num)
		right += num
	}

	minmaxSum := right

	var mid, subarrays, currSum int
	for left <= right {
		mid = left + (right-left)/2

		subarrays = 1 // we're currently in the first subarray

		currSum = 0
		for _, num := range nums {
			currSum += num
			if currSum > mid {
				subarrays++
				currSum = num
			}
		}

		if subarrays <= k {
			minmaxSum = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return minmaxSum
}

// // Approach: Dynamic Programming
// // Time: O(N^2 * K)
// // Space: O(N * K)
// func splitArray(nums []int, k int) int {
// 	var n = len(nums)
// 	var dp = make([][]*int, n)
// 	for idx := range dp {
// 		dp[idx] = make([]*int, k)
// 	}
// 	var dfs func(idx int, k int) int
// 	dfs = func(idx int, k int) int {
// 		if k == 1 {
// 			var sum = 0
// 			for _, num := range nums[idx:] {
// 				sum += num
// 			}
// 			return sum
// 		}
// 		if dp[idx][k-1] != nil {
// 			return *dp[idx][k-1]
// 		}
// 		var minmaxSum = math.MaxInt
// 		var sum = 0
// 		for j := idx; j <= n-k; j++ {
// 			sum += nums[j]
// 			minmaxSum = min(minmaxSum, max(sum, dfs(j+1, k-1)))
// 			if sum > minmaxSum { // optimization
// 				break
// 			}
// 		}
// 		dp[idx][k-1] = &minmaxSum
// 		return minmaxSum
// 	}
// 	return dfs(0, k)
// }
