package bits

// https://leetcode.com/problems/find-subarray-with-bitwise-or-closest-to-k/

// Approach #2: Sliding OR Window with Early Break
// Time: O(n * log MAX_VAL) ≈ O(n * 30) = O(n)
// Space: O(log MAX_VAL) ≈ O(30) = O(1)
func minimumDifference(nums []int, k int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	minDiff := abs(nums[0] - k)

	ors := []int{}

	for _, num := range nums {
		ors = append(ors, 0)
		for i := len(ors) - 1; i >= 0; i-- {
			if ors[i] == ors[i]|num {
				// ORing num with any previous OR values
				// (which are subsets of ors[i]) won't change them either
				break
			}

			ors[i] |= num
			minDiff = min(minDiff, abs(k-ors[i]))
		}
	}

	return minDiff
}

// // Approach #1: Brute-Force
// // Time: O(nn)
// // Space: O(1)
// func minimumDifference(nums []int, k int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	n := len(nums)
// 	minDiff := abs(nums[0] - k)

// 	for i := range n {
// 		currentOr := 0
// 		for j := i; j < n; j++ {
// 			currentOr |= nums[j]
// 			minDiff = min(minDiff, abs(currentOr-k))
// 		}
// 	}

// 	return minDiff
// }
