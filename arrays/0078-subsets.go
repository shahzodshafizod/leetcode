package arrays

// https://leetcode.com/problems/subsets/

// [[]]
// [[],[1]]
// [[],[1],[2],[1,2]]
// [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	var dfs func(idx int) [][]int

	dfs = func(idx int) [][]int {
		if idx < 0 {
			return [][]int{{}}
		}

		subsets := dfs(idx - 1)
		for i, n := 0, len(subsets); i < n; i++ {
			cpy := append([]int{}, subsets[i]...)
			cpy = append(cpy, nums[idx])
			subsets = append(subsets, cpy)
		}

		return subsets
	}
	subsets := dfs(len(nums) - 1)

	return subsets
}

// // iteratively
// func subsets(nums []int) [][]int {
// 	var subsets = [][]int{{}}
// 	for _, num := range nums {
// 		// 1. decision NOT to include num
// 		// 2. decision to include num
// 		for i, len := 0, len(subsets); i < len; i++ {
// 			subsets = append(subsets, append([]int{}, append(subsets[i], num)...))
// 		}
// 	}
// 	return subsets
// }
