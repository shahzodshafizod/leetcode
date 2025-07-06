import heapq
from typing import List
import unittest

# https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/

# python3 -m unittest heaps/2503-maximum-number-of-points-from-grid-queries.py


class Solution(unittest.TestCase):
    # Approach: Sorting Queries + Min-Heap Expansion
    # Time: O(klogk+n⋅mlog(n⋅m))
    # Space: O(n⋅m+k)
    def maxPoints(self, grid: List[List[int]], queries: List[int]) -> List[int]:
        dirs = [1, 0, -1, 0, 1]
        m, n = len(grid), len(grid[0])
        seen = [[False] * n for _ in range(m)]
        queue = [(grid[0][0], 0, 0)]  # val, row, col
        seen[0][0] = True
        points = 0
        answer = [0] * len(queries)
        for limit, index in sorted([(q, i) for i, q in enumerate(queries)]):
            while queue and queue[0][0] < limit:
                _, row, col = heapq.heappop(queue)
                for d in range(1, 5):
                    nr = row + dirs[d - 1]
                    nc = col + dirs[d]
                    if 0 <= nr < m and 0 <= nc < n and not seen[nr][nc]:
                        heapq.heappush(queue, (grid[nr][nc], nr, nc))
                        seen[nr][nc] = True
                points += 1
            answer[index] = points
        return answer

    def test(self):
        for grid, queries, expected in [
            ([[1, 2, 3], [2, 5, 7], [3, 5, 1]], [5, 6, 2], [5, 8, 1]),
            ([[5, 2, 1], [1, 1, 2]], [3], [0]),
        ]:
            output = self.maxPoints(grid, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
