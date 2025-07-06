from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/count-number-of-bad-pairs/

# python3 -m unittest hashes/2364-count-number-of-bad-pairs.py


class Solution(unittest.TestCase):
    def countBadPairs(self, nums: List[int]) -> int:
        n = len(nums)
        pairs = defaultdict(int)
        goods = 0
        for idx, num in enumerate(nums):
            goods += pairs[num - idx]
            pairs[num - idx] += 1
        bads = n * (n - 1) // 2 - goods
        return bads

    def test(self):
        for nums, expected in [
            ([4, 1, 3, 3], 5),
            ([1, 2, 3, 4, 5], 0),
        ]:
            output = self.countBadPairs(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
