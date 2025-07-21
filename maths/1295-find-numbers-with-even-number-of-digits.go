package maths

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

// Approach: Extract Digits
// Time: O(NxM), N=len(nums), M=max(len(nums[i]))
// Space: O(1)
func findNumbers(nums []int) int {
	count := 0
	var digits int
	for _, num := range nums {
		digits = 0
		for num > 0 {
			digits++
			num /= 10
		}
		if digits&1 == 0 {
			count++
		}
	}
	return count
}

// // Approach: Convert to String
// // Time: O(NxM), N=len(nums), M=max(len(nums[i]))
// // Space: O(M)
// func findNumbers(nums []int) int {
// 	var count = 0
// 	for _, num := range nums {
// 		if len(strconv.Itoa(num))&1 == 0 {
// 			count++
// 		}
// 	}
// 	return count
// }
