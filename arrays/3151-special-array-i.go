package arrays

// https://leetcode.com/problems/special-array-i/

func isArraySpecial(nums []int) bool {
	for idx := len(nums) - 1; idx > 0; idx-- {
		if (nums[idx-1]+nums[idx])&1 == 0 {
			return false
		}
	}

	return true
}
