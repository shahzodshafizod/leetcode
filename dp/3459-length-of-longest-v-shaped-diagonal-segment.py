from functools import cache
from typing import List
import unittest

# https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/

# python3 -m unittest dp/3459-length-of-longest-v-shaped-diagonal-segment.py


class Solution(unittest.TestCase):
    # Approach: Memoization Search
    # Time: O(n x m)
    # Space: O(n x m)
    def lenOfVDiagonal(self, grid: List[List[int]]) -> int:
        DIRS = [(-1, -1), (-1, 1), (1, 1), (1, -1)]
        n, m = len(grid), len(grid[0])

        @cache
        def dfs(row: int, col: int, turn: int, dirn: int, target: int) -> int:
            if min(row, col) < 0 or row == n or col == m or grid[row][col] != target:
                return 0
            nrow, ncol = row + DIRS[dirn][0], col + DIRS[dirn][1]
            length = 1 + dfs(nrow, ncol, turn, dirn, 2 - target)
            if turn:
                ndir = (dirn + 1) % 4
                nrow, ncol = row + DIRS[ndir][0], col + DIRS[ndir][1]
                length = max(length, 1 + dfs(nrow, ncol, 0, ndir, 2 - target))
            return length

        length = 0
        for row in range(n):
            for col in range(m):
                if grid[row][col] == 1:
                    for dirn in range(4):
                        length = max(length, 1 + dfs(row + DIRS[dirn][0], col + DIRS[dirn][1], 1, dirn, 2))
        return length

    def test(self):
        for grid, expected in [
            ([[2, 2, 1, 2, 2], [2, 0, 2, 2, 0], [2, 0, 1, 1, 0], [1, 0, 2, 2, 2], [2, 0, 0, 2, 2]], 5),
            ([[2, 2, 2, 2, 2], [2, 0, 2, 2, 0], [2, 0, 1, 1, 0], [1, 0, 2, 2, 2], [2, 0, 0, 2, 2]], 4),
            ([[1, 2, 2, 2, 2], [2, 2, 2, 2, 0], [2, 0, 0, 0, 0], [0, 0, 2, 2, 2], [2, 0, 0, 2, 0]], 5),
            ([[1]], 1),
        ]:
            output = self.lenOfVDiagonal(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
