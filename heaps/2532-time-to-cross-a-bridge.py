import heapq
from typing import List
import unittest

# https://leetcode.com/problems/time-to-cross-a-bridge/

# python3 -m unittest heaps/2532-time-to-cross-a-bridge.py


class Solution(unittest.TestCase):
    # Approach: Simulation
    # Time: O(n log n)
    # Space: O(n)
    def findCrossingTime(self, n: int, k: int, time: List[List[int]]) -> int:
        l2r: List[tuple[int, int]] = []  # ready to move from left to right: (-left-right, -index)
        r2l: List[tuple[int, int]] = []  # ready to move from right to left: (-left-right, -index)
        pick: List[tuple[int, int]] = []  # ready to pick: (time, index)
        put: List[tuple[int, int]] = []  # ready to put: (time, index)

        # Initially, all k workers are waiting on the left side of the bridge.
        for i in range(k):
            heapq.heappush(l2r, (-time[i][0] - time[i][2], -i))

        minutes = 0
        while n:
            while put and put[0][0] <= minutes:
                _, i = heapq.heappop(put)
                heapq.heappush(l2r, (-time[i][0] - time[i][2], -i))
            while pick and pick[0][0] <= minutes:
                _, i = heapq.heappop(pick)
                heapq.heappush(r2l, (-time[i][0] - time[i][2], -i))
            if r2l:
                _, i = heapq.heappop(r2l)
                minutes += time[-i][2]
                heapq.heappush(put, (minutes + time[-i][3], -i))
                n -= 1
            elif l2r and n > len(pick) + len(r2l):
                _, i = heapq.heappop(l2r)
                minutes += time[-i][0]
                heapq.heappush(pick, (minutes + time[-i][1], -i))
            else:
                minutes = (1 << 31) - 1
                if put and n > len(pick) + len(r2l):
                    minutes = min(minutes, put[0][0])
                if pick:
                    minutes = min(minutes, pick[0][0])
        return minutes

    def test(self):
        for n, k, time, expected in [
            (1, 3, [[1, 1, 2, 1], [1, 1, 3, 1], [1, 1, 4, 1]], 6),
            (3, 2, [[1, 5, 1, 8], [10, 10, 10, 10]], 37),
        ]:
            output = self.findCrossingTime(n, k, time)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
