package dp

// https://leetcode.com/problems/house-robber/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n)
// Space: O(1)
func rob(nums []int) int {
	prev, curr := 0, 0
	for _, num := range nums {
		prev, curr = curr, max(prev+num, curr)
	}

	return curr
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(n)
// func rob(nums []int) int {
// 	var n = len(nums)
// 	var memo = make([]*int, n)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx >= n {
// 			return 0
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		memo[idx] = new(int)
// 		*memo[idx] = max(
// 			nums[idx]+dp(idx+2),
// 			dp(idx+1),
// 		)
// 		return *memo[idx]
// 	}
// 	return dp(0)
// }
