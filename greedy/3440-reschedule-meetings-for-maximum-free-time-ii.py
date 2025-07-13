from typing import List
import unittest

# https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-ii/

# python3 -m unittest greedy/3440-reschedule-meetings-for-maximum-free-time-ii.py


class Solution(unittest.TestCase):
    # Approach: Greedy
    # Time: O(n)
    # Space: O(1)
    def maxFreeTime(self, eventTime: int, startTime: List[int], endTime: List[int]) -> int:
        free, n = 0, len(startTime)
        # left to right
        slot, left = 0, 0
        for i in range(n):
            meeting = endTime[i] - startTime[i]
            if slot >= meeting:
                meeting = 0
            right = eventTime if i + 1 == n else startTime[i + 1]
            free = max(free, right - left - meeting)
            slot = max(slot, startTime[i] - left)
            left = endTime[i]
        # right to left
        slot, right = 0, eventTime
        for i in range(n - 1, -1, -1):
            meeting = endTime[i] - startTime[i]
            if slot >= meeting:
                meeting = 0
            left = 0 if i == 0 else endTime[i - 1]
            free = max(free, right - left - meeting)
            slot = max(slot, right - endTime[i])
            right = startTime[i]
        return free

    def test(self):
        for eventTime, startTime, endTime, expected in [
            (5, [1, 3], [2, 5], 2),
            (10, [0, 7, 9], [1, 8, 10], 7),
            (10, [0, 3, 7, 9], [1, 4, 8, 10], 6),
            (5, [0, 1, 2, 3, 4], [1, 2, 3, 4, 5], 0),
            (34, [0, 17], [14, 19], 18),
            (37, [5, 14, 27, 34], [13, 18, 31, 37], 16),
            (54, [9, 24, 45, 50, 53], [15, 26, 50, 53, 54], 30),
            (41, [17, 24], [19, 25], 24),
            (86, [22, 82], [66, 85], 38),
        ]:
            output = self.maxFreeTime(eventTime, startTime, endTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
