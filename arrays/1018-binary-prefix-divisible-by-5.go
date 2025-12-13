package arrays

// https://leetcode.com/problems/binary-prefix-divisible-by-5/

// Approach: Track running value modulo 5
// Time: O(n) - single pass
// Space: O(n) - result array
func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))

	value := 0
	for i, bit := range nums {
		value = (value*2 + bit) % 5
		result[i] = value == 0
	}

	return result
}
