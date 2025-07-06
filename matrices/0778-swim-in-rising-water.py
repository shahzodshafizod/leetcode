import unittest
from typing import List
import heapq

# https://leetcode.com/problems/swim-in-rising-water/

# python3 -m unittest matrices/0778-swim-in-rising-water.py


class Solution(unittest.TestCase):
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        minHeap = [(grid[0][0], 0, 0)]
        grid[0][0] = -1
        directions = [-1, 0, 1, 0, -1]
        t = 0
        while minHeap:
            elevation, row, col = heapq.heappop(minHeap)
            t = max(t, elevation)
            if row == n - 1 and col == n - 1:
                break
            for d in range(1, 5):
                r, c = row + directions[d - 1], col + directions[d]
                if min(r, c) >= 0 and max(r, c) < n and grid[r][c] >= 0:
                    heapq.heappush(minHeap, (grid[r][c], r, c))
                    grid[r][c] = -1
        return t

    def test(self):
        for grid, expected in [
            ([[0, 2], [1, 3]], 3),
            ([[0, 1, 2, 3, 4], [24, 23, 22, 21, 5], [12, 13, 14, 15, 16], [11, 17, 18, 19, 20], [10, 9, 8, 7, 6]], 16),
        ]:
            output = self.swimInWater(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output {output}")
