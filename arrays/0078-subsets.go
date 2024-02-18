package arrays

// https://leetcode.com/problems/subsets/

// [[]]
// [[],[1]]
// [[],[1],[2],[1,2]]
// [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// iteratively
func subsets(nums []int) [][]int {
	var subsets = [][]int{{}}
	for _, num := range nums {
		// 1. decision NOT to include num
		// 2. decision to include num
		for i, len := 0, len(subsets); i < len; i++ {
			subsets = append(subsets, append([]int{}, append(subsets[i], num)...))
		}
	}
	return subsets
}

// func subsets(nums []int) [][]int {
// 	var subsets = make([][]int, 0)
// 	var subset = make([]int, 0)
// 	var dfs func(int)
// 	var length = len(nums)
// 	dfs = func(idx int) {
// 		if idx >= length {
// 			subsets = append(subsets, append([]int{}, subset...))
// 			return
// 		}
// 		// decision to include nums[idx]
// 		subset = append(subset, nums[idx])
// 		dfs(idx + 1)
// 		// decision NOT to include nums[idx]
// 		subset = subset[:len(subset)-1]
// 		dfs(idx + 1)
// 	}
// 	dfs(0)
// 	return subsets
// }
