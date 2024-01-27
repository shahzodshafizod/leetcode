package arrays

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

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

// // Approach#1
// Hoare's Quickselect Algorithm
func quickSelect(nums []int, left, right int, k int) int {
	/*
		how to use:
		return quickSelect(nums, 0, len(nums)-1, k-1)
	*/

	randIdx := left + (right-left)/2
	nums[randIdx], nums[right] = nums[right], nums[randIdx]
	var partition = left
	for i := left; i < right; i++ {
		if nums[i] > nums[right] {
			nums[i], nums[partition] = nums[partition], nums[i]
			partition++
		}
	}
	nums[right], nums[partition] = nums[partition], nums[right]

	if partition > k {
		return quickSelect(nums, left, partition-1, k)
	} else if partition < k {
		return quickSelect(nums, partition+1, right, k)
	}
	return nums[k]
}

/*
how to use:
h := IntHeap(nums[:k])
heap.Init(&h)
for _, num := range nums {
	if num > h[0] {
		heap.Pop(&h)
		heap.Push(&h, num)
	}
}
return h[0]
*/

// // Approach#2
// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x any) {
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() any {
// 	value := (*h)[h.Len()-1]
// 	*h = (*h)[:h.Len()-1]
// 	return value
// }

// // Approach#3
// func quickSort(nums []int, left int, right int) {
// 	/*
// 		how to use:
// 		(1)
// 			quickSort(nums, 0, len(nums)-1)
// 			return nums[k-1]
// 		(2)
// 			sort.Ints(nums)
// 			return nums[len(nums)-k]
// 	*/
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
// 	quickSort(nums, left, partition-1)
// 	quickSort(nums, partition+1, right)
// }

// // Approach#4: Priority Queue (array as tree)
// func findKthLargest(nums []int, k int) int {
// 	var pq = trees.NewPriorityQueue()
// 	for _, num := range nums {
// 		pq.Push(num)
// 	}
// 	for i := 1; i < k; i++ {
// 		pq.Pop()
// 	}
// 	return pq.Peek()
// }
