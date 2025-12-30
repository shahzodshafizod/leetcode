package arrays

// https://leetcode.com/problems/count-of-range-sum/

// Approach: Divide and Conquer with Merge Sort
// Key insight: For range sum S(i,j) = prefix[j+1] - prefix[i]
// We need: lower <= prefix[j+1] - prefix[i] <= upper
// Rearranging: prefix[i] + lower <= prefix[j+1] <= prefix[i] + upper
//
// During merge sort of prefix sums:
//   - For each prefix[i] in left half, count how many prefix[j] in right half
//     satisfy: prefix[i] + lower <= prefix[j] <= prefix[i] + upper
//   - Use two pointers since both halves are sorted
//
// Time: O(n log n) - merge sort with linear merge
// Space: O(n) - for prefix array and merge temporary array
func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	// Build prefix sum array
	prefix := make([]int, n+1)
	for i := range n {
		prefix[i+1] = prefix[i] + nums[i]
	}

	var mergeSort func(start, end int) int

	mergeSort = func(start, end int) int {
		if end-start <= 1 {
			return 0
		}

		mid := (start + end) / 2
		count := mergeSort(start, mid) + mergeSort(mid, end)

		// Count valid pairs where i in [start, mid) and j in [mid, end)
		// We need: lower <= prefix[j] - prefix[i] <= upper
		// Which means: prefix[i] + lower <= prefix[j] <= prefix[i] + upper
		jLow, jHigh := mid, mid
		for i := start; i < mid; i++ {
			// Find the range [jLow, jHigh) where prefix[j] satisfies the condition
			for jLow < end && prefix[jLow]-prefix[i] < lower {
				jLow++
			}

			for jHigh < end && prefix[jHigh]-prefix[i] <= upper {
				jHigh++
			}

			count += jHigh - jLow
		}

		// Merge the two sorted halves
		sorted := make([]int, 0, end-start)

		left, right := start, mid
		for left < mid && right < end {
			if prefix[left] <= prefix[right] {
				sorted = append(sorted, prefix[left])
				left++
			} else {
				sorted = append(sorted, prefix[right])
				right++
			}
		}

		sorted = append(sorted, prefix[left:mid]...)
		sorted = append(sorted, prefix[right:end]...)

		// Copy back to prefix array
		copy(prefix[start:end], sorted)

		return count
	}

	return mergeSort(0, n+1)
}

// // Approach 1: Brute Force
// // Generate all possible range sums and count those in [lower, upper]
// // For each starting position i, compute all range sums starting from i
// // Time: O(n^2) - nested loops to generate all ranges
// // Space: O(n) - prefix sum array
// func countRangeSum(nums []int, lower int, upper int) int {
// 	n := len(nums)
// 	// Build prefix sum array for O(1) range sum queries
// 	prefix := make([]int, n+1)
// 	for i := range n {
// 		prefix[i+1] = prefix[i] + nums[i]
// 	}

// 	count := 0
// 	// Try all possible ranges [i, j]
// 	for i := range n {
// 		for j := i; j < n; j++ {
// 			// Range sum S(i, j) = prefix[j+1] - prefix[i]
// 			rangeSum := prefix[j+1] - prefix[i]
// 			if lower <= rangeSum && rangeSum <= upper {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
