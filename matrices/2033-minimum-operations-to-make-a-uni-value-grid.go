package matrices

import (
	"sort"
)

// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/

func minOperations2033(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	total := 0
	nums := make([]int, 0, m*n)
	mod := grid[0][0] % x

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col]%x != mod {
				return -1
			}

			total += grid[row][col]
			nums = append(nums, grid[row][col])
		}
	}

	sort.Ints(nums)
	mid := nums[m*n/2]
	res := 0

	for _, num := range nums {
		if num > mid {
			res += (num - mid) / x
		} else {
			res += (mid - num) / x
		}
	}

	return res
}
