from typing import List
import unittest

# https://leetcode.com/problems/count-days-without-meetings/

# python3 -m unittest intervals/3169-count-days-without-meetings.py


class Solution(unittest.TestCase):
    # Approach: Sorting
    # Time: O(NLogN)
    # Space: O(N)
    def countDays(self, days: int, meetings: List[List[int]]) -> int:
        meetings.sort()
        last = 0
        for start, end in meetings:
            start = max(start, last + 1)
            if end >= start:
                days -= end - start + 1
            last = max(last, end)
        return days

    def test(self):
        for days, meetings, expected in [
            (10, [[5, 7], [1, 3], [9, 10]], 2),
            (5, [[2, 4], [1, 3]], 1),
            (6, [[1, 6]], 0),
            (14, [[6, 11], [7, 13], [8, 9], [5, 8], [3, 13], [11, 13], [1, 3], [5, 10], [8, 13], [3, 9]], 1),
        ]:
            output = self.countDays(days, meetings)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
