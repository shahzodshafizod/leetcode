from typing import List
import unittest

# https://leetcode.com/problems/largest-number-at-least-twice-of-others/

# python3 -m unittest arrays/0747-largest-number-at-least-twice-of-others.py


class Solution(unittest.TestCase):
    # # Approach 1: Two-Pass Solution
    # # First pass: Find the largest element and its index
    # # Second pass: Check if largest is at least twice every other element
    # # Time: O(n) - two passes through the array
    # # Space: O(1) - only using a few variables
    # def dominantIndex(self, nums: List[int]) -> int:
    #     # First pass: find max and its index
    #     max_val = max(nums)
    #     max_idx = nums.index(max_val)

    #     # Second pass: check if max is at least twice every other element
    #     for num in nums:
    #         if num != max_val and max_val < 2 * num:
    #             return -1

    #     return max_idx

    # Approach 2: One-Pass Solution (Optimized)
    # Track both the largest and second largest elements in one pass
    # Then check if largest >= 2 * second_largest
    # Time: O(n) - single pass through the array
    # Space: O(1) - only using a few variables
    def dominantIndex(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0

        max_val = -1
        max_idx = -1
        second_max = -1

        # Find both max and second max in one pass
        for i, num in enumerate(nums):
            if num > max_val:
                second_max = max_val
                max_val = num
                max_idx = i
            elif num > second_max:
                second_max = num

        # Check if max is at least twice the second max
        if max_val >= 2 * second_max:
            return max_idx
        return -1

    def test(self):
        for nums, expected in [
            ([3, 6, 1, 0], 1),
            ([1, 2, 3, 4], -1),
            ([1], 0),
            ([0, 0, 0, 1], 3),
            ([1, 0], 0),
            ([0, 0, 3, 2], -1),
        ]:
            output = self.dominantIndex(nums[:])
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
