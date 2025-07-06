from typing import List
import unittest

# https://leetcode.com/problems/longest-nice-subarray/

# python3 -m unittest slidingwindows/2401-longest-nice-subarray.py


class Solution(unittest.TestCase):
    # Approach: Sliding Window + Bit Manipulation
    # Time: O(n)
    # Space: O(1)
    def longestNiceSubarray(self, nums: List[int]) -> int:
        mask, length = 0, 0
        left = 0
        for right, num in enumerate(nums):
            while left < right and num & mask:
                mask ^= nums[left]
                left += 1
            mask |= num
            length = max(length, right - left + 1)
        return length

    def test(self):
        for nums, expected in [
            ([1, 3, 8, 48, 10], 3),
            ([3, 1, 5, 11, 13], 1),
        ]:
            output = self.longestNiceSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
