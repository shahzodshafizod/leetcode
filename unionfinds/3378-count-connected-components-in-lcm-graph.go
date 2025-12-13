package unionfinds

// https://leetcode.com/problems/count-connected-components-in-lcm-graph/

// Approach #2: Union-Find with Multiples Connection
// Key insight: For each value that exists in nums and is <= threshold,
// connect it with all its multiples (2*val, 3*val, ...) that are also <= threshold.
// This works because LCM(a, k*a) = k*a, so if k*a <= threshold, they connect.
// Time: O(threshold * log(threshold) + n) - harmonic series sum
// Space: O(threshold) for DSU structure
func countComponents(nums []int, threshold int) int {
	// DSU operates on values up to threshold
	parent := make([]int, threshold+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	union := func(x, y int) {
		px := find(x)

		py := find(y)
		if px != py {
			parent[py] = px
		}
	}

	// For each value in nums that's <= threshold, connect it with all multiples
	// Time complexity: O(threshold/1 + threshold/2 + ... + threshold/k) = O(threshold * log(threshold))
	for _, num := range nums {
		if num <= threshold {
			// Connect num with all its multiples: 2*num, 3*num, etc.
			for multiple := 2 * num; multiple <= threshold; multiple += num {
				union(num, multiple)
			}
		}
	}

	// Count distinct components
	components := make(map[int]bool)
	largeCount := 0

	for _, num := range nums {
		if num > threshold {
			// Each number > threshold is its own component
			largeCount++
		} else {
			components[find(num)] = true
		}
	}

	return len(components) + largeCount
}

// // Approach #1: Brute Force - Check all pairs with Union Find
// // For each pair, compute LCM and connect if LCM <= threshold
// // Time: O(n^2 * log(max(nums))) for GCD computation
// // Space: O(n) for union find
// func countComponents(nums []int, threshold int) int {
// 	gcd := func(a, b int) int {
// 		for b != 0 {
// 			a, b = b, a%b
// 		}

// 		return a
// 	}

// 	lcm := func(a, b int) int {
// 		return a * b / gcd(a, b)
// 	}

// 	n := len(nums)

// 	parent := make([]int, n)
// 	for i := range n {
// 		parent[i] = i
// 	}

// 	var find func(int) int

// 	find = func(x int) int {
// 		if parent[x] != x {
// 			parent[x] = find(parent[x])
// 		}

// 		return parent[x]
// 	}

// 	union := func(x, y int) bool {
// 		px := find(x)

// 		py := find(y)
// 		if px == py {
// 			return false
// 		}

// 		parent[py] = px

// 		return true
// 	}

// 	// Check all pairs
// 	for i := range n {
// 		for j := i + 1; j < n; j++ {
// 			if lcm(nums[i], nums[j]) <= threshold {
// 				union(i, j)
// 			}
// 		}
// 	}

// 	// Count distinct components
// 	components := make(map[int]bool)
// 	for i := range n {
// 		components[find(i)] = true
// 	}

// 	return len(components)
// }
