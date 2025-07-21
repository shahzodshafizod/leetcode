package arrays

/*
ON and OFF:
1. In the beginning all switches (elements) are ON (positive)
2. Traverse elements of array and turn elements where the current pointer is pointing OFF (negative)
3. If an element where the current pointer is pointing is already OFF (negative), it is a duplicate.
*/

// https://leetcode.com/problems/find-all-duplicates-in-an-array/

func findDuplicates(nums []int) []int {
	duplicates := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		num--
		if nums[num] < 0 {
			duplicates = append(duplicates, num+1)
		} else {
			nums[num] = -nums[num]
		}
	}
	return duplicates
}

// // Bit Manipulation
// func findDuplicates(nums []int) []int {
// 	var duplicates = make([]int, 0)
// 	var number, count int
// 	for _, pointer := range nums {
// 		pointer--
// 		pointer &= 0xFFFF
// 		number = nums[pointer] & 0xFFFF
// 		count = nums[pointer] >> 16
// 		count++
// 		if count == 2 {
// 			duplicates = append(duplicates, pointer+1)
// 		}
// 		nums[pointer] = (count << 16) + number
// 	}
// 	return duplicates
// }
