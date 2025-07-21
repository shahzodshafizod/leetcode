package bits

// https://leetcode.com/problems/sum-of-all-subset-xor-totals/

// Approach #2: Bit Manipulation
// Time: O(n)
// Space: O(1)
func subsetXORSum(nums []int) int {
	total := 0
	for _, num := range nums {
		total |= num
	}
	return total << (len(nums) - 1)
}

// // Approach #1: Backtracking
// // Time: O(2^n)
// // Space: O(2^n)
// func subsetXORSum(nums []int) int {
// 	var n = len(nums)
// 	var dp func(idx int, total int) int
// 	dp = func(idx int, total int) int {
// 		if idx == n {
// 			return total
// 		}
// 		return dp(idx+1, total) +
// 			dp(idx+1, total^nums[idx])
// 	}
// 	return dp(0, 0)
// }
