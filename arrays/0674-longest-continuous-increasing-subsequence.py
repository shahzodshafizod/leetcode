from typing import List
import unittest

# https://leetcode.com/problems/longest-continuous-increasing-subsequence/

# python3 -m unittest arrays/0674-longest-continuous-increasing-subsequence.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check all subarrays
    # # For each starting position, extend while increasing
    # # Time: O(N^2) - nested loops checking all possible subarrays
    # # Space: O(1) - only storing counters
    # def findLengthOfLCIS(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     max_length = 1
    
    #     for start in range(n):
    #         length = 1
    #         for end in range(start + 1, n):
    #             if nums[end] > nums[end - 1]:
    #                 length += 1
    #                 max_length = max(max_length, length)
    #             else:
    #                 break
    
    #     return max_length

    # Approach #2: Optimized - Single pass with counter
    # Track current sequence length, reset when not increasing
    # Time: O(N) - single pass through array
    # Space: O(1) - only storing two counters
    def findLengthOfLCIS(self, nums: List[int]) -> int:
        if not nums:
            return 0

        max_length = 1
        current_length = 1

        for i in range(1, len(nums)):
            if nums[i] > nums[i - 1]:
                current_length += 1
                max_length = max(max_length, current_length)
            else:
                current_length = 1

        return max_length

    def test(self):
        for nums, expected in [
            ([1, 3, 5, 4, 7], 3),
            ([2, 2, 2, 2, 2], 1),
            ([1], 1),
            ([1, 2, 3, 4, 5], 5),
            ([5, 4, 3, 2, 1], 1),
            ([1, 3, 5, 4, 7, 8, 9], 4),
        ]:
            output = self.findLengthOfLCIS(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
