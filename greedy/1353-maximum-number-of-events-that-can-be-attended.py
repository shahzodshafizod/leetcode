import heapq
import unittest
from typing import List

# https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/

# python3 -m unittest greedy/1353-maximum-number-of-events-that-can-be-attended.py


class Solution(unittest.TestCase):
    def maxEvents(self, events: List[List[int]]) -> int:
        events.sort()
        pq = []
        count, i, n = 0, 0, len(events)
        max_day = max(end for _, end in events)
        for day in range(1, max_day + 1):
            while i < n and events[i][0] == day:
                heapq.heappush(pq, events[i][1])
                i += 1
            while pq and pq[0] < day:
                heapq.heappop(pq)
            if pq:
                count += 1
                heapq.heappop(pq)
        return count

    def test(self):
        for events, expected in [
            ([[1, 10], [1, 10], [1, 10], [1, 10]], 4),
            ([[1, 100], [50, 100], [60, 70], [90, 95]], 4),
            ([[1, 1], [2, 2], [3, 3], [4, 4]], 4),
            ([[1, 2], [2, 2], [1, 2], [1, 2]], 2),
            ([[1, 2], [3, 4], [5, 6]], 3),
            ([[1, 2], [4, 5], [7, 8], [10, 11]], 4),
            ([[1, 1], [1, 1], [1, 1], [1, 1]], 1),
            ([[1, 3], [2, 4], [3, 5], [4, 6], [5, 7]], 5),
            ([[1, 3], [1, 3], [1, 3], [1, 3]], 3),
            ([[1, 3], [5, 8], [6, 10], [9, 11], [12, 15]], 5),
            ([[1, 2], [2, 3], [3, 4]], 3),
            ([[1, 2], [2, 3], [3, 4], [1, 2]], 4),
            ([[1, 4], [4, 4], [2, 2], [3, 4], [1, 1]], 4),
            ([[1, 100000]], 1),
            ([[1, 1], [1, 2], [1, 3], [1, 4], [1, 5], [1, 6], [1, 7]], 7),
            ([[1, 1], [2, 2], [3, 3], [4, 4], [5, 5]], 5),
            ([[1, 2], [1, 2], [3, 3], [1, 5], [1, 5]], 5),
            ([[1, 10], [2, 10], [3, 10], [4, 10]], 4),
            ([[1, 2], [1, 2], [1, 6], [1, 2], [1, 2]], 3),
            ([[1, 10], [2, 2], [2, 2], [2, 2], [2, 2]], 2),
            ([[1, 2], [2, 2], [3, 3], [3, 4], [3, 4]], 4),
            ([[1, 4], [1, 1], [2, 2], [3, 4], [4, 4]], 4),
            ([[1, 2], [2, 3], [3, 4], [1, 3], [2, 5]], 5),
            ([[1, 5], [1, 5], [1, 5], [2, 3], [2, 3]], 5),
        ]:
            output = self.maxEvents(events)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
