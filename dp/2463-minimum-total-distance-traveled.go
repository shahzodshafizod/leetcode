package dp

import (
	"sort"
)

// https://leetcode.com/problems/minimum-total-distance-traveled/

// Approach: Bottom-Up Dynamic Programming
// Time: O(RF), R=len(robot), F=sum(factory[i][1])
// Space: O(RF)
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	abs := func(x int) int64 {
		if x < 0 {
			x = -x
		}
		return int64(x)
	}
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	factories := make([]int, 0, len(factory))
	for fid := range factory {
		for limit := factory[fid][1]; limit > 0; limit-- {
			factories = append(factories, factory[fid][0])
		}
	}
	rlen, flen := len(robot), len(factories)
	dp := make([][]int64, rlen+1)
	for idx := range dp {
		dp[idx] = make([]int64, flen+1)
		dp[idx][flen] = 1e12
	}
	dp[rlen][flen] = 0
	var distance int64
	for rid := rlen - 1; rid >= 0; rid-- {
		for fid := flen - 1; fid >= 0; fid-- {
			distance = abs(robot[rid] - factories[fid])
			dp[rid][fid] = min(
				distance+dp[rid+1][fid+1], // include factory
				dp[rid][fid+1],            // skip factory
			)
		}
	}
	return dp[0][0]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(RF), R=len(robot), F=sum(factory[i][1])
// // Space: O(RF)
// func minimumTotalDistance(robot []int, factory [][]int) int64 {
// 	var abs = func(x int) int64 {
// 		if x < 0 {
// 			x = -x
// 		}
// 		return int64(x)
// 	}
// 	sort.Ints(robot)
// 	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
// 	var factories = make([]int, 0, len(factory))
// 	for fid := range factory {
// 		for limit := factory[fid][1]; limit > 0; limit-- {
// 			factories = append(factories, factory[fid][0])
// 		}
// 	}
// 	var rlen, flen = len(robot), len(factories)
// 	var memo = make([][]*int64, rlen)
// 	for idx := range memo {
// 		memo[idx] = make([]*int64, flen)
// 	}
// 	var dp func(rid int, fid int) int64
// 	dp = func(rid int, fid int) int64 {
// 		if rid == rlen {
// 			return 0
// 		}
// 		if fid == flen {
// 			return 1e12
// 		}
// 		if memo[rid][fid] != nil {
// 			return *memo[rid][fid]
// 		}
// 		var distance = abs(robot[rid] - factories[fid])
// 		memo[rid][fid] = new(int64)
// 		*memo[rid][fid] = min(
// 			distance+dp(rid+1, fid+1), // include factory
// 			dp(rid, fid+1),            // skip factory
// 		)
// 		return *memo[rid][fid]
// 	}
// 	return dp(0, 0)
// }
