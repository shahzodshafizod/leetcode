package arrays

// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/

// Approach #2: Binary Search
// time: O(N∗Log(Max(Bloomday)))
// space: O(1)
func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	getNumOfBouquets := func(max int) int {
		flowers, bouquets := 0, 0
		for _, day := range bloomDay { // O(N)
			if day <= max {
				flowers++
			} else {
				flowers = 0
			}
			if flowers == k {
				bouquets++
				flowers = 0
			}
		}
		return bouquets
	}
	start, end := bloomDay[0], bloomDay[0]
	for _, day := range bloomDay { // O(N)
		start = min(start, day)
		end = max(end, day)
	}
	days := -1
	var mid int
	for start <= end { // O(Log Max(bloomDay))
		mid = (start + end) / 2
		if getNumOfBouquets(mid) >= m { // O(N)
			days = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return days
}

// // Approach #1: Brute Force
// // time: O(n^2)
// // space: O(1)
// func minDays(bloomDay []int, m int, k int) int {
// 	var days = -1
// 	for _, max := range bloomDay { // O(n)
// 		var flowers, bouquets = 0, 0
// 		for _, day := range bloomDay { // O(n)
// 			if day <= max {
// 				flowers++
// 			} else {
// 				flowers = 0
// 			}
// 			if flowers == k {
// 				bouquets++
// 				flowers = 0
// 			}
// 		}
// 		if bouquets >= m {
// 			if days == -1 || max < days {
// 				days = max
// 			}
// 		}
// 	}
// 	return days
// }
