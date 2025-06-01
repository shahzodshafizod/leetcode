package maths

// https://leetcode.com/problems/distribute-candies-among-children-ii/

// Approach 2: Inclusion-Exclusion Principle
// Time: O(1)
// Space: O(1)
func distributeCandies2929(n int, limit int) int64 {
	var count = func(candies int) int64 {
		// to avoid selecting 3 children, we add 2 more dividers
		candies += 2
		if candies < 0 {
			return 0
		}
		return int64(candies) * int64(candies-1) / 2
	}

	// step 1: calculate number of ways to
	// distribute n candies without any limit
	var ways = count(n)

	// step 2: subtract the ways where at
	// least 1 child receives more than limit
	ways -= 3 * count(n-(limit+1))

	// step 3: in the step (2), we may subtract some
	// distributions multiple times - specifically, the ones
	// where two or more children receive more than limit candies.
	ways += 3 * count(n-2*(limit+1))

	// step 4: in the step (3), we've overcounted
	// the distributions where all three children
	// exceed the limit, so we subtract those again.
	ways -= count(n - 3*(limit+1))

	return ways
}

// // Approach 1: Enumeration
// // Time: O(min(n, limit))
// // Space: O(1)
// func distributeCandies2929(n int, limit int) int64 {
// 	var ways int64 = 0
// 	for i := min(n, limit); i >= 0; i-- {
// 		if n-i <= 2*limit {
// 			ways += int64(min(n-i, limit) - max(0, n-i-limit) + 1)
// 		}
// 	}
// 	return ways
// }
