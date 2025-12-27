package graphs

// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-ii/

// Approach: O(n) with Difference Array and Exact Mathematical Formulas
// Key insight: For each house i, apply formulas that mark distance changes in O(1).
// The formulas account for paths using the x-y shortcut and adjust for overcounting.
//
// Mathematical Analysis:
// For each house i (0-indexed):
// 1. Add paths via shortcut to x and y
// 2. Subtract overcounts at left/right boundaries
// 3. Subtract cycle adjustment for even/odd paths
//
// After all updates, apply prefix sum to convert difference array to final counts.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func countOfPairs(n int, x int, y int) []int64 {
	// Convert to 0-indexed
	x--
	y--

	// Ensure x <= y
	if x > y {
		x, y = y, x
	}

	diff := make([]int64, n)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}

		return a
	}

	// For each house i, apply the exact mathematical formulas
	for i := range n {
		// Base case: every house connects to neighbors
		diff[0] += 2

		// Formula 1: Distance via shortcut from house x
		diff[min(abs(i-x), abs(i-y)+1)]++

		// Formula 2: Distance via shortcut from house y
		diff[min(abs(i-y), abs(i-x)+1)]++

		// Formula 3: Subtract overcount to left endpoint (house 0)
		diff[min(abs(i-0), abs(i-y)+1+abs(x-0))]--

		// Formula 4: Subtract overcount to right endpoint (house n-1)
		diff[min(abs(i-(n-1)), abs(i-x)+1+abs(y-(n-1)))]--

		// Formula 5: Subtract cycle adjustment for even path
		diff[max(x-i, 0)+max(i-y, 0)+(y-x)/2]--

		// Formula 6: Subtract cycle adjustment for odd path
		diff[max(x-i, 0)+max(i-y, 0)+(y-x+1)/2]--
	}

	// Apply difference array via prefix sum
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	return diff
}
