package greedy

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-cost-to-hire-k-workers/

// time: O(nlogn + nlogk)
// space: O(n + k)
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	type candidate struct {
		quality int
		ratio   float64
	}
	var candidates = pkg.NewPQ(
		make([]*candidate, 0),
		func(x, y *candidate) bool { return x.ratio > y.ratio },
	)
	for idx := range quality { // O(n)
		candidates.Push(&candidate{
			quality: quality[idx],
			ratio:   float64(wage[idx]) / float64(quality[idx]),
		})
	}
	var money float64 = math.MaxFloat64
	var top *candidate
	var qualities float64 = 0
	var maxheap = pkg.NewPQ(make([]int, 0), func(x, y int) bool { return x < y })
	for candidates.Len() > 0 { // O(n log k)
		top = candidates.Pop()
		qualities += float64(top.quality)
		maxheap.Push(top.quality)
		if maxheap.Len() > k {
			qualities -= float64(maxheap.Pop())
		}
		if maxheap.Len() == k {
			money = min(money, qualities*top.ratio)
		}
	}
	return money
}

// // Approach: Brute Force
// // time: O(n^3 * log n)
// // space: O(n)
// func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
// 	var ratio, offer float64
// 	var money float64
// 	var minmoney = math.MaxFloat64
// 	for captain, n := 0, len(quality); captain < n; captain++ {
// 		ratio = float64(wage[captain]) / float64(quality[captain])
// 		var chosen = design.NewPQ(make([]float64, 0), func(x, y float64) bool { return x > y })
// 		for candidate := 0; candidate < n; candidate++ {
// 			offer = float64(quality[candidate]) * ratio
// 			if offer >= float64(wage[candidate]) {
// 				chosen.Push(offer)
// 			}
// 		}
// 		if chosen.Len() >= k {
// 			money = 0
// 			for j := 0; j < k; j++ {
// 				money += chosen.Pop()
// 			}
// 			minmoney = min(minmoney, money)
// 		}
// 	}
// 	return minmoney
// }
