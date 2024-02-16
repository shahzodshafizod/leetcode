package arrays

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/sort-an-array/

// Heap Sort
func sortArray(nums []int) []int {
	// 1. create a max heap
	var maxHeap = design.NewPriorityQueue(nums, func(x, y int) bool { return x < y })
	maxHeap.Init()
	// 2. sort ascending
	// move max to the end, and continue with the array[0:len-1]
	maxHeap.Sort()
	nums = maxHeap.Array()
	return nums
}

// // Quick Sort
// func sortArray(nums []int) []int {
// 	sortArrayQuick(nums, 0, len(nums)-1)
// 	return nums
// }

// func sortArrayQuick(nums []int, left int, right int) {
// 	if left >= right {
// 		return
// 	}
// 	var mid = sortArrayPartition(nums, left, right)
// 	sortArrayQuick(nums, left, mid-1)
// 	sortArrayQuick(nums, mid+1, right)
// }

// func sortArrayPartition(nums []int, left int, right int) int {
// 	var pivot = nums[right]
// 	var mid = left
// 	for left < right {
// 		if nums[left] < pivot {
// 			nums[left], nums[mid] = nums[mid], nums[left]
// 			mid++
// 		}
// 		left++
// 	}
// 	nums[mid], nums[right] = nums[right], nums[mid]
// 	return mid
// }
