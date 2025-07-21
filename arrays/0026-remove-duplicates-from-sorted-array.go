package arrays

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	index := 0
	for _, num := range nums {
		if num != nums[index] {
			index++
			nums[index] = num
		}
	}
	return index + 1
}

// func removeDuplicates(nums []int) int {
// 	return len(slices.Compact[[]int](nums))
// }

// func removeDuplicates(nums []int) int {
// 	var len = len(nums)
// 	var duplicates = 0
// 	for idx := 1; idx < len; idx++ {
// 		if nums[idx] == nums[idx-1] {
// 			duplicates++
// 		} else if duplicates > 0 {
// 			nums[idx-duplicates] = nums[idx]
// 		}
// 	}
// 	return len - duplicates
// }
