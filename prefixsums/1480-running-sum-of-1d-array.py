from typing import List
import unittest

# https://leetcode.com/problems/running-sum-of-1d-array/

# python3 -m unittest prefixsums/1480-running-sum-of-1d-array.py


class Solution(unittest.TestCase):
    def runningSum(self, nums: List[int]) -> List[int]:
        for i in range(1, len(nums)):
            nums[i] += nums[i - 1]
        return nums

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4], [1, 3, 6, 10]),
            ([1, 1, 1, 1, 1], [1, 2, 3, 4, 5]),
            ([3, 1, 2, 10, 1], [3, 4, 6, 16, 17]),
        ]:
            output = self.runningSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
