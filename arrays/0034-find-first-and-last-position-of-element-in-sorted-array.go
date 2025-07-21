package arrays

/*
O(LogN) for Binary Search is O(Log2(N))
LogA(N) = B, A^B = N; Log2(16) = 4, 2^4 = 16

Problem:
Given an array of integers sorted in ascending order, return the starting
and ending index of a given target value in an array, i.e. [x, y].
Your solution should run in O(logN) time.

Step 1: Verify the constraints:
	- What do we return if the target is not found in the array?
		: Return -1, all values in the array are positive.
Step 2: Write out some test cases:
	- [1, 3, 3, 5, 5, 5, 8, 9], 5 => [3, 5]
	- [1, 2, 3, 4, 5, 6], 4 => [3, 3]
	- [1, 2, 3, 4, 5], 9 => [-1, -1]
	- [], 3 => [-1, -1]
	- [4, 4, 4, 4, 4, 4], 4 => [0, 5]
Step 3: Figure out a solution without a code

*/

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}
	left, right := 0, length-1
	position := binarySearch(nums, &left, &right, target)

	first := -1
	for pos := position; pos != -1; {
		first = pos
		pos--
		pos = binarySearch(nums, &left, &pos, target)
	}

	last := -1
	for pos := position; pos != -1; {
		last = pos
		pos++
		pos = binarySearch(nums, &pos, &right, target)
	}

	return []int{first, last}
}

func binarySearch(nums []int, left *int, right *int, target int) int {
	for *left <= *right {
		mid := (*left + *right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			*left = mid + 1
		} else {
			*right = mid - 1
		}
	}
	return -1
}
