from typing import List
import unittest

# https://leetcode.com/problems/array-partition/

# python3 -m unittest arrays/0561-array-partition.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Same as optimized
    # # This is already optimal - sorting ensures max sum of minimums.
    # # Time: O(n log n)
    # # Space: O(1)
    # def arrayPairSum(self, nums: List[int]) -> int:
    #     nums.sort()
    #     total = 0
    #     for i in range(0, len(nums), 2):
    #         total += nums[i]
    #     return total

    # Approach #2: Sort and sum odd indices
    # Sort array, then pair adjacent elements. Sum of mins is sum of elements at even indices.
    # Time: O(n log n) - sorting
    # Space: O(1) - ignoring sort space
    def arrayPairSum(self, nums: List[int]) -> int:
        nums.sort()
        return sum(nums[i] for i in range(0, len(nums), 2))

    def test(self):
        for nums, expected in [
            ([1, 4, 3, 2], 4),  # (1,2) + (3,4) = 1+3 = 4
            ([6, 2, 6, 5, 1, 2], 9),  # (1,2) + (2,5) + (6,6) = 1+2+6 = 9
            ([1, 2], 1),
        ]:
            output = self.arrayPairSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
