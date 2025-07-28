package arrays

// https://leetcode.com/problems/product-of-array-except-self/

// Follow up: Can you solve the problem in O(1) extra space complexity?
// The output array does not count as extra space for space complexity analysis.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// time: O(2n) = O(n)
// space: O(1)
func productExceptSelf(nums []int) []int {
	len := len(nums)
	answer := make([]int, len)
	product := 1

	for idx := 0; idx < len; idx++ { // from left to right
		answer[idx] = product
		product *= nums[idx]
	}

	product = 1
	for idx := len - 1; idx >= 0; idx-- { // from right to left
		answer[idx] *= product
		product *= nums[idx]
	}

	return answer
}

// func productExceptSelf(nums []int) []int {
// 	var len = len(nums)
// 	var answer = make([]int, len)
// 	var left = 1
// 	for i := 0; i < len; i++ {
// 		answer[i] = left
// 		for j := i + 1; j < len; j++ {
// 			answer[i] *= nums[j]
// 		}
// 		left *= nums[i]
// 	}
// 	return answer
// }
