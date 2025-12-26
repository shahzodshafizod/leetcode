package maths

// https://leetcode.com/problems/count-number-of-trapezoids-ii/

// Approach 2: Optimized - Hash by Slope and Intercept
// Key insights from hints:
// 1. Hash every point-pair by its slope
// 2. For each slope, group segments by their line (using intercept)
// 3. Count pairs of segments with same slope but different lines
// 4. Subtract overcounts from parallelograms (counted twice)
//
// Algorithm:
// - For each pair of points (segment), calculate:
//   - Slope (as floating point for simplicity)
//   - Intercept (to distinguish parallel lines)
//   - Midpoint (to identify parallelograms)
//
// - Count all pairs of parallel segments (potential trapezoids)
// - Subtract segments sharing endpoints
// - Subtract parallelogram overcounts
//
// Time Complexity: O(n^2) for processing pairs + O(n^2) for counting
// Space Complexity: O(n^2) for storing segment information
func countTrapezoids(points [][]int) int {
	n := len(points)
	const inf = 1e9 + 7

	// Map: slope -> list of intercepts
	slopeToIntercepts := make(map[float64][]float64)

	// Map: midpoint -> list of slopes
	midToSlopes := make(map[int][]float64)

	// Process all pairs of points (segments)
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x1 - x2
			dy := y1 - y2

			var k, b float64

			// Calculate slope
			if x2 == x1 {
				// Vertical line: use special marker
				k = inf
				b = float64(x1) // Intercept is x-coordinate
			} else {
				// Slope = (y2 - y1) / (x2 - x1)
				k = float64(y2-y1) / float64(x2-x1)
				// Intercept: y - slope * x
				// b = y1 - k * x1 = (y1 * dx - x1 * dy) / dx
				b = float64(y1*dx-x1*dy) / float64(dx)
			}

			// Midpoint of segment (to identify parallelograms)
			// Use encoding to avoid complex key types
			mid := (x1+x2)*10000 + (y1 + y2)

			slopeToIntercepts[k] = append(slopeToIntercepts[k], b)
			midToSlopes[mid] = append(midToSlopes[mid], k)
		}
	}

	ans := 0

	// Count all pairs of parallel segments (trapezoid candidates)
	for _, intercepts := range slopeToIntercepts {
		if len(intercepts) < 2 {
			continue
		}

		// Group by intercept (segments on same line don't form trapezoid)
		interceptCount := make(map[float64]int)
		for _, b := range intercepts {
			interceptCount[b]++
		}

		// Count pairs of segments on different parallel lines
		// For each new intercept, count pairs with all previous ones
		totalSum := 0
		for _, count := range interceptCount {
			ans += totalSum * count
			totalSum += count
		}
	}

	// Subtract parallelogram overcounts
	// Parallelograms are counted twice (once for each pair of parallel sides)
	for _, slopes := range midToSlopes {
		if len(slopes) < 2 {
			continue
		}

		// Group by slope
		slopeCount := make(map[float64]int)
		for _, k := range slopes {
			slopeCount[k]++
		}

		// Subtract pairs of segments with same slope and same midpoint
		totalSum := 0
		for _, count := range slopeCount {
			ans -= totalSum * count
			totalSum += count
		}
	}

	return ans
}

// // Approach 1: Brute Force - Check All Quadrilaterals
// // Try all combinations of 4 points and check if they form a trapezoid.
// // A trapezoid has at least one pair of parallel sides, but NOT all collinear.
// // Key: For any 4 points, there are exactly 3 ways to pair them into
// // opposite sides: (p1,p2)-(p3,p4), (p1,p3)-(p2,p4), (p1,p4)-(p2,p3)
// // If exactly one or two pairs have the same slope (not all three), it's a trapezoid.
// //
// // Time Complexity: O(n^4) where n = number of points
// // Space Complexity: O(1)
// // This approach will TLE for n = 500 (over 6 billion operations)
// func countTrapezoids(points [][]int) int {
// 	// Helper to calculate GCD
// 	gcd := func(a, b int) int {
// 		for b != 0 {
// 			a, b = b, a%b
// 		}
// 		return a
// 	}

// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}
// 		return x
// 	}

// 	// Helper to get normalized slope between two points
// 	getSlope := func(p1, p2 []int) (int, int) {
// 		dx := p2[0] - p1[0]
// 		dy := p2[1] - p1[1]

// 		if dx == 0 {
// 			return 1, 0 // Vertical line
// 		}
// 		if dy == 0 {
// 			return 0, 1 // Horizontal line
// 		}

// 		// Reduce fraction and normalize sign
// 		if dx < 0 {
// 			dx, dy = -dx, -dy
// 		}
// 		g := gcd(abs(dx), abs(dy))
// 		return dy / g, dx / g
// 	}

// 	// Check if 4 points form a trapezoid
// 	isTrapezoid := func(p1, p2, p3, p4 []int) bool {
// 		// For 4 points, check all 3 ways to pair them as opposite sides
// 		s12dy, s12dx := getSlope(p1, p2)
// 		s13dy, s13dx := getSlope(p1, p3)
// 		s14dy, s14dx := getSlope(p1, p4)
// 		s23dy, s23dx := getSlope(p2, p3)
// 		s24dy, s24dx := getSlope(p2, p4)
// 		s34dy, s34dx := getSlope(p3, p4)

// 		// Check if each pairing has parallel segments
// 		slope1234 := s12dy == s34dy && s12dx == s34dx // (p1,p2) || (p3,p4)
// 		slope1324 := s13dy == s24dy && s13dx == s24dx // (p1,p3) || (p2,p4)
// 		slope1423 := s14dy == s23dy && s14dx == s23dx // (p1,p4) || (p2,p3)

// 		// Count how many pairs are parallel
// 		parallelCount := 0
// 		if slope1234 {
// 			parallelCount++
// 		}
// 		if slope1324 {
// 			parallelCount++
// 		}
// 		if slope1423 {
// 			parallelCount++
// 		}

// 		// Trapezoid: at least one pair parallel, but NOT all collinear
// 		// If all 3 pairs are parallel, points are collinear (not a valid quadrilateral)
// 		return parallelCount > 0 && parallelCount < 3
// 	}

// 	n := len(points)
// 	count := 0

// 	// Try all combinations of 4 points
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			for k := j + 1; k < n; k++ {
// 				for l := k + 1; l < n; l++ {
// 					if isTrapezoid(points[i], points[j], points[k], points[l]) {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return count
// }
