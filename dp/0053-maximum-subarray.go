package dp

/*
Kadane's algorithm
The fundamental concept behind Kadane's algorithm is the idea of
"local maximum" and "global maximum".
	1. The local maximum at each position is the maximum of the current number
		and the sum of the current number and the previous local maximum.
		This basically means we have two choices at each step:
		to add the current number to the subarray
		or start a new subarray with the current number.
		We choose the option that gives the maximum sum.
	2. The global maximum is the maximum sum of all subarrays that we've seen so far.

Pseudocode of Kadane's algorithm:
Initialize:
    max_so_far = INT_MIN
    max_ending_here = 0

Loop for each element of the array
  (a) max_ending_here = max_ending_here + a[i]
  (b) if(max_so_far < max_ending_here)
            max_so_far = max_ending_here
  (c) if(max_ending_here < 0)
            max_ending_here = 0
return max_so_far

Example: nums:[-2, 1, -3, 4, -1, 2, 1, -5, 4], max:6

[[-2], 1, -3, 4, -1, 2, 1, -5, 4], sum:-2, max:-2, sum:0
[-2, [1], -3, 4, -1, 2, 1, -5, 4], sum:1, max:1
[-2, [1, -3], 4, -1, 2, 1, -5, 4], sum:-2, max:1, sum:0
[-2, 1, -3, [4], -1, 2, 1, -5, 4], sum:4, max:4
[-2, 1, -3, [4, -1], 2, 1, -5, 4], sum:3, max:4
[-2, 1, -3, [4, -1, 2], 1, -5, 4], sum:5, max:5
[-2, 1, -3, [4, -1, 2, 1], -5, 4], sum:6, max:6
[-2, 1, -3, [4, -1, 2, 1, -5], 4], sum:1, max:6
[-2, 1, -3, [4, -1, 2, 1, -5, 4]], sum:5, max:6
*/

// https://leetcode.com/problems/maximum-subarray/

// Kadane's algorithm: O(n)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0
	for _, num := range nums {
		currSum += num
		maxSum = max(maxSum, currSum)
		if currSum < 0 {
			// beginning of a new subsequence
			currSum = 0
		}
	}
	return maxSum
}

// // Brute Force: O(n^2)
// func maxSubArray(nums []int) int {
// 	var maxSum = math.MinInt
// 	var sum int
// 	for i, len := 0, len(nums); i < len; i++ {
// 		sum = 0
// 		for j := i; j < len; j++ {
// 			sum += nums[j]
// 			maxSum = max(sum, maxSum)
// 		}
// 	}
// 	return maxSum
// }

// Follow up: If you have figured out the O(n) solution,
// try coding another solution using the divide and conquer approach, which is more subtle.

// // Divide and Conquer (Optimized)
// // time: O(n)
// // space: O(n)
// func maxSubArray(nums []int) int {
// 	var n = len(nums)
// 	prevMaxSum := make([]int, n)
// 	prevMaxSum[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		prevMaxSum[i] = nums[i] + max(0, prevMaxSum[i-1])
// 	}
// 	postMaxSum := make([]int, n)
// 	postMaxSum[n-1] = nums[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		postMaxSum[i] = nums[i] + max(0, postMaxSum[i+1])
// 	}
// 	var dp func(left int, right int) int
// 	dp = func(left int, right int) int {
// 		if left == right {
// 			return nums[left]
// 		}
// 		var mid = (left + right) / 2
// 		return max(
// 			max(dp(left, mid), dp(mid+1, right)),
// 			prevMaxSum[mid]+postMaxSum[mid+1],
// 		)
// 	}
// 	return dp(0, n-1)
// }

// // Divide and Conquer
// // time: O(n log n)
// // space: O(1)
// func maxSubArray(nums []int) int {
// 	var n = len(nums)
// 	var dp func(left int, right int) int
// 	dp = func(left int, right int) int {
// 		if left == right {
// 			return nums[left]
// 		}
// 		var mid = (left + right) / 2
// 		prevMaxSum, sum := math.MinInt, 0
// 		for i := mid; i >= left; i-- {
// 			sum += nums[i]
// 			prevMaxSum = max(prevMaxSum, sum)
// 		}
// 		postMaxSum, sum := math.MinInt, 0
// 		for i := mid; i <= right; i++ {
// 			sum += nums[i]
// 			postMaxSum = max(postMaxSum, sum)
// 		}
// 		return max(
// 			max(dp(left, mid), dp(mid+1, right)),
// 			prevMaxSum+postMaxSum-nums[mid],
// 		)
// 	}
// 	return dp(0, n-1)
// }
