package hashes

/*
Problem:
Given an array of integers, return the indices of the two numbers that add up to a given target.

Step 1: Verify the constraints
	- Are all the numbers positive or can there be negatives?
		: All numbers are positive.
	- Are there duplicate numbers in the array?
		: No, there areno duplicates.
	- Will there always be a solution available?
		: No, there may not always be a solution.
	- What do we return if there's no solution?
		: Just return null.
	- Can there be multiple pairs that add up to the target?
		: No, only 1 pair of numbers will add up to the target.
Step 2: Write out some test cases
	- [1, 3, 7, 9, 2], 11 -> [3, 4]
	- [1, 3, 7, 9, 2], 25 -> null
	- [], 1 -> null
	- [5], 5 -> null
	- [1, 6], 7 -> [0, 1]
Step 3: Figure out a solution without code
Step 4: Write out a solution in code
Step 5: Double check for errors
Step 6: Test the code with the test cases
Step 7: Space & Time Complexity
Step 8: Can we optimize the solution?
*/

// https://leetcode.com/problems/two-sum/

// Approach: Hash Table
// Time: O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		if j, exists := hash[nums[i]]; exists {
			return []int{i, j}
		}
		hash[target-nums[i]] = i
	}
	return nil
}

// // Approach: Two Pointers
// // Time: O(n log n)
// // Space: O(1)
// func twoSum(nums []int, target int) []int {
// 	var n = len(nums)
// 	var indices = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		indices[idx] = idx
// 	}
// 	sort.Slice(indices, func(i int, j int) bool {
// 		return nums[indices[i]] < nums[indices[j]]
// 	})
// 	var lo, hi = 0, n - 1
// 	for lo < hi {
// 		if nums[indices[lo]]+nums[indices[hi]] > target {
// 			hi--
// 		} else if nums[indices[lo]]+nums[indices[hi]] < target {
// 			lo++
// 		} else {
// 			return []int{indices[lo], indices[hi]}
// 		}
// 	}
// 	return nil
// }
