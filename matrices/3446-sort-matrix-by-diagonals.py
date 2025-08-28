from typing import List
import unittest

# https://leetcode.com/problems/sort-matrix-by-diagonals/

# python3 -m unittest matrices/3446-sort-matrix-by-diagonals.py


class Solution(unittest.TestCase):
    def sortMatrix(self, grid: List[List[int]]) -> List[List[int]]:
        n = len(grid)
        for i in range(n):
            nums = sorted([grid[i + j][j] for j in range(n - i)], reverse=True)
            for j in range(n - i):
                grid[i + j][j] = nums[j]
        for j in range(1, n):
            nums = sorted([grid[i][i + j] for i in range(n - j)])
            for i in range(n - j):
                grid[i][i + j] = nums[i]
        return grid

    def test(self):
        for grid, expected in [
            ([[1, 7, 3], [9, 8, 2], [4, 5, 6]], [[8, 2, 3], [9, 6, 7], [4, 5, 1]]),
            ([[0, 1], [1, 2]], [[2, 1], [1, 0]]),
            ([[1]], [[1]]),
        ]:
            output = self.sortMatrix(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
