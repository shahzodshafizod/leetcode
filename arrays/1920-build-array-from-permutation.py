from typing import List
import unittest

# https://leetcode.com/problems/build-array-from-permutation/

# python3 -m unittest arrays/1920-build-array-from-permutation.py

class Solution(unittest.TestCase):
    def buildArray(self, nums: List[int]) -> List[int]:
        return [nums[nums[_]] for _ in range(len(nums))]

    def test(self):
        for nums, expected in [
            ([0,2,1,5,3,4], [0,1,2,4,5,3]),
            ([5,0,1,2,3,4], [4,5,0,1,2,3]),
        ]:
            output = self.buildArray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
