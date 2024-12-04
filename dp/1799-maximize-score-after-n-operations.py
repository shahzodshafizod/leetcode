from typing import List
from math import gcd
import unittest

# https://leetcode.com/problems/maximize-score-after-n-operations/

# python3 -m unittest dp/1799-maximize-score-after-n-operations.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(n^2 + 2^n + log(max(nums)))
    # # Space: O(2^n)
    # def maxScore(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     memo = [None] * (1<<n)
    #     def dp(mask: int, op: int) -> int:
    #         if memo[mask] != None:
    #             return memo[mask]
    #         memo[mask] = 0
    #         for i in range(n):
    #             for j in range(i+1,n):
    #                 if mask & (1<<i) or mask & (1<<j):
    #                     continue
    #                 new_mask = mask | (1<<i) | (1<<j)
    #                 memo[mask] = max(
    #                     memo[mask],
    #                     op * gcd(nums[i], nums[j]) + dp(new_mask, op+1)
    #                 )
    #         return memo[mask]
    #     return dp(0, 1)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(n^2 + 2^n + log(max(nums)))
    # Space: O(2^n)
    def maxScore(self, nums: List[int]) -> int:
        n = len(nums)
        limit = 1<<n
        dp = [0] * limit
        for mask in range(limit):
            count = mask.bit_count()
            if count&1 == 1: continue
            for i in range(n):
                if mask&(1<<i) == 0: continue
                for j in range(i+1, n):
                    if mask&(1<<j) == 0: continue
                    new_mask = mask ^ (1<<i) ^ (1<<j)
                    dp[mask] = max(
                        dp[mask],
                        count//2 * gcd(nums[i], nums[j]) + dp[new_mask]
                    )
        return dp[limit-1]

    def test(self) -> None:
        for nums, expected in [
            ([1,2], 1),
            ([2,3,4,6], 8),
            ([3,4,6,8], 11),
            ([1,2,3,4,5,6], 14),
            ([9,17,16,15,18,13,18,20], 91),
            ([697035,181412,384958,575458], 869),
            ([109497,983516,698308,409009,310455,528595,524079,18036,341150,641864,913962,421869,943382,295019], 527),
        ]:
            output = self.maxScore(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
