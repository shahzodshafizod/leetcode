from typing import List
import unittest

# https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/

# python3 -m unittest dp/3202-find-the-maximum-length-of-valid-subsequence-ii.py


class Solution(unittest.TestCase):
    def maximumLength(self, nums: List[int], k: int) -> int:
        dp = [[0] * k for _ in range(k)]
        length = 0
        for curr in nums:
            curr %= k
            for prev in range(k):
                dp[curr][prev] = dp[prev][curr] + 1
                length = max(length, dp[curr][prev])
        return length

    def test(self):
        for nums, k, expected in [
            ([1, 2, 3, 4, 5], 2, 5),
            ([1, 4, 2, 3, 1, 4], 3, 4),
        ]:
            output = self.maximumLength(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
