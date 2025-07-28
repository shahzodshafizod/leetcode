package dp

// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/

// // Approach #3: Bit Manipulation
// // Time: O(n x maxor), maxor=2^n
// // Space: O(1)
// func countMaxOrSubsets(nums []int) int {
// 	maxor := 0
// 	for _, num := range nums {
// 		maxor |= num
// 	}
// 	count := 0
// 	n := len(nums)
// 	var logicsum int
// 	for subset, sn := 1, 1<<n; subset < sn; subset++ {
// 		logicsum = 0
// 		for idx := 0; idx < n; idx++ {
// 			if (1<<idx)&subset != 0 {
// 				logicsum |= nums[idx]
// 			}
// 		}
// 		if logicsum == maxor {
// 			count++
// 		}
// 	}
// 	return count
// }

// // Approach #2: Dynamic Programming (Bottom-Up)
// // Time: O(n x maxor), maxor=2^n
// // Space: O(maxor)
// func countMaxOrSubsets(nums []int) int {
// 	dp := map[int]int{0: 1}
// 	maxor := 0
// 	for _, curr := range nums {
// 		maxor |= curr
// 		newdp := maps.Clone(dp)
// 		for prev, cnt := range dp {
// 			newdp[prev|curr] += cnt
// 		}
// 		dp = newdp
// 	}
// 	return dp[maxor]
// }

// Approach #1: Dynamic Programming (Top-Down)
// Caching is inefficient:
// w/ cache: 16*131071 = 2_097_136
// w/o cache: 2^16 = 65_536
// Time: O(2^n)
// Space: O(2^n)
func countMaxOrSubsets(nums []int) int {
	maxor := 0
	for _, num := range nums {
		maxor |= num
	}
	n := len(nums)
	// bin(int(1e5)) = 11000011010100000
	// dec(11111111111111111) = 131071
	var dfs func(idx int, logicsum int) int
	dfs = func(idx int, logicsum int) int {
		if idx == n {
			if logicsum == maxor {
				return 1
			}
			return 0
		}
		// decision to skip
		count := dfs(idx+1, logicsum)
		// decision to include
		count += dfs(idx+1, logicsum|nums[idx])
		return count
	}
	return dfs(0, 0)
}
