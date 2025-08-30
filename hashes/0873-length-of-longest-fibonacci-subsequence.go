package hashes

// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/

// // Approach #1: Hash
// // Time: O(nn logm), n=len(arr), m=len(fib seq)
// // Space: O(n)
// func lenLongestFibSubseq(arr []int) int {
// 	var indices = make(map[int]bool, len(arr))
// 	for _, num := range arr {
// 		indices[num] = true
// 	}
// 	var longest = 0
// 	var prev, curr, next, length int
// 	for i, n := 0, len(arr); i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			prev, curr = arr[i], arr[j]
// 			next = prev + curr
// 			length = 2
// 			for indices[next] {
// 				length++
// 				longest = max(longest, length)
// 				prev, curr = curr, next
// 				next = prev + curr
// 			}
// 		}
// 	}
// 	return longest
// }

// Approach #2: Dynamic Programming
// Time: O(nn), n=len(arr)
// Space: O(nn)
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	indices := make(map[int]int, n)
	dp := make([][]int, n)

	for i := range n {
		indices[arr[i]] = i

		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}

	longest := 0

	var prev, curr, next int
	for prev = range n {
		for curr = prev + 1; curr < n; curr++ {
			next = indices[arr[prev]+arr[curr]]
			if next > curr {
				dp[curr][next] = max(
					dp[curr][next],
					dp[prev][curr]+1,
				)
				longest = max(longest, dp[curr][next])
			}
		}
	}

	return longest
}
