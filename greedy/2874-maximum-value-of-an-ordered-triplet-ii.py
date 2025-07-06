# pylint: disable=duplicate-code
from typing import List
import unittest

# https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/

# python3 -m unittest greedy/2874-maximum-value-of-an-ordered-triplet-ii.py


class Solution(unittest.TestCase):
    def maximumTripletValue(self, nums: List[int]) -> int:
        value, dmax, lmax = 0, 0, 0
        for num in nums:
            value = max(value, dmax * num)
            dmax = max(dmax, lmax - num)
            lmax = max(lmax, num)
        return value

    def test(self):
        for nums, expected in [
            ([12, 6, 1, 2, 7], 77),
            ([1, 10, 3, 4, 19], 133),
            ([1, 2, 3], 0),
        ]:
            output = self.maximumTripletValue(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
