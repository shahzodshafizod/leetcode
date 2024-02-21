package dp

/*
Dynamic Programming
5 Simple Steps For Solving Dynamic Programming Problems:
	- Visualize Examples
	- Find an appropriate subproblem
	- Find relationships among subproblems
	- Generalize the relationship
	- Implement by solving subproblems in order
*/

// https://leetcode.com/problems/longest-increasing-subsequence/

// O(n log n)
// CACHE - does not keep the real subsequence
// its length is the length of the longest increasing subsequence
// we'll keep many subsequences in the same array by replacing previous values
func lengthOfLIS(nums []int) int {
	var cache = make([]int, 0)
	var size = 0
	for _, num := range nums {
		if size == 0 || num > cache[size-1] {
			cache = append(cache, num)
			size++
			continue
		}
		// it's not a binary search, it's building blocks to the full solution
		idx, right := 0, size
		for idx < right {
			mid := (idx + right) / 2
			if cache[mid] < num {
				idx = mid + 1
			} else {
				right = mid
			}
		}
		cache[idx] = num
	}
	return size
}

// // Brute Force: O(n)
// func lengthOfLIS(nums []int) int {
// 	var lis = 0
// 	var cache = make([]int, 0, len(nums))
// 	for idx, num := range nums {
// 		var max = 0
// 		for j := 0; j < idx; j++ {
// 			if nums[j] < num && cache[j] > max {
// 				max = cache[j]
// 			}
// 		}
// 		max++
// 		if max > lis {
// 			lis = max
// 		}
// 		cache = append(cache, max)
// 	}
// 	return lis
// }
