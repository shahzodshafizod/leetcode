from typing import List
import unittest

# https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/

# python3 -m unittest greedy/2099-find-subsequence-of-length-k-with-the-largest-sum.py


class Solution(unittest.TestCase):
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        return [
            num
            for idx, num in sorted(sorted(enumerate(nums), key=lambda x: x[1], reverse=True)[:k], key=lambda x: x[0])
        ]

    def test(self):
        for nums, k, expected in [
            ([2, 1, 3, 3], 2, [3, 3]),
            ([-1, -2, 3, 4], 3, [-1, 3, 4]),
            ([3, 4, 3, 3], 2, [3, 4]),
        ]:
            output = self.maxSubsequence(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
