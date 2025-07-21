from typing import List
import unittest

# https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/

# python3 -m unittest dp/3201-find-the-maximum-length-of-valid-subsequence-i.py


class Solution(unittest.TestCase):
    def maximumLength(self, nums: List[int]) -> int:
        cnt, end = [0] * 2, [0] * 2
        for num in nums:
            num %= 2
            cnt[num] += 1
            end[num] = end[1 - num] + 1
        return max(cnt[0], cnt[1], end[0], end[1])

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4], 4),
            ([1, 2, 1, 1, 2, 1, 2], 6),
            ([1, 3], 2),
        ]:
            output = self.maximumLength(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
