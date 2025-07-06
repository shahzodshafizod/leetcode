from typing import List
import unittest

# https://leetcode.com/problems/maximum-ascending-subarray-sum/

# python3 -m unittest arrays/1800-maximum-ascending-subarray-sum.py


class Solution(unittest.TestCase):
    def maxAscendingSum(self, nums: List[int]) -> int:
        max_sum = total = nums[0]
        for idx in range(1, len(nums)):
            if nums[idx - 1] < nums[idx]:
                total += nums[idx]
            else:
                total = nums[idx]
            max_sum = max(max_sum, total)
        return max_sum

    def test(self):
        for nums, expected in [
            ([10, 20, 30, 5, 10, 50], 65),
            ([10, 20, 30, 40, 50], 150),
            ([12, 17, 15, 13, 10, 11, 12], 33),
        ]:
            output = self.maxAscendingSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
