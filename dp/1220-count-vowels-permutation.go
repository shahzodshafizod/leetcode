package dp

// e -> a
// a -> e
// i -> e
// a -> i
// e -> i
// o -> i
// u -> i
// i -> o
// u -> o
// a -> u

// a:0, e:1, i:2, o:3, u:4

// https://leetcode.com/problems/count-vowels-permutation/

// Approach: Bottom-Up Dynamic Programming
// Time: O(N)
// Space: O(1)
func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7

	prev := [5]int{1, 1, 1, 1, 1}

	var curr [5]int

	count := 5

	for ; n > 1; n-- {
		curr[0] = (prev[1]) % mod
		curr[1] = (prev[0] + prev[2]) % mod
		curr[2] = (prev[0] + prev[1] + prev[3] + prev[4]) % mod
		curr[3] = (prev[2] + prev[4]) % mod
		curr[4] = (prev[0]) % mod
		count = 0

		for i := range 5 {
			prev[i] = curr[i]
			count = (count + curr[i]) % mod
		}
	}

	return count
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(N)
// // Space: O(N)
// func countVowelPermutation(n int) int {
// 	const mod int = 1e9 + 7
// 	var adjList = map[int][]int{
// 		-1: {0, 1, 2, 3, 4},
// 		0:  {1, 2, 4},
// 		1:  {0, 2},
// 		2:  {1, 3},
// 		3:  {2},
// 		4:  {2, 3},
// 	}
// 	var dp = make([][5]int, n)
// 	var dfs func(idx int, prev int) int
// 	dfs = func(idx int, prev int) int {
// 		if idx == n {
// 			return 1
// 		}
// 		if prev != -1 && dp[idx][prev] != 0 {
// 			return dp[idx][prev]
// 		}
// 		var count = 0
// 		for _, next := range adjList[prev] {
// 			count = (count + dfs(idx+1, next)) % mod
// 		}
// 		if prev != -1 {
// 			dp[idx][prev] = count
// 		}
// 		return count
// 	}
// 	return dfs(0, -1)
// }
