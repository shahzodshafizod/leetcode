package bits

// https://leetcode.com/problems/bitwise-xor-of-all-pairings/

// Approach: Bit Manipulation
// Time: O(n1 + n2)
// Space: O(1)
func xorAllNums(nums1 []int, nums2 []int) int {
	var num1 = 0
	for _, num := range nums1 {
		num1 ^= num
	}
	var num2 = 0
	for _, num := range nums2 {
		num2 ^= num
	}
	return (len(nums2) % 2 * num1) ^ (len(nums1) % 2 * num2)
}
