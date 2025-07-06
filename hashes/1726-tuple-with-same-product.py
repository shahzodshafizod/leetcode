from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/tuple-with-same-product/

# python3 -m unittest hashes/1726-tuple-with-same-product.py


class Solution(unittest.TestCase):
    def tupleSameProduct(self, nums: List[int]) -> int:
        squares = defaultdict(int)
        n = len(nums)
        for i in range(n):
            for j in range(i + 1, n):
                squares[nums[i] * nums[j]] += 1
        return sum(8 * (cnt * (cnt - 1) // 2) for cnt in squares.values())

    def test(self):
        for nums, expected in [
            ([2, 3, 4, 6], 8),
            ([1, 2, 4, 5, 10], 16),
            ([2, 3, 5, 7], 0),
            ([2, 3, 4, 6, 8, 12], 40),
            ([1, 2, 4, 8, 16, 32], 56),
            ([10, 5, 15, 8, 6, 12, 20, 4], 72),
            ([20, 10, 6, 24, 15, 5, 4, 30], 48),
            ([19, 32, 39, 38, 43, 47, 9, 5, 35, 22], 0),
            ([1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192], 1288),
        ]:
            output = self.tupleSameProduct(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
