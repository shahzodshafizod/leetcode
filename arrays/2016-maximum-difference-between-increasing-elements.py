from typing import List
import unittest

# https://leetcode.com/problems/maximum-difference-between-increasing-elements/

# python3 -m unittest arrays/2016-maximum-difference-between-increasing-elements.py

class Solution(unittest.TestCase):
    def maximumDifference(self, nums: List[int]) -> int:
        premin, diff = nums[0], -1
        for curr in nums:
            if curr > premin:
                diff = max(diff, curr - premin)
            else:
                premin = curr
        return diff

    def test(self):
        for nums, expected in [
            ([7,1,5,4], 4),
            ([9,4,3,2], -1),
            ([1,5,2,10], 9),
        ]:
            output = self.maximumDifference(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
