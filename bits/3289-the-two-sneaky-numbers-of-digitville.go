package bits

// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/

// Approach #3: Bit Manipulations
// Time: O(n)
// Space: O(1)
func getSneakyNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	n := len(nums) - 2
	for idx := range n {
		xor ^= idx
	}

	randomBit := xor & -xor
	// -xor means: ~xor (invert bits) and +1
	var num1, num2 int

	for _, num := range nums {
		if num&randomBit == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	for idx := range n {
		if idx&randomBit == 0 {
			num1 ^= idx
		} else {
			num2 ^= idx
		}
	}

	return []int{num1, num2}
}

// // Approach #2: Sorting
// // Time: O(n log n)
// // Space: O(1)
// func getSneakyNumbers(nums []int) []int {
// 	sort.Ints(nums)

// 	sneakies := make([]int, 0, 2)

// 	for i, n := 1, len(nums); i < n; i++ {
// 		if nums[i-1] == nums[i] {
// 			sneakies = append(sneakies, nums[i])
// 		}
// 	}

// 	return sneakies
// }

// // Approach #1: Counting
// // Time: O(n)
// // Space: O(n)
// func getSneakyNumbers(nums []int) []int {
// 	n := len(nums)
// 	count := make([]int, n-2)
// 	sneakies := make([]int, 0, 2)

// 	for _, num := range nums {
// 		count[num]++
// 		if count[num] > 1 {
// 			sneakies = append(sneakies, num)
// 		}
// 	}

// 	return sneakies
// }
