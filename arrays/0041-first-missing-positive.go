package arrays

// https://leetcode.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	var len = len(nums)
	// 1. delete unneccessary elements (set to zero)
	// and sort elements: O(2N) = O(N)
	for idx := range nums {
		for nums[idx] >= 1 &&
			nums[idx] <= len &&
			nums[idx] != idx+1 &&
			nums[idx] != nums[nums[idx]-1] {

			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		}
		if nums[idx] != idx+1 {
			nums[idx] = 0
		}
	}
	// 2. find first missing positive element: O(N)
	var missing int
	for idx := range nums {
		if nums[idx] == 0 {
			missing = idx + 1
			break
		}
	}
	// if all elements in array exist, then the first
	// missing is the next element after the last one.
	if missing == 0 {
		missing = len + 1
	}
	return missing
}
