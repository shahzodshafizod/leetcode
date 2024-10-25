package dp

// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/

// Approach #4: Stack
// Time: O(N*D)
// Space: O(N)
func minDifficulty(jobDifficulty []int, d int) int {
	var n = len(jobDifficulty)
	if n < d {
		return -1
	}
	var prev, curr = make([]int, n), make([]int, n)
	for idx := range curr {
		curr[idx] = 300_001
	}
	for day := 0; day < d; day++ {
		prev, curr = curr, prev
		var stack = make([]int, 0)
		var slen = 0
		for idx := day; idx < n; idx++ {
			curr[idx] = jobDifficulty[idx]
			if idx > 0 {
				curr[idx] += prev[idx-1]
			}
			for slen > 0 && jobDifficulty[stack[slen-1]] <= jobDifficulty[idx] {
				j := stack[slen-1]
				slen--
				stack = stack[:slen]
				curr[idx] = min(curr[idx], curr[j]-jobDifficulty[j]+jobDifficulty[idx])
			}
			if slen > 0 {
				curr[idx] = min(curr[idx], curr[stack[slen-1]])
			}
			stack = append(stack, idx)
			slen++
		}
	}
	return curr[n-1]
}

// // Approach #3: Bottom-Up Dynamic Programming (Optimized Memory)
// // Time: O(N^2 * D)
// // Space: O(N)
// func minDifficulty(jobDifficulty []int, d int) int {
// 	var n = len(jobDifficulty)
// 	if n < d {
// 		// not enough tasks to split into each day
// 		return -1
// 	}
// 	var prev, curr = make([]int, n), make([]int, n)
// 	curr[0] = jobDifficulty[0]
// 	for idx := 1; idx < n; idx++ {
// 		// day=1-1, only one day left, we have to do all the remaining jobs
// 		// fill out first column based on max(jobDiffculty[0:i])
// 		curr[idx] = max(jobDifficulty[idx], curr[idx-1])
// 	}
// 	var max_diff int
// 	for day := 1; day < d; day++ {
// 		prev, curr = curr, prev
// 		for idx := day; idx < n; idx++ {
// 			max_diff = 0
// 			curr[idx] = 300_001
// 			for j := idx; j >= day; j-- {
// 				max_diff = max(max_diff, jobDifficulty[j])
// 				curr[idx] = min(curr[idx], max_diff+prev[j-1])
// 			}
// 		}
// 	}
// 	return curr[n-1]
// }

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(N^2 * D)
// // Space: O(N * D)
// func minDifficulty(jobDifficulty []int, d int) int {
// 	var n = len(jobDifficulty)
// 	if n < d {
// 		// not enough tasks to split into each day
// 		return -1
// 	}
// 	var dp = make([][]int, n)
// 	var curmax = 0
// 	for idx := range dp {
// 		dp[idx] = make([]int, d)
// 		curmax = max(curmax, jobDifficulty[idx])
// 		dp[idx][0] = curmax
// 	}
// 	var max_diff int
// 	for day := 1; day < d; day++ {
// 		for idx := day; idx < n; idx++ {
// 			max_diff = 0
// 			dp[idx][day] = 300_001
// 			for j := idx; j >= day; j-- {
// 				max_diff = max(max_diff, jobDifficulty[j])
// 				dp[idx][day] = min(dp[idx][day], max_diff+dp[j-1][day-1])
// 			}
// 		}
// 	}
// 	return dp[n-1][d-1]
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(N^2 * D)
// // Space: O(N*D)
// func minDifficulty(jobDifficulty []int, d int) int {
// 	var n = len(jobDifficulty)
// 	if n < d {
// 		return -1
// 	}
// 	var dp = make([][]*int, n)
// 	for idx := range dp {
// 		dp[idx] = make([]*int, d+1)
// 	}
// 	var maxRigth = make([]int, n+1)
// 	for idx := n - 1; idx >= 0; idx-- {
// 		maxRigth[idx] = max(maxRigth[idx+1], jobDifficulty[idx])
// 	}
// 	var dfs func(idx int, day int) int
// 	dfs = func(idx int, day int) int {
// 		if day == 1 {
// 			return maxRigth[idx]
// 		}
// 		if dp[idx][day] != nil {
// 			return *dp[idx][day]
// 		}
// 		var maxDiff, minDiff = 0, 300_001
// 		for j := idx; j < n-day+1; j++ {
// 			maxDiff = max(maxDiff, jobDifficulty[j])
// 			minDiff = min(minDiff, maxDiff+dfs(j+1, day-1))
// 		}
// 		dp[idx][day] = &minDiff
// 		return minDiff
// 	}
// 	return dfs(0, d)
// }
