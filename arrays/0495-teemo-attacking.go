package arrays

// https://leetcode.com/problems/teemo-attacking/

// Approach 1: Brute Force - Mark All Poisoned Seconds
// Create a boolean array to mark each second as poisoned or not.
// For each attack, mark the next 'duration' seconds as poisoned.
// Count total poisoned seconds.
// Time: O(n * duration) - could be very large if duration is big
// Space: O(max_time) - need array up to the last attack time + duration
// This approach would TLE for large inputs.

// // Approach 2: Optimal - Calculate Overlapping Intervals
// // Key insight: For each attack, add either 'duration' seconds or
// // the time until the next attack (whichever is smaller).
// // If there's overlap, the poison timer resets.
// //
// // Time: O(n) where n is length of timeSeries
// // Space: O(1)

// func findPoisonedDuration(timeSeries []int, duration int) int {
// 	if len(timeSeries) == 0 {
// 		return 0
// 	}

// 	totalTime := 0

// 	// Process each attack
// 	for i := range len(timeSeries) - 1 {
// 		// Calculate the time between current and next attack
// 		gap := timeSeries[i+1] - timeSeries[i]

// 		// Add either the full duration or the gap (if next attack comes sooner)
// 		if gap < duration {
// 			totalTime += gap
// 		} else {
// 			totalTime += duration
// 		}
// 	}

// 	// Don't forget the last attack - it always lasts full duration
// 	totalTime += duration

// 	return totalTime
// }

// Alternative implementation with slightly different logic
func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}

	totalTime := 0
	poisonEnd := 0 // Track when poison effect ends

	for _, attackTime := range timeSeries {
		if attackTime >= poisonEnd {
			// No overlap - add full duration
			totalTime += duration
		} else {
			// Overlap - only add the additional time
			totalTime += attackTime + duration - poisonEnd
		}
		// Update poison end time
		poisonEnd = attackTime + duration
	}

	return totalTime
}
