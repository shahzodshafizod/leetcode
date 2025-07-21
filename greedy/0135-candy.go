package greedy

// https://leetcode.com/problems/candy/

// Approach #2: Greedy
// Time: O(n)
// Space: O(n)
func candy(ratings []int) int {
	n := len(ratings)
	cnts := make([]int, n)
	for idx := 1; idx < n; idx++ {
		if ratings[idx] > ratings[idx-1] {
			cnts[idx] = cnts[idx-1] + 1
		}
	}
	candies := n + cnts[n-1]
	for idx := n - 2; idx >= 0; idx-- {
		if ratings[idx] > ratings[idx+1] {
			cnts[idx] = max(cnts[idx], cnts[idx+1]+1)
		}
		candies += cnts[idx]
	}
	return candies
}

// // Approach #1: Sorting
// // Time: O(nlogn)
// // Space: O(n)
// func candy(ratings []int) int {
// 	var n = len(ratings)
// 	var indices = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		indices[idx] = idx
// 	}
// 	sort.Slice(indices, func(i, j int) bool {
// 		return ratings[indices[i]] < ratings[indices[j]]
// 	})
// 	var counts = make([]int, n)
// 	var candies = n
// 	for _, idx := range indices {
// 		if idx > 0 && ratings[idx] > ratings[idx-1] {
// 			counts[idx] = counts[idx-1] + 1
// 		}
// 		if idx+1 < n && ratings[idx] > ratings[idx+1] {
// 			counts[idx] = max(counts[idx], counts[idx+1]+1)
// 		}
// 		candies += counts[idx]
// 	}
// 	return candies
// }
