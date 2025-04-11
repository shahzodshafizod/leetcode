package dp

import (
	"strconv"
)

// TODO: Solve it by YOURSELF!!!

// https://leetcode.com/problems/count-the-number-of-powerful-integers/

// Approach: Digital Dynamic Programming
// Time: O(len(finish))
// Space: O(len(finish))
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	var count = func(finish int64) int64 {
		var suff1, suff2 int64 = 0, 0
		var decimal int64 = 1
		for idx := len(s) - 1; idx >= 0; idx-- {
			if finish == 0 {
				return 0
			}
			suff1 += decimal * (finish % 10)
			suff2 += decimal * int64(s[idx]-'0')
			finish /= 10
			decimal *= 10
		}
		var flow = strconv.FormatInt(finish, 10)
		var n = len(flow)
		var dp = make([][2]int64, n+1)
		dp[n][0] = 1
		if suff1 >= suff2 {
			dp[n][1] = 1
		}
		for idx := n - 1; idx >= 0; idx-- {
			dp[idx][0] = int64(limit+1) * dp[idx+1][0]
			if int(flow[idx]-'0') <= limit {
				dp[idx][1] = int64(flow[idx]-'0')*dp[idx+1][0] + dp[idx+1][1]
			} else {
				dp[idx][1] = int64(limit+1) * dp[idx+1][0]
			}
		}
		return dp[0][1]
	}
	return count(finish) - count(start-1)
}
