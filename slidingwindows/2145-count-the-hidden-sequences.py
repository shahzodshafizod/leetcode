from typing import List
import unittest

# https://leetcode.com/problems/count-the-hidden-sequences/

# python3 -m unittest slidingwindows/2145-count-the-hidden-sequences.py


class Solution(unittest.TestCase):
    def numberOfArrays(self, differences: List[int], lower: int, upper: int) -> int:
        peak_val, deep_val = upper, lower
        peak, deep = upper, lower
        for diff in differences:
            peak_val += diff
            peak = min(peak, upper - (peak_val - upper))
            deep_val += diff
            deep = max(deep, lower + (lower - deep_val))
            if peak < deep:
                return 0
        return peak - deep + 1

    def test(self):
        for differences, lower, upper, expected in [
            ([1, -3, 4], 1, 6, 2),
            ([3, -4, 5, 1, -2], -4, 5, 4),
            ([4, -7, 2], 3, 6, 0),
            ([2, -2, 3, -3, 4, -4, 5, -5], 0, 10, 6),
            ([1, -1, 2, -2, 3, -3, 4, -4, 5, -5], -3, 7, 6),
            ([3, -2, 1, -3, 2, -1, 3, -3], 1, 8, 4),
            ([1, 1, 1, 1, 1, -5], 0, 5, 1),
            ([5, -4, 3, -2, 1, -1, -2, 2, -3, 3], -2, 6, 3),
            ([1, 2, -1, -2, 3, -3, 4, -4, 5, -5], 0, 10, 6),
            ([-40], -46, 53, 60),
            ([83702, -5216], -82788, 14602, 13689),
        ]:
            output = self.numberOfArrays(differences, lower, upper)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
