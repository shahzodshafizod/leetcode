package arrays

// https://leetcode.com/problems/build-array-from-permutation/

// Space: O(1)
func buildArray(nums []int) []int {
	const shift = 10

	const mask = (1 << shift) - 1
	for i := range nums {
		nums[i] |= (nums[nums[i]] & mask) << shift
	}

	for i := range nums {
		nums[i] >>= shift
	}

	return nums
}

// // Space: O(n)
// func buildArray(nums []int) []int {
// 	var ans = make([]int, len(nums))
// 	for idx := range nums {
// 		ans[idx] = nums[nums[idx]]
// 	}
// 	return ans
// }
