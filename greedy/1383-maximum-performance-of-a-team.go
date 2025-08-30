package greedy

import (
	"container/heap"
	"slices"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/maximum-performance-of-a-team/

// Time: O(n Log n)
// Space: O(n)
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engineers := make([][2]int, 0, n)
	for idx := range n {
		engineers = append(engineers, [2]int{efficiency[idx], speed[idx]})
	}

	slices.SortFunc(engineers, func(a, b [2]int) int { return b[0] - a[0] })

	speeds := pkg.NewHeap(
		make([]int, 0),
		func(x, y int) bool { return x < y },
	)
	performance, totalSpeed := 0, 0

	for _, engineer := range engineers {
		if speeds.Len() == k {
			speed, ok := heap.Pop(speeds).(int)
			_ = ok
			totalSpeed -= speed
		}

		totalSpeed += engineer[1]
		heap.Push(speeds, engineer[1])
		performance = max(performance, totalSpeed*engineer[0])
	}

	return performance % (1e9 + 7)
}

// // Approach: Brute-Force
// // Time: O(2^N)
// func maxPerformance(n int, speed []int, efficiency []int, k int) int {
// 	var dp func(idx int, count int, spd int, eff int) int
// 	dp = func(idx int, count int, spd int, eff int) int {
// 		if idx == n || count == k {
// 			return spd * eff
// 		}
// 		return max(
// 			// decision to include
// 			dp(idx+1, count+1, spd+speed[idx], min(eff, efficiency[idx])),
// 			// decision NOT to include
// 			dp(idx+1, count, spd, eff),
// 		)
// 	}
// 	return dp(0, 0, 0, 1e8) % (1e9 + 7)
// }
