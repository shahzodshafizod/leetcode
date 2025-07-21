package dp

// https://leetcode.com/problems/string-compression-ii/

// Approach: Top-Down Dynamic Programming
// Time: O(nnk)
// Space: O(nk)
func getLengthOfOptimalCompression(s string, k int) int {
	// ''->'a', 'a'->'a2', 'a9'->'a10', 'a99'->'a100'
	incr := map[int]int{1: 1, 2: 1, 10: 1, 100: 1}
	n := len(s)
	memo := make([][]*int, n)
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
		// delete s[idx]
		res := dp(idx+1, k-1)
		// keep s[idx] and same characters
		diff, same, length := 0, 0, 0
		for j := idx; j < n; j++ {
			if s[j] == s[idx] {
				same++
				length += incr[same]
			} else {
				diff++
				if diff > k {
					break
				}
			}
			res = min(res, length+dp(j+1, k-diff))
		}
		memo[idx][k] = &res
		return res
	}
	return dp(0, k)
}
