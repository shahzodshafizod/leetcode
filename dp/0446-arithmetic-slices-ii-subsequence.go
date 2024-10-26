package dp

// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/

/*
+-------------+---+---+---+---+-----+-------+
| diff \ nums | 0 | 1 | 2 | 2 | 3   |     4 |
|-------------+---+---+---+---+-----+-------|
| 0           |   |   |   | 1 |     |       |
| 1           |   | 1 | 2 | 2 | 6   |     7 |
| 2           |   |   | 1 | 1 | 1   |     4 |
| 3           |   |   |   |   | 1   |     1 |
| 4           |   |   |   |   |     |     1 |
+-------------+---+---+---+---+-----+-------+
| count       |   | 0 | 1 | 1 | 2+2 | 6+1+1 |
| result      |   | 0 | 1 | 2 |  6  |  "14" |
+-------------+---+---+---+---+-----+-------+
*/

// Approach: Bottom-Up Dynamic Programming
// Time: O(nn), n=len(nums)
// Space: O(nd), d=max(nums)-min(nums)
func numberOfArithmeticSlicesSubsequence(nums []int) int {
	var n = len(nums)
	var dp = make([]map[int]int, n)
	for idx := range dp {
		dp[idx] = make(map[int]int)
	}
	var difference int
	var count = 0
	for curr := 1; curr < n; curr++ {
		for prev := 0; prev < curr; prev++ {
			difference = nums[curr] - nums[prev]
			dp[curr][difference] += dp[prev][difference] + 1
			count += dp[prev][difference]
		}
	}
	return count
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(nnn), n=len(nums)
// // Space: O(nd), d=max(nums)-min(nums)
// func numberOfArithmeticSlicesSubsequence(nums []int) int {
// 	var n = len(nums)
// 	var dp = make([]map[int]int, n)
// 	for idx := range dp {
// 		dp[idx] = make(map[int]int)
// 	}
// 	var dfs func(prev int, curr int) int
// 	dfs = func(prev int, curr int) int {
// 		var difference = nums[curr] - nums[prev]
// 		var count = 0
// 		for next := curr + 1; next < n; next++ {
// 			if nums[next]-nums[curr] == difference {
// 				count += 1 + dfs(curr, next)
// 			}
// 		}
// 		return count
// 	}
// 	var count = 0
// 	for curr := 0; curr < n; curr++ {
// 		for prev := 0; prev < curr; prev++ {
// 			count += dfs(prev, curr)
// 		}
// 	}
// 	return count
// }
