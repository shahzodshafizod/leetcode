package dp

import (
	"math"
	"strings"
)

// https://leetcode.com/problems/freedom-trail/

// Memoization
// time: O(k * r^2)
// space: O(r * k)
func findRotateSteps(ring string, key string) int {
	var rn, kn = len(ring), len(key)
	var cache = make([][]*int, rn)
	for idx := range cache {
		cache[idx] = make([]*int, kn)
	}
	var dfs func(r int, k int) int
	dfs = func(r int, k int) int {
		if k == kn {
			return 0
		}
		if cache[r][k] != nil {
			return *cache[r][k]
		}
		var result = math.MaxInt
		var dist, idx int
		var skips = 0
		for idx = strings.IndexByte(ring, key[k]); idx >= 0; {
			idx += skips
			dist = int(math.Abs(float64(idx - r)))
			dist = min(dist, rn-dist)
			result = min(result, dist+1+dfs(idx, k+1))
			skips = idx + 1
			idx = strings.IndexByte(ring[skips:], key[k])
		}
		cache[r][k] = &result
		return result
	}
	return dfs(0, 0)
}

// // Tabulation
// // time: O(kn * rn^2)
// // space: O(rn)
// func findRotateSteps(ring string, key string) int {
// 	var rn = len(ring)
// 	var dp = make([]int, rn)
// 	var makeinfarr = func(n int) []int {
// 		var arr = make([]int, n)
// 		for idx := range arr {
// 			arr[idx] = math.MaxInt
// 		}
// 		return arr
// 	}
// 	var dist int
// 	for k := len(key) - 1; k >= 0; k-- {
// 		var nextdp = makeinfarr(rn)
// 		for r := range ring {
// 			for idx := range ring {
// 				if ring[idx] == key[k] {
// 					dist = int(math.Abs(float64(r - idx)))
// 					dist = min(dist, rn-dist)
// 					nextdp[r] = min(nextdp[r], dist+1+dp[idx])
// 				}
// 			}
// 		}
// 		dp = nextdp
// 	}
// 	return dp[0]
// }
