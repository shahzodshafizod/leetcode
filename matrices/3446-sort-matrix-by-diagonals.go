package matrices

import "sort"

// https://leetcode.com/problems/sort-matrix-by-diagonals/

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	for d := n - 1; d >= 0; d-- {
		var nums []int
		for col := d; col >= 0; col-- {
			nums = append(nums, grid[n-1-(d-col)][col])
		}

		sort.Ints(nums)

		for col := d; col >= 0; col-- {
			grid[n-1-d+col][col] = nums[d-col]
		}
	}

	for d := 1; d < n; d++ {
		var nums []int
		for col := d; col < n; col++ {
			nums = append(nums, grid[col-d][col])
		}

		sort.Ints(nums)

		for col := d; col < n; col++ {
			grid[col-d][col] = nums[col-d]
		}
	}

	return grid
}
