from typing import List
import unittest

# https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/

# python3 -m unittest prefixsums/1749-maximum-absolute-sum-of-any-subarray.py


class Solution(unittest.TestCase):
    # Approach: Kadane's algorithm
    # Time: O(n)
    # Space: O(1)
    def maxAbsoluteSum(self, nums: List[int]) -> int:
        pos_prefix_sum, max_sum = 0, 0
        neg_prefix_sum, min_sum = 0, 0
        for num in nums:
            pos_prefix_sum = max(0, pos_prefix_sum + num)
            max_sum = max(max_sum, pos_prefix_sum)
            neg_prefix_sum = min(0, neg_prefix_sum + num)
            min_sum = min(min_sum, neg_prefix_sum)
        return max(max_sum, -min_sum)

    def test(self):
        for nums, expected in [
            ([1, -3, 2, 3, -4], 5),
            ([2, -5, 1, -4, 3, -2], 8),
        ]:
            output = self.maxAbsoluteSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
