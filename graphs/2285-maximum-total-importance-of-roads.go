package graphs

import "slices"

// https://leetcode.com/problems/maximum-total-importance-of-roads/

// Approach #2: Counting (Bucket) Sorting
// R = len(roads) = N^2
// time: O(R+3N) = O(R+N) = O(N^2)
// space: O(2N) = O(N)
func maximumImportance(n int, roads [][]int) int64 {
	degrees := make([]int, n)
	for _, road := range roads { // O(R)
		degrees[road[0]]++
		degrees[road[1]]++
	}
	max := slices.Max(degrees) // O(N)
	count := make([]int, max+1)
	for _, degree := range degrees { // O(N)
		count[degree]++
	}
	var importance int64 = 0
	for degree := max; degree >= 0; degree-- { // O(MAX) || O(N)
		for count[degree] > 0 {
			importance += int64(degree * n)
			count[degree]--
			n--
		}
	}
	return importance
}

// // Approach #1: Sorting + Greedy
// // R = len(roads) = N^2
// // time: O(2R+NLogN) = O(R+NLogN) = O(N^2)
// // space: O(N)
// func maximumImportance(n int, roads [][]int) int64 {
// 	var degrees = make([]int, n)
// 	for _, road := range roads { // O(R)
// 		degrees[road[0]]++
// 		degrees[road[1]]++
// 	}
// 	sort.Ints(degrees) // O(N Log N)
// 	var importance int64 = 0
// 	for idx, count := range degrees { // O(R)
// 		importance += int64((idx + 1) * count)
// 	}
// 	return importance
// }
