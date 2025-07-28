package dp

// https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/

// Approach: Top-Down Dynamic Programming
// Time: O(logn)
// Space: O(logn)
func minDays(n int) int {
	// We have three options:
	// - east 1, remains n-1: 1+n-1 = n
	// - east n/2, remains n/2, n/2+n/2 = n
	// - east 2*n/3, remains n/3, 2*n/3+n/3 = 3*n/3 = n
	memo := make(map[int]int)

	var dp func(n int) int

	dp = func(n int) int {
		if n <= 1 {
			return n
		}

		if cache, found := memo[n]; found {
			return cache
		}

		memo[n] = 1 + min(
			n%3+dp(n/3), // make n divisable to 3 and add the remaining
			n%2+dp(n/2), // make n divisable to 2 and add the remaining
		)

		return memo[n]
	}

	return dp(n)
}

// // Approach: Breadth-First Search
// // Time: O(logn)
// // Space: O(n)
// func minDays(n int) int {
// 	var queue = []int{n}
// 	var seen = map[int]bool{n: true}
// 	var days = 0
// 	for size := len(queue); size > 0; size = len(queue) {
// 		for idx := 0; idx < size; idx++ {
// 			n = queue[idx]
// 			if n == 0 {
// 				return days
// 			}
// 			if n%3 == 0 && !seen[n/3] {
// 				queue = append(queue, n/3)
// 				seen[n/3] = true
// 			}
// 			if n%2 == 0 && !seen[n/2] {
// 				queue = append(queue, n/2)
// 				seen[n/2] = true
// 			}
// 			if !seen[n-1] {
// 				queue = append(queue, n-1)
// 				seen[n-1] = true
// 			}
// 		}
// 		queue = queue[size:]
// 		days++
// 	}
// 	return -1
// }
