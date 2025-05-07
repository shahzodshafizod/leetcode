import heapq
from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/

# python3 -m unittest graphs/3341-find-minimum-time-to-reach-last-room-i.py

class Solution(unittest.TestCase):
    # Approach: Shortest Path + Dijkstra
    # Time: O(nmlog(nm))
    # Space: O(nm)
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        m, n = len(moveTime), len(moveTime[0])
        directions = [1,0,-1,0,1]
        min_heap = [(0,0,0)] # (time,row,col)
        seen = set([0,0])
        while min_heap:
            time, row, col = heapq.heappop(min_heap)
            if row == m-1 and col == n-1:
                return time
            for dir in range(1, 5):
                new_r = row + directions[dir-1]
                new_c = col + directions[dir]
                if min(new_r, new_c) >= 0 and new_r < m and new_c < n and (new_r,new_c) not in seen:
                    heapq.heappush(min_heap, (max(time,moveTime[new_r][new_c])+1, new_r, new_c))
                    seen.add((new_r, new_c))
        return -1

    def test(self):
        for moveTime, expected in [
            ([[0,4],[4,4]], 6),
            ([[0,0,0],[0,0,0]], 3),
            ([[0,1],[1,2]], 3),
        ]:
            output = self.minTimeToReach(moveTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
