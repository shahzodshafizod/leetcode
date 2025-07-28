package dp

// https://leetcode.com/problems/filling-bookcase-shelves/

// Approach: NeetCode's Bottom-Up Dynamic Programming (Tabulation)
// W = shelfWidth
// N = len(books)
// Time: O(N x W)
// Space: O(N)
func minHeightShelves(books [][]int, shelfWidth int) int {
	// books[i] = [2]int{thickness, height}
	n := len(books)
	dp := make([]int, n+1)

	for idx := n - 1; idx >= 0; idx-- {
		space := shelfWidth
		maxHeight := 0
		dp[idx] = 1_000 * shelfWidth // 1_000 is max height

		for jdx := idx; jdx < n && space >= books[jdx][0]; jdx++ {
			space -= books[jdx][0]
			maxHeight = max(maxHeight, books[jdx][1])
			dp[idx] = min(dp[idx], dp[jdx+1]+maxHeight)
		}
	}

	return dp[0]
}

// // Approach: NeetCode's Top-Down Dynamic Programming (Memoization)
// // W = shelfWidth
// // N = len(books)
// // Time: O(N x W)
// // Space: O(N)
// func minHeightShelves(books [][]int, shelfWidth int) int {
// // books[i] = [2]int{thickness, height}
// 	var n = len(books)
// 	var cache = make([]int, n)
// 	var dfs func(idx int) int
// 	dfs = func(idx int) int {
// 		if idx >= n {
// 			return 0
// 		}
// 		if cache[idx] != 0 {
// 			return cache[idx]
// 		}
// 		var space = shelfWidth
// 		var maxHeight = 0
// 		cache[idx] = math.MaxInt
// 		for jdx := idx; jdx < n && space >= books[jdx][0]; jdx++ {
// 			space -= books[jdx][0]
// 			maxHeight = max(maxHeight, books[jdx][1])
// 			cache[idx] = min(
// 				cache[idx],
// 				maxHeight+dfs(jdx+1),
// 			)
// 		}
// 		return cache[idx]
// 	}
// 	return dfs(0)
// }

// // Approach: Top-Down Dynamic Programming (Memoization)
// // N = len(books)
// // W = shelfWidth
// // Time: O(N x W)
// // Space: O(N x W)
// func minHeightShelves(books [][]int, shelfWidth int) int {
// 	// books[i] = [2]int{thickness, height}
// 	var n = len(books)
// 	var cache = make([][]int, n)
// 	for idx := range cache {
// 		cache[idx] = make([]int, shelfWidth+1)
// 	}
// 	var dfs func(idx int, space int, maxHeight int) int
// 	dfs = func(idx int, space int, maxHeight int) int {
// 		if idx == n {
// 			return maxHeight
// 		}
// 		if cache[idx][space] != 0 {
// 			return cache[idx][space]
// 		}
// 		// decision NOT to include: create a new shelf
// 		var height = maxHeight + dfs(idx+1, shelfWidth-books[idx][0], books[idx][1])
// 		// decision to include
// 		if space >= books[idx][0] {
// 			height = min(height, dfs(idx+1, space-books[idx][0], max(maxHeight, books[idx][1])))
// 		}
// 		cache[idx][space] = height
// 		return height
// 	}
// 	return dfs(0, shelfWidth, 0)
// }
