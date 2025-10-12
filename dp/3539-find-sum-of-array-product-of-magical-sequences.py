from functools import lru_cache
from typing import List
import unittest
import math

# https://leetcode.com/problems/find-sum-of-array-product-of-magical-sequences/

# python3 -m unittest dp/3539-find-sum-of-array-product-of-magical-sequences.py


class Solution(unittest.TestCase):
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        MOD = 10**9 + 7

        @lru_cache(None)
        def dp(i: int, m: int, k: int, mask: int) -> int:
            if m == 0 and mask.bit_count() == k:
                return 1
            if m == 0 or i == len(nums):
                return 0
            total = dp(i + 1, m, k - (mask & 1), mask >> 1)
            for freq in range(1, m + 1):
                ways = math.comb(m, freq) * pow(nums[i], freq, MOD) % MOD
                nmask = mask + freq
                nxt = dp(i + 1, m - freq, k - (nmask & 1), nmask >> 1)
                total = (total + ways * nxt) % MOD
            return total

        return dp(0, m, k, 0)

    def test(self):
        for m, k, nums, expected in [
            (5, 5, [1, 10, 100, 10000, 1000000], 991600007),
            (2, 2, [5, 4, 3, 2, 1], 170),
            (1, 1, [28], 28),
        ]:
            output = self.magicalSum(m, k, nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
