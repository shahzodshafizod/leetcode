import heapq
from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/

# python3 -m unittest graphs/3342-find-minimum-time-to-reach-last-room-ii.py

class Solution(unittest.TestCase):
    # Approach: Shortest Path + Dijkstra
    # Time: O(nm log nm)
    # Space: O(nm)
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        n, m = len(moveTime), len(moveTime[0])
        directions = [1,0,-1,0,1]
        dist = [[float("inf")] * m for _ in range(n)]
        dist[0][0] = 0
        min_heap = [(0,0,0)] # (time,row,col)
        while min_heap:
            time, row, col = heapq.heappop(min_heap)
            if row == n-1 and col == m-1:
                break
            for dir in range(1, 5):
                new_r = row + directions[dir-1]
                new_c = col + directions[dir]
                if min(new_r, new_c) < 0 or new_r >= n or new_c >= m:
                    continue
                new_t = max(time, moveTime[new_r][new_c]) + (row+col)%2 + 1
                if new_t < dist[new_r][new_c]:
                    heapq.heappush(min_heap, (new_t, new_r, new_c))
                    dist[new_r][new_c] = new_t
        return dist[n-1][m-1]

    def test(self):
        for moveTime, expected in [
            ([[0,4],[4,4]], 7),
            ([[0,0,0,0],[0,0,0,0]], 6),
            ([[0,1],[1,2]], 4),
        ]:
            output = self.minTimeToReach(moveTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
