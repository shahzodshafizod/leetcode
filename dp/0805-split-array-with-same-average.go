package dp

/*
S, N = sum(nums), len(nums)
SA, NA = sum(A), len(A)
SB, NB = sum(B), len(B)
nums[i] -> A or B
SA + SB = S
NA + NB = N

SA/NA = SB/NB
SA/NA = (S-SA)/NB
SA*NB = (S-SA)*NA
SA*NB = S*NA - SA*NA
SA*NB + SA*NA = S*NA
SA*(NA+NB) = S*NA
SA*N = S*NA
SA/NA = S/N

If we partition nums into two arrays with equal average,
then nums itself has to have this average as well
*/

// https://leetcode.com/problems/split-array-with-same-average/

// Approach #4: Bit Manipulation
// Time: O(N*TotalSum)
// Space: O(TotalSum)
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}

	total := 0
	for _, num := range nums {
		total += num
	}

	if total == 0 {
		return true
	}

	possible := false

	for size := 1; size < n; size++ {
		if total*size%n == 0 {
			possible = true

			break
		}
	}

	if !possible {
		return false
	}

	dp := make([]int, total+1)
	dp[0] = 1

	for _, num := range nums {
		for target := total; target >= 0; target-- {
			if dp[target] != 0 {
				dp[target+num] |= dp[target] << 1
			}
		}
	}

	for target := 1; target <= total; target++ {
		if target*n%total == 0 {
			shift := target * n / total
			if 0 < shift && shift < n && (dp[target]&(1<<shift)) != 0 {
				return true
			}
		}
	}

	return false
}

// // Approach #3: Bottom-Up Dynamic Programming
// // Time: O(N^2 * TotalSum)
// // Space: O(N * TotalSum)
// func splitArraySameAverage(nums []int) bool {
// 	var n = len(nums)
// 	var total = 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	var possible = false
// 	for size := 1; size < n; size++ {
// 		if total*size%n == 0 {
// 			possible = true
// 			break
// 		}
// 	}
// 	if !possible {
// 		return false
// 	}
// 	var dp = make([][]bool, n/2+1)
// 	for idx := range dp {
// 		dp[idx] = make([]bool, total/2+1)
// 	}
// 	dp[0][0] = true
// 	for _, num := range nums {
// 		for size := n / 2; size > 0; size-- {
// 			for target := total / 2; target >= num; target-- {
// 				dp[size][target] = dp[size][target] || dp[size-1][target-num]
// 			}
// 		}
// 	}
// 	for size := 1; size <= n/2; size++ {
// 		if total*size%n == 0 && dp[size][total*size/n] {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Approach #2: Top-Down Dynamic Programming (fatal error: runtime: out of memory)
// // Time: O(N^3)
// // Space: O(N^2 * TotalSum)
// func splitArraySameAverage(nums []int) bool {
// 	var total = 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	var n = len(nums)
// 	var dp = make([][][]*bool, n+1)
// 	for idx := range dp {
// 		dp[idx] = make([][]*bool, n+1)
// 		for j := range dp[idx] {
// 			dp[idx][j] = make([]*bool, total+1)
// 		}
// 	}
// 	var dfs func(start int, size int, target int) bool
// 	dfs = func(start int, size int, target int) bool {
// 		if size == 0 && target == 0 {
// 			return true
// 		}
// 		if start == n || size == 0 || target < 0 {
// 			return false
// 		}
// 		if dp[start][size][target] != nil {
// 			return *dp[start][size][target]
// 		}
// 		dp[start][size][target] = new(bool)
// 		*dp[start][size][target] = dfs(start+1, size, target) ||
// 			dfs(start+1, size-1, target-nums[start])
// 		return *dp[start][size][target]
// 	}
// 	for size := 1; size <= n/2; size++ {
// 		if total*size%n == 0 {
// 			if dfs(0, size, total*size/n) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
// // Time: O(2^N)
// // Space: O(2^N)
// func splitArraySameAverage(nums []int) bool {
// 	var total = 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	var n = len(nums)
// 	var dfs func(idx int, sum1 int, len1 int) bool
// 	dfs = func(idx int, sum1 int, len1 int) bool {
// 		if idx == n {
// 			// sum1/len1 = sum2/len2
// 			// sum1*len2 = sum2*len1
// 			// len2 = n-len1
// 			// sum2 = total-sum1
// 			return 0 < len1 && len1 < n && sum1*(n-len1) == (total-sum1)*len1
// 		}
// 		// decision NOT to include
// 		if dfs(idx+1, sum1, len1) {
// 			return true
// 		}
// 		// decision to include
// 		return dfs(idx+1, sum1+nums[idx], len1+1)
// 	}
// 	return dfs(0, 0, 0)
// }
