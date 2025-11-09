from typing import List
import unittest

# https://leetcode.com/problems/minimum-replacements-to-sort-the-array/

# python3 -m unittest greedy/2366-minimum-replacements-to-sort-the-array.py


class Solution(unittest.TestCase):
    # Approach: Greedy
    # Time: O(n)
    # Space: O(1)
    def minimumReplacement(self, nums: List[int]) -> int:
        ops = 0
        height = nums[-1]
        for i in range(len(nums) - 2, -1, -1):
            parts = (nums[i] + height - 1) // height
            ops += parts - 1
            height = nums[i] // parts
        return ops

    def test(self):
        for nums, expected in [
            ([3, 9, 3], 2),
            ([1, 2, 3, 4, 5], 0),
            ([19, 7, 2, 24, 11, 16, 1, 11, 23], 73),
        ]:
            output = self.minimumReplacement(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
