package arrays

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/sort-an-array/

// Approach: Heap Sort
// Time: O(N Log N)
// Space: O(1)
func sortArray(nums []int) []int {
	var maxHeap = design.NewPQSort(
		len(nums),
		func(i, j int) bool { return nums[i] < nums[j] },
		func(i, j int) { nums[i], nums[j] = nums[j], nums[i] },
	)
	maxHeap.Sort()
	return nums
}
