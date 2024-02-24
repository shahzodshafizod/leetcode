package arrays

import (
	"container/heap"

	"github.com/shahzodshafizod/alkhwarizmi/design"
)

/*
Data Structures: 2-D Arrays/ Matrixes, Binary Trees, Heaps, Graphs
Algorithmic approaches: Sorting, Greedy Method, Divide and Conquer, Dynamic Programming, Backtracking

Tail Recursion:
	- Normal recursion space: O(N)
		: func factorial(x int) int {
			if x <= 1 {
				return 1
			}
			return x * factorial(x - 1)
		}
	- Tail recursion space: O(1)
		: func factorial(x int, total int) int {
			if x <= 1 {
				return total
			}
			return factorial(x - 1, total * x)
		}

Problem:
Given an unsorted array, return the kth largest element.
It is the kth largest element in sorted order, not the kth distinct element.

Step 1: Verify the constraints
	- Can we get an array where k is larger than the array length?
		: No, assume an answer is always available.
Step 2: Write out some test cases

Algorithmic Paradigm: Divide & Conquer
	1. Multi-branched recursion.
	2. Breaks a problem into multiple smaller but same sub-problems.
	3. Combines the solutions of sub-problems into the solution for the original problem.
*/

// https://leetcode.com/problems/kth-largest-element-in-an-array/

// Approach#1: Heap (Priority Queue)
func findKthLargest(nums []int, k int) int {
	numsHeap := design.NewHeap(nums, func(x, y int) bool { return x > y })
	heap.Init(numsHeap)
	var kth int
	for k > 0 {
		k--
		kth = heap.Pop(numsHeap).(int)
	}
	return kth
}

// // Approach#2: Hoare's Quickselect Algorithm
// func findKthLargest(nums []int, k int) int {
// 	return quickSelect(nums, 0, len(nums)-1, k-1)
// }

// func quickSelect(nums []int, left, right int, k int) int {
// 	randIdx := left + (right-left)/2
// 	nums[randIdx], nums[right] = nums[right], nums[randIdx]
// 	var partition = left
// 	for i := left; i < right; i++ {
// 		if nums[i] > nums[right] {
// 			nums[i], nums[partition] = nums[partition], nums[i]
// 			partition++
// 		}
// 	}
// 	nums[right], nums[partition] = nums[partition], nums[right]

// 	if partition > k {
// 		return quickSelect(nums, left, partition-1, k)
// 	} else if partition < k {
// 		return quickSelect(nums, partition+1, right, k)
// 	}
// 	return nums[k]
// }

// // Approach#3 sort library
// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }

// // Approach#4 Quick Sort
// func findKthLargest(nums []int, k int) int {
// 	kthLargestQuickSort(nums, 0, len(nums)-1)
// 	return nums[k-1]
// }

// func kthLargestQuickSort(nums []int, left int, right int) {
// 	if left >= right {
// 		return
// 	}
// 	randIdx := left + rand.Intn(right-left+1)
// 	nums[right], nums[randIdx] = nums[randIdx], nums[right]
// 	var partition = left
// 	for i := left; i < right; i++ {
// 		if nums[i] > nums[right] {
// 			nums[i], nums[partition] = nums[partition], nums[i]
// 			partition++
// 		}
// 	}
// 	nums[partition], nums[right] = nums[right], nums[partition]
// 	kthLargestQuickSort(nums, left, partition-1)
// 	kthLargestQuickSort(nums, partition+1, right)
// }

// // Approach#5: Priority Queue (array as tree)
// func findKthLargest(nums []int, k int) int {
// 	var pq = design.NewPQ(func(x, y int) bool { return x < y })
// 	for _, num := range nums {
// 		pq.Push(num)
// 	}
// 	for i := 1; i < k; i++ {
// 		pq.Pop()
// 	}
// 	return pq.Peek()
// }
