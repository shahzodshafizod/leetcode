package backtracking

import (
	"fmt"
	"math/bits"
)

// https://leetcode.com/problems/binary-watch/

// Approach 1: Brute Force - Try All Possible Times
// Iterate through all possible hours (0-11) and minutes (0-59).
// For each combination, count the number of bits set.
// If it matches turnedOn, add to result.
// Time: O(12 * 60) = O(720) = O(1) - constant time
// Space: O(1) - not counting output

// Approach 2: Backtracking
// Use backtracking to generate all combinations of positions
// where LEDs can be turned on.
// Time: O(C(10, turnedOn)) - combinatorial
// Space: O(turnedOn) - recursion depth

// Approach 3: Bit Manipulation (Optimal for this problem)
// Count set bits in hour (0-11) and minute (0-59).
// If total set bits equals turnedOn, add to result.
// This is actually simpler than backtracking for this specific problem.
// Time: O(12 * 60) = O(1)
// Space: O(1)
func readBinaryWatch(turnedOn int) []string {
	result := []string{}

	// Try all possible hours (0-11) and minutes (0-59)
	for hour := range 12 {
		for minute := range 60 {
			// Count total bits set in both hour and minute
			if bits.OnesCount(uint(hour))+bits.OnesCount(uint(minute)) == turnedOn {
				// Format: h:mm (hour without leading zero, minute with leading zero)
				result = append(result, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}
	}

	return result
}

// // Alternative: Backtracking approach
// func readBinaryWatch(turnedOn int) []string {
// 	result := []string{}

// 	// positions[i] represents the value at position i
// 	// First 4 positions are hours: 8, 4, 2, 1
// 	// Last 6 positions are minutes: 32, 16, 8, 4, 2, 1
// 	positions := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}

// 	var backtrack func(index, count, hours, minutes int)

// 	backtrack = func(index, count, hours, minutes int) {
// 		// Pruning: if already invalid, stop
// 		if hours >= 12 || minutes >= 60 {
// 			return
// 		}

// 		// Base case: used all turnedOn LEDs
// 		if count == turnedOn {
// 			result = append(result, fmt.Sprintf("%d:%02d", hours, minutes))

// 			return
// 		}

// 		// Try turning on each remaining LED
// 		for i := index; i < len(positions); i++ {
// 			if i < 4 {
// 				// Hour LED
// 				backtrack(i+1, count+1, hours+positions[i], minutes)
// 			} else {
// 				// Minute LED
// 				backtrack(i+1, count+1, hours, minutes+positions[i])
// 			}
// 		}
// 	}

// 	backtrack(0, 0, 0, 0)

// 	return result
// }
