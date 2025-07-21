from typing import List
import unittest

# https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/

# python3 -m unittest slidingwindows/3439-reschedule-meetings-for-maximum-free-time-i.py


class Solution(unittest.TestCase):
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        free, busy, n = 0, 0, len(startTime)
        for i in range(n):
            busy += endTime[i] - startTime[i]
            left = 0 if i + 1 <= k else endTime[i - k]
            right = eventTime if i + 1 == n else startTime[i + 1]
            free = max(free, right - left - busy)
            if i + 1 >= k:
                busy -= endTime[i + 1 - k] - startTime[i + 1 - k]
        return free

    def test(self):
        for eventTime, k, startTime, endTime, expected in [
            (5, 1, [1, 3], [2, 5], 2),
            (10, 1, [0, 2, 9], [1, 4, 10], 6),
            (5, 2, [0, 1, 2, 3, 4], [1, 2, 3, 4, 5], 0),
        ]:
            output = self.maxFreeTime(eventTime, k, startTime, endTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
