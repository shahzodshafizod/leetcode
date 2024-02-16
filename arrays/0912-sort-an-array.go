package arrays

// https://leetcode.com/problems/sort-an-array/

// Heap Sort
func sortArray(nums []int) []int {
	len := len(nums)
	var (
		getLeft   = func(parent int) int { return 2*parent + 1 }
		getParent = func(child int) int { return (child - 1) / 2 }
		compare   = func(i int, j int) bool { return nums[i] < nums[j] }
		swap      = func(i int, j int) { nums[i], nums[j] = nums[j], nums[i] }
	)
	// 1. create a max heap
	for i := 1; i < len; i++ {
		// siftUp(i)
		child := i
		parent := getParent(child)
		for parent >= 0 && compare(parent, child) {
			swap(parent, child)
			child = parent
			parent = getParent(child)
		}
	}
	// 2. sort ascending
	// move max to the end, and continue with the array[0:len-1]
	for len > 0 {
		len--
		swap(0, len)
		// siftDown(0, len)
		parent := 0
		child := getLeft(parent)
		for child < len {
			if child+1 < len && compare(child, child+1) { // right
				child++
			}
			if !compare(parent, child) {
				break
			}
			swap(parent, child)
			parent = child
			child = getLeft(parent)
		}
	}
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
