package binarysearch

// https://leetcode.com/problems/maximum-running-time-of-n-computers/

// Approach 1: Brute Force (Would be infeasible for large inputs)
// We could try every possible time from 1 to sum(batteries), and for each time,
// check if we can run all n computers for that duration.
// Time: O((sum of batteries) * m) where m is number of batteries - TLE
// Space: O(1)

// // Approach 2: Binary Search + Greedy
// // The key insight is that if we can run all n computers for time t,
// // then we can also run them for any time < t (monotonic property).
// // This suggests binary search on the answer.
// //
// // For a given time t, we can run all n computers if:
// // sum(min(battery[i], t)) >= n * t
// //
// // The reasoning: Each battery can contribute at most min(battery[i], t) minutes.
// // If the total contribution is >= n * t, we can distribute the power to run
// // all n computers for t minutes (we can swap batteries).
// //
// // Time: O(m * log(sum/n)) where m is number of batteries
// // Space: O(1)

// func maxRunTime(n int, batteries []int) int64 {
// 	// Binary search on the answer
// 	// Lower bound: minimum is 1 minute
// 	// Upper bound: total sum divided by n computers
// 	var totalSum int64 = 0
// 	for _, battery := range batteries {
// 		totalSum += int64(battery)
// 	}

// 	left, right := int64(1), totalSum/int64(n)

// 	// Binary search for maximum time
// 	for left < right {
// 		mid := right - (right-left)/2 // Use right-biased mid to find maximum

// 		// Check if we can run all n computers for 'mid' minutes
// 		var availablePower int64 = 0
// 		for _, battery := range batteries {
// 			// Each battery contributes min(battery, mid) minutes
// 			batteryPower := int64(battery)
// 			if batteryPower > mid {
// 				batteryPower = mid
// 			}

// 			availablePower += batteryPower
// 		}

// 		// If total available power >= n * mid, we can run all computers for mid minutes
// 		if availablePower >= int64(n)*mid {
// 			left = mid // Try for longer time
// 		} else {
// 			right = mid - 1 // Reduce time
// 		}
// 	}

// 	return left
// }

// Alternative implementation using slices.Max for clearer upper bound
func maxRunTime(n int, batteries []int) int64 {
	var totalSum int64
	for _, battery := range batteries {
		totalSum += int64(battery)
	}

	// Binary search on possible running time
	left, right := int64(1), totalSum/int64(n)

	var result int64

	for left <= right {
		mid := left + (right-left)/2

		// Check if we can run n computers for mid minutes
		var powerSum int64

		for _, battery := range batteries {
			if int64(battery) < mid {
				powerSum += int64(battery)
			} else {
				powerSum += mid
			}
		}

		if powerSum >= int64(n)*mid {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
