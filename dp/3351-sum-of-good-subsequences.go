package dp

// https://leetcode.com/problems/sum-of-good-subsequences/

// // Approach 1: Brute Force (Top-Down Dynamic Programming)
// // Time: O(nn)
// // Space: O(nn)
// func sumOfGoodSubsequences(nums []int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	type pair struct {
// 		length int
// 		total  int
// 	}

// 	const mod int = 1e9 + 7

// 	n := len(nums)

// 	memo := make([][]*pair, n)
// 	for i := range memo {
// 		memo[i] = make([]*pair, n)
// 	}

// 	var dp func(prev int, curr int) *pair

// 	dp = func(prev int, curr int) *pair {
// 		if curr == n {
// 			return &pair{}
// 		}

// 		if memo[prev+1][curr] != nil {
// 			return memo[prev+1][curr]
// 		}

// 		pr := dp(prev, curr+1)
// 		if prev == -1 || abs(nums[prev]-nums[curr]) == 1 {
// 			npr := dp(curr, curr+1)
// 			pr.total = (pr.total + nums[curr]*(npr.length+1) + npr.total) % mod
// 			pr.length = (pr.length + npr.length + 1) % mod
// 		}

// 		memo[prev+1][curr] = pr

// 		return pr
// 	}

// 	return dp(-1, 0).total
// }

// Approach 2: Bottom-Up Dynamic Programming (Tabulation)
// Time: O(n)
// Space: O(n)
func sumOfGoodSubsequences(nums []int) int {
	const mod int = 1e9 + 7

	sum := make(map[int]int)
	cnt := make(map[int]int)

	var total int

	for _, num := range nums {
		numCnt := (1 + cnt[num-1] + cnt[num+1]) % mod
		numSum := (num*numCnt + sum[num-1] + sum[num+1]) % mod

		cnt[num] = (cnt[num] + numCnt) % mod
		sum[num] = (sum[num] + numSum) % mod
		total = (total + numSum) % mod
	}

	return total
}
