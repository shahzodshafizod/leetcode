package dp

// https://leetcode.com/problems/count-all-possible-routes/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n^2 * fuel)
// Space: O(n * fuel)
func countRoutes(locations []int, start int, finish int, fuel int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	const mod = 1e9 + 7

	n := len(locations)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, fuel+1)
	}

	dp[start][fuel] = 1

	for f := fuel; f >= 0; f-- {
		for city := range n {
			if dp[city][f] == 0 {
				continue
			}

			for nextCity := range n {
				if nextCity == city {
					continue
				}

				remainingFuel := f - abs(locations[city]-locations[nextCity])

				if remainingFuel >= 0 {
					dp[nextCity][remainingFuel] = (dp[nextCity][remainingFuel] + dp[city][f]) % mod
				}
			}
		}
	}

	result := 0
	for f := 0; f <= fuel; f++ {
		result = (result + dp[finish][f]) % mod
	}

	return result
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n^2 * fuel)
// // Space: O(n * fuel)
// func countRoutes(locations []int, start int, finish int, fuel int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	const mod = 1e9 + 7

// 	n := len(locations)

// 	memo := make([][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([]*int, fuel+1)
// 	}

// 	var dfs func(city, fuel int) int

// 	dfs = func(city, fuel int) int {
// 		if memo[city][fuel] != nil {
// 			return *memo[city][fuel]
// 		}

// 		count := 0
// 		if city == finish {
// 			count = 1
// 		}

// 		for nextCity := range n {
// 			if nextCity == city {
// 				continue
// 			}

// 			remainingFuel := fuel - abs(locations[city]-locations[nextCity])
// 			if remainingFuel >= 0 {
// 				count = (count + dfs(nextCity, remainingFuel)) % mod
// 			}
// 		}

// 		memo[city][fuel] = &count

// 		return count
// 	}

// 	return dfs(start, fuel)
// }
