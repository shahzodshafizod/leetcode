from typing import List
import unittest

# https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/

# python3 -m unittest arrays/3423-maximum-difference-between-adjacent-elements-in-a-circular-array.py

class Solution(unittest.TestCase):
    def maxAdjacentDistance(self, nums: List[int]) -> int:
        return max(abs(nums[i]-nums[i-1]) for i in range(len(nums)))

    def test(self):
        for nums, expected in [
            ([1,2,4], 3),
            ([-5,-10,-5], 5),
        ]:
            output = self.maxAdjacentDistance(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
