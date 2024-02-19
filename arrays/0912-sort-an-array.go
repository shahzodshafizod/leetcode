package arrays

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/sort-an-array/

// Heap Sort
func sortArray(nums []int) []int {
	var maxHeap = design.NewPriorityQueue(
		nums,
		func(x, y int) bool { return x < y },
	)
	maxHeap.Heapify()
	maxHeap.Sort()
	return maxHeap.Array()
}
