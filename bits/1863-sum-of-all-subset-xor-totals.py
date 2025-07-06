from functools import cache  # pylint: disable=unused-import
from typing import List
import unittest

# https://leetcode.com/problems/sum-of-all-subset-xor-totals/

# python3 -m unittest bits/1863-sum-of-all-subset-xor-totals.py


class Solution(unittest.TestCase):
    # # Approach #1: Backtracking
    # # Time: O(2^n)
    # # Space: O(n)
    # def subsetXORSum(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     @cache
    #     def dp(idx: int, total: int) -> int:
    #         if idx == n:
    #             return total
    #         return dp(idx+1, total^nums[idx]) + dp(idx+1, total)
    #     return dp(0, 0)

    # Approach #2: Bit Manipulation
    # Time: O(n)
    # Space: O(1)
    def subsetXORSum(self, nums: List[int]) -> int:
        total = 0
        for num in nums:
            total |= num
        return total << (len(nums) - 1)

    def test(self):
        for nums, expected in [
            ([1, 3], 6),
            ([5, 1, 6], 28),
            ([3, 4, 5, 6, 7, 8], 480),
        ]:
            output = self.subsetXORSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
