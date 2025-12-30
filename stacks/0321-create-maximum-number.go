package stacks

// https://leetcode.com/problems/create-maximum-number/

// Approach: Greedy + Monotonic Stack + Merge
// This problem can be broken down into 3 sub-problems:
// 1. Find max subsequence of length k from a single array (monotonic stack)
// 2. Merge two arrays to get lexicographically largest result
// 3. Try all possible splits: i digits from nums1, k-i from nums2
//
// Time: O(k * (m+n)^2) where m=len(nums1), n=len(nums2)
//   - O(k) iterations for different splits
//   - O(m+n) for maxSubsequence for each array
//   - O((m+n)^2) for merge comparison in worst case
//
// Space: O(k) for storing subsequences and result
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// Get the maximum subsequence of given length from nums
	// while maintaining relative order using monotonic stack
	maxSubsequence := func(nums []int, length int) []int {
		if length == 0 {
			return []int{}
		}

		stack := []int{}
		toDrop := len(nums) - length // How many elements we can/must drop

		for _, num := range nums {
			// Pop smaller elements from stack if we still have drops left
			for len(stack) > 0 && stack[len(stack)-1] < num && toDrop > 0 {
				stack = stack[:len(stack)-1]
				toDrop--
			}

			stack = append(stack, num)
		}

		// Return only the first 'length' elements
		// (we might have more if we didn't drop enough)
		if len(stack) > length {
			return stack[:length]
		}

		return stack
	}

	// Helper to compare two slices lexicographically
	isGreater := func(nums1 []int, i int, nums2 []int, j int) bool {
		for i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				return true
			}

			if nums1[i] < nums2[j] {
				return false
			}

			i++
			j++
		}

		return i < len(nums1)
	}

	// Merge two arrays to get the lexicographically largest result
	merge := func(nums1 []int, nums2 []int) []int {
		result := make([]int, 0, len(nums1)+len(nums2))
		i, j := 0, 0

		for i < len(nums1) && j < len(nums2) {
			// Compare remaining subarrays lexicographically
			// Choose the one that gives larger result
			if isGreater(nums1, i, nums2, j) {
				result = append(result, nums1[i])
				i++
			} else {
				result = append(result, nums2[j])
				j++
			}
		}

		// Append remaining elements
		result = append(result, nums1[i:]...)
		result = append(result, nums2[j:]...)

		return result
	}

	// Helper to compare two slices
	greater := func(a, b []int) bool {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] > b[i] {
				return true
			}

			if a[i] < b[i] {
				return false
			}
		}

		return len(a) > len(b)
	}

	m, n := len(nums1), len(nums2)
	maxResult := []int{}

	// Try all possible splits: i from nums1, k-i from nums2
	for i := 0; i <= k; i++ {
		j := k - i

		// Check if this split is valid
		if i > m || j > n {
			continue
		}

		// Get max subsequence from each array
		sub1 := maxSubsequence(nums1, i)
		sub2 := maxSubsequence(nums2, j)

		// Merge and compare with current max
		merged := merge(sub1, sub2)
		if greater(merged, maxResult) {
			maxResult = merged
		}
	}

	return maxResult
}
