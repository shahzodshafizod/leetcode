package twopointers

/*
DNF algorithm 🇳🇱
The Dutch National Flag Algorithm, also known as the DNF algorithm or
the Three-Way Partitioning Algorithm, is a simple and efficient approach
to sorting an array containing three distinct elements.
*/

// https://leetcode.com/problems/sort-colors/

// Approach: Three Pointers (DNF Algorithm)
// Time: O(n)
// Space: O(1)
func sortColors(nums []int) {
	for idx, left, right := 0, 0, len(nums)-1; idx <= right; idx++ {
		switch nums[idx] {
		case 0:
			nums[idx], nums[left] = nums[left], nums[idx]
			left++
		case 2:
			nums[idx], nums[right] = nums[right], nums[idx]
			right--
			idx--
		default:
		}
	}
}

// // Approach: Count Sort
// // Time: O(n)
// // Space: O(3)
// func sortColors(nums []int) {
// 	var buckets [3]int
// 	for _, num := range nums {
// 		buckets[num]++
// 	}
// 	var idx = 0
// 	for num, count := range buckets {
// 		for count > 0 {
// 			count--
// 			nums[idx] = num
// 			idx++
// 		}
// 	}
// }
