package arrays

/*
DNF algorithm 🇳🇱
The Dutch National Flag Algorithm, also known as the DNF algorithm or
the Three-Way Partitioning Algorithm, is a simple and efficient approach
to sorting an array containing three distinct elements.
*/

// https://leetcode.com/problems/sort-colors/

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
		}
	}
}

// func sortColors(nums []int) {
// 	var buckets = make([]int, 3)
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
