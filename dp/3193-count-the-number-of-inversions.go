package dp

// https://leetcode.com/problems/count-the-number-of-inversions/

// // Approach 1: Brute-force - Generate all permutations and count inversions
// // Time: O(n! * n^2) - generate all permutations and count inversions for each
// // Space: O(n) - for storing permutation and recursion stack
// func numberOfPermutations(n int, requirements [][]int) int {
// 	const mod = int(1e9) + 7

// 	countInversions := func(perm []int, end int) int {
// 		count := 0

// 		for i := 0; i <= end; i++ {
// 			for j := i + 1; j <= end; j++ {
// 				if perm[i] > perm[j] {
// 					count++
// 				}
// 			}
// 		}

// 		return count
// 	}

// 	perm := make([]int, n)
// 	used := make([]bool, n)

// 	var backtrack func(pos int) int

// 	backtrack = func(pos int) int {
// 		if pos == n {
// 			// Check if this permutation satisfies all requirements
// 			for _, req := range requirements {
// 				end, cnt := req[0], req[1]
// 				if countInversions(perm, end) != cnt {
// 					return 0
// 				}
// 			}

// 			return 1
// 		}

// 		var count int

// 		for i := range n {
// 			if !used[i] {
// 				used[i] = true
// 				perm[pos] = i
// 				count = (count + backtrack(pos+1)) % mod

// 				used[i] = false
// 			}
// 		}

// 		return count
// 	}

// 	return backtrack(0)
// }

// // Approach 2: Top-down DP (Memoization)
// // Time: O(n * max_cnt * (n + max_cnt))
// // Space: O(n * max_cnt) for memoization
// func numberOfPermutations(n int, requirements [][]int) int {
// 	const mod = int(1e9) + 7

// 	reqs := make([]*int, n)

// 	var mxCnt, end, cnt int
// 	for _, req := range requirements {
// 		end, cnt = req[0], req[1]
// 		reqs[end] = new(int)
// 		*reqs[end] = cnt
// 		mxCnt = max(mxCnt, cnt)
// 	}

// 	// // Check if position 0 has impossible requirement
// 	// if cnt, exists := reqMap[0]; exists && cnt != 0 {
// 	// 	return 0
// 	// }

// 	memo := make([][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([]*int, mxCnt+1)
// 	}

// 	var dp func(i, cnt int) int

// 	dp = func(i, cnt int) int {
// 		if i == n {
// 			return 1
// 		}

// 		if cnt > mxCnt {
// 			return 0
// 		}

// 		if memo[i][cnt] != nil {
// 			return *memo[i][cnt]
// 		}

// 		var total int

// 		for ncnt := cnt; ncnt <= i+cnt; ncnt++ {
// 			if reqs[i] == nil || *reqs[i] == ncnt {
// 				total = (total + dp(i+1, ncnt)) % mod
// 			}
// 		}

// 		memo[i][cnt] = &total

// 		return total
// 	}

// 	return dp(0, 0)
// }

// Approach 3: Bottom-up DP (Tabulation)
// Time: O(n * maxInv^2) where maxInv <= 400
// Space: O(maxInv) with space optimization
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = int(1e9) + 7

	reqs := make([]*int, n)

	var end, cnt, mxcnt int
	for _, req := range requirements {
		end, cnt = req[0], req[1]
		reqs[end] = new(int)
		*reqs[end] = cnt
		mxcnt = max(mxcnt, cnt)
	}

	if reqs[0] != nil && *reqs[0] > 0 {
		return 0
	}

	dp := make([]int, mxcnt+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		prefix := make([]int, mxcnt+2)
		for cnt := range mxcnt + 1 {
			prefix[cnt+1] = (prefix[cnt] + dp[cnt]) % mod
		}

		ndp := make([]int, mxcnt+1)
		for cnt := range mxcnt + 1 {
			ndp[cnt] = (prefix[cnt+1] - prefix[max(0, cnt-(i-1))] + mod) % mod
		}

		if reqs[i-1] != nil {
			for cnt := range mxcnt + 1 {
				if cnt != *reqs[i-1] {
					ndp[cnt] = 0
				}
			}
		}

		dp = ndp
	}

	var total int
	for _, cnt := range dp {
		total = (total + cnt) % mod
	}

	return total
}
