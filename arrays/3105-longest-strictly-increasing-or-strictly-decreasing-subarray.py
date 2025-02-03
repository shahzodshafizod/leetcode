from typing import List
import unittest

# https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/

# python3 -m unittest arrays/3105-longest-strictly-increasing-or-strictly-decreasing-subarray.py

class Solution(unittest.TestCase):
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        length = 1
        inc, dec = 1, 1
        for idx in range(1, len(nums)):
            if nums[idx-1] < nums[idx]:
                inc += 1
                dec = 1
            elif nums[idx-1] > nums[idx]:
                dec += 1
                inc = 1
            else:
                inc = dec = 1
            length = max(length, inc, dec)
        return length

    def test(self):
        for nums, expected in [
            ([1,4,3,3,2], 2),
            ([3,3,3,3], 1),
            ([3,2,1], 3),
        ]:
            output = self.longestMonotonicSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
