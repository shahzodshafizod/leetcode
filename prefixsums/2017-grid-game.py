from typing import List
import unittest

# https://leetcode.com/problems/grid-game/

# python3 -m unittest prefixsums/2017-grid-game.py


class Solution(unittest.TestCase):
    # Approach: Prefix Sum
    # Time: O(n)
    # Space: O(1)
    def gridGame(self, grid: List[List[int]]) -> int:
        bottom, top = 0, sum(grid[0])
        robot2 = float("inf")
        for col in range(len(grid[0])):
            top -= grid[0][col]
            robot2 = min(robot2, max(top, bottom))
            bottom += grid[1][col]
        return robot2

    def test(self):
        for grid, expected in [
            ([[2, 5, 4], [1, 5, 1]], 4),
            ([[3, 3, 1], [8, 5, 2]], 4),
            ([[1, 3, 1, 15], [1, 3, 3, 1]], 7),
        ]:
            output = self.gridGame(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
