package dp

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/

// Approach #2: Meet in the Middle + Binary Search
// Time: O(2^n * n) - generate all subsets for each half, then binary search
// Space: O(2^n) - storing all possible sums for each half
func minimumDifference(nums []int) int {
	n := len(nums) / 2

	total := 0
	for _, num := range nums {
		total += num
	}

	target := total / 2

	// Split array into two halves
	leftHalf := nums[:n]
	rightHalf := nums[n:]

	// For each half, generate all possible sums grouped by count of elements
	// leftSums[k] = list of all possible sums using exactly k elements from left half
	leftSums := make([][]int, n+1)
	rightSums := make([][]int, n+1)

	// Generate all subsets for left half using bitmask
	for mask := range 1 << n {
		sum := 0
		count := 0

		for i := range n {
			if mask&(1<<i) != 0 {
				sum += leftHalf[i]
				count++
			}
		}

		leftSums[count] = append(leftSums[count], sum)
	}

	// Generate all subsets for right half using bitmask
	for mask := range 1 << n {
		sum := 0
		count := 0

		for i := range n {
			if mask&(1<<i) != 0 {
				sum += rightHalf[i]
				count++
			}
		}

		rightSums[count] = append(rightSums[count], sum)
	}

	// Sort right_sums for binary search
	for count := 0; count <= n; count++ {
		sort.Ints(rightSums[count])
	}

	minDiff := math.MaxInt

	// Try all possible combinations
	// We take k elements from left and (n-k) elements from right
	for leftCount := 0; leftCount <= n; leftCount++ {
		rightCount := n - leftCount

		for _, leftSum := range leftSums[leftCount] {
			// We want: leftSum + rightSum ≈ target
			// So we need: rightSum ≈ target - leftSum
			needed := target - leftSum

			// Binary search for closest value in rightSums[rightCount]
			rightList := rightSums[rightCount]

			// Find position where 'needed' would be inserted
			idx := sort.SearchInts(rightList, needed)

			// Check idx and idx-1 for closest values
			candidates := []int{}
			if idx > 0 {
				candidates = append(candidates, idx-1)
			}

			if idx < len(rightList) {
				candidates = append(candidates, idx)
			}

			for _, i := range candidates {
				rightSum := rightList[i]
				sum1 := leftSum + rightSum
				sum2 := total - sum1

				diff := int(math.Abs(float64(sum1 - sum2)))
				if diff < minDiff {
					minDiff = diff
				}
			}
		}
	}

	return minDiff
}

// // Approach #1: Brute-Force (Try all 2^(2n) partitions)
// // Time: O(2^(2n)) - exponential, trying all possible partitions
// // Space: O(2n) - recursion stack
// func minimumDifference(nums []int) int {
// 	n := len(nums) / 2
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	minDiff := math.MaxInt

// 	var backtrack func(idx int, count int, currentSum int)
// 	backtrack = func(idx int, count int, currentSum int) {
// 		// We need exactly n elements in first partition
// 		if count > n || (len(nums)-idx) < (n-count) {
// 			return
// 		}

// 		if idx == len(nums) {
// 			if count == n {
// 				// Sum of second partition
// 				otherSum := total - currentSum
// 				diff := int(math.Abs(float64(currentSum - otherSum)))
// 				if diff < minDiff {
// 					minDiff = diff
// 				}
// 			}
// 			return
// 		}

// 		// Include current element in first partition
// 		backtrack(idx+1, count+1, currentSum+nums[idx])

// 		// Exclude current element (goes to second partition)
// 		backtrack(idx+1, count, currentSum)
// 	}

// 	backtrack(0, 0, 0)
// 	return minDiff
// }
