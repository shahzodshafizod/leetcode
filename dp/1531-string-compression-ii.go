package dp

// https://leetcode.com/problems/string-compression-ii/

// Approach: Top-Down Dynamic Programming
// Time: O(nnk)
// Space: O(nk)
func getLengthOfOptimalCompression(s string, k int) int {
	var n = len(s)
	var memo = make([][]*int, n)
	for idx := range memo {
		memo[idx] = make([]*int, k+1)
	}
	var dp func(idx int, k int) int
	dp = func(idx int, k int) int {
		if idx+k >= n {
			return 0
		}
		if k < 0 {
			return 101
		}
		if memo[idx][k] != nil {
			return *memo[idx][k]
		}
		// skip this character
		var res = dp(idx+1, k-1)
		var diff, same, length = 0, 0, 0
		// for all continuous s[i] characters (can skip k characters)
		for j := idx; j < n; j++ {
			if k-diff < 0 {
				break
			}
			if s[j] == s[idx] {
				same++
				if same <= 2 || same == 10 || same == 100 {
					length++
				}
			} else {
				diff++
			}
			res = min(res, length+dp(j+1, k-diff))
		}
		memo[idx][k] = &res
		return res
	}
	return dp(0, k)
}
