package graphs

// https://leetcode.com/problems/parallel-courses-ii/

// Approach #1: Top-Down Dynamic Programming
// Time: O(2^n)
// Space: O(2^n)
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	adjList := make([][]int, n)
	var prev, next int
	for idx := range relations {
		prev, next = relations[idx][0]-1, relations[idx][1]-1
		adjList[prev] = append(adjList[prev], next)
	}
	memo := make([]*int, 1<<n)
	var dfs func(mask int) int
	dfs = func(mask int) int {
		if mask == (1<<n)-1 {
			return 0
		}
		if memo[mask] != nil {
			return *memo[mask]
		}
		indegree := make([]int, n)
		for node := 0; node < n; node++ {
			if mask&(1<<node) != 0 {
				continue
			}
			for _, child := range adjList[node] {
				indegree[child]++
			}
		}
		availableMask := 0
		count := 0
		for node := 0; node < n; node++ {
			if indegree[node] == 0 && mask&(1<<node) == 0 {
				availableMask |= 1 << node
				count++
			}
		}
		result := n + 1
		nextMask := availableMask
		if count > k {
			for nextMask > 0 {
				count = 0
				tmp := nextMask
				for tmp > 0 {
					tmp &= tmp - 1
					count++
				}
				if count != k {
					nextMask = (nextMask - 1) & availableMask
					continue
				}
				result = min(result, 1+dfs(nextMask|mask))
				nextMask = (nextMask - 1) & availableMask
			}
		} else {
			result = min(result, 1+dfs(nextMask|mask))
		}
		memo[mask] = &result
		return result
	}
	return dfs(0)
}
