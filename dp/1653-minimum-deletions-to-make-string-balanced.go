package dp

// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/

// Approach #5: Optimized DP
// Time: O(N)
// Space: O(1)
func minimumDeletions(s string) int {
	var deletions = 0
	var bcount = 0
	for _, c := range s {
		if c == 'b' {
			bcount++
		} else {
			deletions = min(deletions+1, bcount)
		}
	}
	return deletions
}

// // Approach #4: Using DP (One Pass)
// // Time: O(N)
// // Space: O(N)
// func minimumDeletions(s string) int {
// 	var n = len(s)
// 	var dp = make([]int, n+1)
// 	var bcount = 0
// 	// dp[i]: The number of deletions required to
// 	// balance the substring s[0, i)
// 	for idx := 0; idx < n; idx++ {
// 		if s[idx] == 'b' {
// 			dp[idx+1] = dp[idx]
// 			bcount++
// 		} else {
// 			// Two cases: remove 'a' or keep 'a'
// 			dp[idx+1] = min(dp[idx]+1, bcount)
// 		}
// 	}
// 	return dp[n]
// }

// // Approach #3: Using stack (one pass)
// // Time: O(N)
// // Space: O(N)
// func minimumDeletions(s string) int {
// 	var stack = design.NewStack[rune]()
// 	var deleteCount = 0
// 	for _, c := range s {
// 		if c == 'a' && !stack.Empty() && stack.Top() == 'b' {
// 			deleteCount++
// 			stack.Pop()
// 		} else {
// 			stack.Push(c)
// 		}
// 	}
// 	return deleteCount
// }

// // Approach #2: Two-Variable Method
// // Time: O(2N) = O(N)
// // Space: O(1)
// func minimumDeletions(s string) int {
// 	var acount = 0 // strings.Count(s, "a")
// 	for _, c := range s {
// 		if c == 'a' {
// 			acount++
// 		}
// 	}
// 	var bcount = 0
// 	var res = len(s)
// 	for _, c := range s {
// 		if c == 'a' {
// 			acount--
// 		}
// 		res = min(res, acount+bcount)
// 		if c == 'b' {
// 			bcount++
// 		}
// 	}
// 	return res
// }

// // Approach #1: Three-Pass Count Method
// // Time: O(3N) = O(N)
// // Space: O(2n) = O(N)
// func minimumDeletions(s string) int {
// 	var n = len(s)
// 	var leftBs = make([]int, n)
// 	var count = 0
// 	for idx := 0; idx < n; idx++ {
// 		leftBs[idx] = count
// 		if s[idx] == 'b' {
// 			count++
// 		}
// 	}
// 	var rightAs = make([]int, n)
// 	count = 0
// 	for idx := n - 1; idx >= 0; idx-- {
// 		rightAs[idx] = count
// 		if s[idx] == 'a' {
// 			count++
// 		}
// 	}
// 	var res = n
// 	for idx := 0; idx < n; idx++ {
// 		res = min(res, leftBs[idx]+rightAs[idx])
// 	}
// 	return res
// }
