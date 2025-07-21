package dp

/*
Dynamic Programming
5 Simple Steps For Solving Dynamic Programming Problems:
	- Visualize Examples
	- Find an appropriate subproblem
	- Find relationships among subproblems
	- Generalize the relationship
	- Implement by solving subproblems in order
*/

// https://leetcode.com/problems/longest-increasing-subsequence/

// Approach: Binary Search
// Time: O(nlogn)
// Space: O(n)
// CACHE - does not keep the real subsequence
// its length is the length of the longest increasing subsequence
// we'll keep many subsequences in the same array by replacing previous values
func lengthOfLIS(nums []int) int {
	lis := make([]int, 0)
	for _, num := range nums {
		if len(lis) == 0 || num > lis[len(lis)-1] {
			lis = append(lis, num)
			continue
		}
		// it's not a binary search, it's building blocks to the full solution
		left, right := 0, len(lis)
		for left < right {
			mid := left + (right-left)/2
			if lis[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// the current element (num) is either equal to or smaller than
		// the existing element (lis[left]), and this allows for potentially
		// more elements to be added to the subsequence in the future.
		lis[left] = num
	}
	return len(lis)
}

// // Approach: Dynamic Programming (Bottom-Up)
// // Time: O(nn)
// // Space: O(n)
// func lengthOfLIS(nums []int) int {
// 	var length = 0
// 	var dp = make([]int, len(nums))
// 	for curr := range nums {
// 		dp[curr] = 1
// 		for prev := 0; prev < curr; prev++ {
// 			if nums[prev] < nums[curr] {
// 				dp[curr] = max(dp[curr], 1+dp[prev])
// 			}
// 		}
// 		length = max(length, dp[curr])
// 	}
// 	return length
// }
