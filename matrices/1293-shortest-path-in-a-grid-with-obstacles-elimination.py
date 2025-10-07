from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

# python3 -m unittest matrices/1293-shortest-path-in-a-grid-with-obstacles-elimination.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(mnk)
    # Space: O(mnk)
    def shortestPath(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])
        directions = [0, -1, 0, 1, 0]
        chance = [[-1] * n for _ in range(m)]
        chance[0][0] = k
        queue = deque([(0, 0, k)])  # (row, col, obstacles)
        length = 0
        while queue:
            for _ in range(len(queue)):
                row, col, obs = queue.popleft()
                if row == m - 1 and col == n - 1:
                    return length
                for d in range(4):
                    nrow, ncol = row + directions[d], col + directions[d + 1]
                    if min(nrow, ncol) < 0 or nrow == m or ncol == n:
                        continue
                    nobs = obs - grid[nrow][ncol]
                    if nobs <= chance[nrow][ncol]:
                        continue
                    chance[nrow][ncol] = nobs
                    queue.append((nrow, ncol, nobs))
            length += 1
        return -1

    def test(self):
        for grid, k, expected in [
            ([[0, 0, 0], [1, 1, 0], [0, 0, 0], [0, 1, 1], [0, 0, 0]], 1, 6),
            ([[0, 1, 1], [1, 1, 1], [1, 0, 0]], 1, -1),
        ]:
            output = self.shortestPath(grid, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
