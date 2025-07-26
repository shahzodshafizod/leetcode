from typing import List
import unittest

# https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/

# python3 -m unittest greedy/3487-maximum-unique-subarray-sum-after-deletion.py


class Solution(unittest.TestCase):
    def maxSum(self, nums: List[int]) -> int:
        uniq = set(x for x in nums if x > 0)
        return sum(uniq) if uniq else max(nums)

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4, 5], 15),
            ([1, 1, 0, 1, 1], 1),
            ([1, 2, -1, -2, 1, 0, -1], 3),
        ]:
            output = self.maxSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
