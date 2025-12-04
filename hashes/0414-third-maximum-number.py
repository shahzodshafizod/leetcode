from typing import List
import unittest

# https://leetcode.com/problems/third-maximum-number/

# python3 -m unittest hashes/0414-third-maximum-number.py


class Solution(unittest.TestCase):
    # Approach: Sorting
    # Time: O(nlogn)
    # Space: O(1)
    def thirdMax(self, nums: List[int]) -> int:
        nums = sorted(list(set(nums)), reverse=True)
        return nums[0] if len(nums) < 3 else nums[2]

    def test(self):
        for nums, expected in [
            ([3, 2, 1], 1),
            ([1, 2], 2),
            ([2, 2, 3, 1], 1),
        ]:
            output = self.thirdMax(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
