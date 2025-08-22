from typing import List
import unittest

# https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/

# python3 -m unittest matrices/3195-find-the-minimum-area-to-cover-all-ones-i.py


class Solution(unittest.TestCase):
    def minimumArea(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        left, right = n, 0
        top, down = m, 0
        for row in range(m):
            for col in range(n):
                if grid[row][col]:
                    left = min(left, col)
                    right = max(right, col)
                    top = min(top, row)
                    down = max(down, row)
        if left == n:
            return 0
        height = down - top + 1
        width = right - left + 1
        return height * width

    def test(self):
        for grid, expected in [
            ([[0, 1, 0], [1, 0, 1]], 6),
            ([[1, 0], [0, 0]], 1),
        ]:
            output = self.minimumArea(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
