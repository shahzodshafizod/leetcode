package arrays

// https://leetcode.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 1. delete unnecessary elements (set to zero)
	// and sort elements: O(2N) = O(N)
	for idx := range nums {
		for nums[idx] >= 1 &&
			nums[idx] <= n &&
			nums[idx] != idx+1 &&
			nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		}

		if nums[idx] != idx+1 {
			nums[idx] = 0
		}
	}
	// 2. find first missing positive element: O(N)
	for idx := range nums {
		if nums[idx] == 0 {
			return idx + 1
		}
	}
	// if all elements in array exist, then the first
	// missing is the next element after the last one.
	return n + 1
}
