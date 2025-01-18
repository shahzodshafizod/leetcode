from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/

# python3 -m unittest matrices/1368-minimum-cost-to-make-at-least-one-valid-path-in-a-grid.py

class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(mxn)
    # Space: O(mxn)
    def minCost(self, grid: List[List[int]]) -> int:
        directions = [(0,1), (0,-1), (1,0), (-1,0)]
        m, n = len(grid), len(grid[0])
        queue = deque([(0, 0, 0)]) # (row, col, cost)
        mincost = [[float("inf")] * n for _ in range(m)]
        while queue:
            row, col, cost = queue.popleft()
            if min(row, col) < 0 or row == m or col == n or cost >= mincost[row][col]:
                continue
            mincost[row][col] = cost
            if row == m-1 and col == n-1:
                break
            for dir, (dr, dc) in enumerate(directions):
                if dir+1 == grid[row][col]:
                    queue.appendleft((row+dr, col+dc, cost))
                else:
                    queue.append((row+dr, col+dc, cost+1))
        return mincost[m-1][n-1]

    def test(self):
        for grid, expected in [
            ([[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]], 3),
            ([[1,1,3],[3,2,2],[1,1,4]], 0),
            ([[1,2],[4,3]], 1),
            ([[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]], 3),
            ([[1,1,3],[3,2,2],[1,1,4]], 0),
            ([[1,1,3],[4,3,2],[4,2,4]], 1),
        ]:
            output = self.minCost(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
