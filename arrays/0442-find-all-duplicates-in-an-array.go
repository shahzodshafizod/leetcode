package arrays

/*
ON and OFF:
1. In the beginning all switches (elements) are ON (positive)
2. Traverse elements of array and turn elements where the current pointer is pointing OFF (negative)
3. If an element where the current pointer is pointing is already OFF (negative), it is a duplicate.
*/

// https://leetcode.com/problems/find-all-duplicates-in-an-array/

func findDuplicates(nums []int) []int {
	var duplicates = make([]int, 0)
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
