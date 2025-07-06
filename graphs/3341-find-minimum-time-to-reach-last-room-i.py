import heapq
from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/

# python3 -m unittest graphs/3341-find-minimum-time-to-reach-last-room-i.py


class Solution(unittest.TestCase):
    # Approach: Shortest Path + Dijkstra
    # Time: O(nm log nm)
    # Space: O(nm)
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        m, n = len(moveTime), len(moveTime[0])
        dirs = [1, 0, -1, 0, 1]
        min_heap = [(0, 0, 0)]  # (time,row,col)
        dist = [[float("inf")] * n for _ in range(m)]
        dist[0][0] = 0
        while min_heap:
            time, row, col = heapq.heappop(min_heap)
            if row == m - 1 and col == n - 1:
                break
            for d in range(1, 5):
                new_r = row + dirs[d - 1]
                new_c = col + dirs[d]
                if min(new_r, new_c) < 0 or new_r >= m or new_c >= n:
                    continue
                new_t = max(time, moveTime[new_r][new_c]) + 1
                if new_t < dist[new_r][new_c]:
                    heapq.heappush(min_heap, (new_t, new_r, new_c))
                    dist[new_r][new_c] = new_t
        return dist[m - 1][n - 1]

    def test(self):
        for moveTime, expected in [
            ([[0, 4], [4, 4]], 6),
            ([[0, 0, 0], [0, 0, 0]], 3),
            ([[0, 1], [1, 2]], 3),
        ]:
            output = self.minTimeToReach(moveTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
