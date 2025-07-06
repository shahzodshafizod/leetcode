import heapq
from typing import List
import unittest

# https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/

# python3 -m unittest graphs/2577-minimum-time-to-visit-a-cell-in-a-grid.py


class Solution(unittest.TestCase):
    # Approach: Dijkstra's
    # Time: O(mnlog(mn))
    # Space: O(mn)
    def minimumTime(self, grid: List[List[int]]) -> int:
        if min(grid[0][1], grid[1][0]) > 1:
            return -1
        directions = [-1, 0, 1, 0, -1]
        m, n = len(grid), len(grid[0])
        seen = [[False] * n for _ in range(m)]
        min_heap = [(0, 0, 0)]  # (time, row, col)
        while min_heap:
            time, row, col = heapq.heappop(min_heap)
            if (row, col) == (m - 1, n - 1):
                return time
            for d in range(1, 5):
                nr = row + directions[d - 1]
                nc = col + directions[d]
                if min(nr, nc) < 0 or nr == m or nc == n or seen[nr][nc]:
                    continue
                wait = 1 if abs(grid[nr][nc] - time) & 1 == 0 else 0
                ntime = max(time + 1, grid[nr][nc] + wait)
                heapq.heappush(min_heap, (ntime, nr, nc))
                seen[nr][nc] = True
        return -1

    def test(self):
        for grid, expected in [
            ([[0, 2, 4], [0, 5, 5], [5, 4, 3]], 6),
            ([[0, 2, 4], [3, 2, 1], [1, 0, 4]], -1),
            ([[0, 5, 1], [0, 7, 6], [7, 7, 1]], 8),
            ([[0, 1, 3, 2], [5, 1, 2, 5], [4, 3, 8, 6]], 7),
            ([[0, 7, 6, 6], [1, 6, 8, 6], [1, 5, 8, 3], [4, 7, 0, 1]], 10),
            ([[0, 1, 12], [19, 39, 97], [75, 88, 33], [21, 2, 88]], 89),
            # ([[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[25,99,99,99,99,99,99,99,99,99,99,99,99,99,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,99,99,99,99,99,99,99,99,99,99,99,99,99,99],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]], 42),
        ]:
            output = self.minimumTime(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
