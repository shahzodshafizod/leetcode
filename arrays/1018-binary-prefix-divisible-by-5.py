from typing import List
import unittest

# https://leetcode.com/problems/binary-prefix-divisible-by-5/

# python3 -m unittest arrays/1018-binary-prefix-divisible-by-5.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Build full binary number
    # # Time: O(n^2) - for each position, build number from scratch
    # # Space: O(n) - result array
    # def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
    #     result: List[bool] = []
    #     for i in range(len(nums)):
    #         value = 0
    #         for j in range(i + 1):
    #             value = value * 2 + nums[j]
    #         result.append(value % 5 == 0)
    #     return result

    # Approach #2: Track running value modulo 5
    # Time: O(n) - single pass
    # Space: O(n) - result array
    def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
        result: List[bool] = []
        value = 0
        for bit in nums:
            value = (value * 2 + bit) % 5
            result.append(value == 0)
        return result

    def test(self):
        for nums, expected in [
            ([0, 1, 1], [True, False, False]),
            ([1, 1, 1], [False, False, False]),
            ([0, 1, 1, 1, 1, 1], [True, False, False, False, True, False]),
        ]:
            output = self.prefixesDivBy5(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
