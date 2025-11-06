from collections import defaultdict
from typing import Dict, List
import unittest

# https://leetcode.com/problems/sum-of-good-subsequences/

# python3 -m unittest dp/3351-sum-of-good-subsequences.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force (Top-Down Dynamic Programming)
    # # Time: O(nn)
    # # Space: O(nn)
    # def sumOfGoodSubsequences(self, nums: List[int]) -> int:
    #     MOD = 10**9 + 7
    #     n = len(nums)
    #     memo: List[List[tuple[int, int]]] = [[(-1, -1)] * n for _ in range(n)]

    #     def dp(prev: int, curr: int) -> tuple[int, int]:
    #         if curr == n:
    #             return (0, 0)
    #         if memo[prev + 1][curr] != (-1, -1):
    #             return memo[prev + 1][curr]
    #         total, cnt = dp(prev, curr + 1)
    #         if prev == -1 or abs(nums[curr] - nums[prev]) == 1:
    #             nsum, ncnt = dp(curr, curr + 1)
    #             total = (total + nums[curr] * (ncnt + 1) + nsum) % MOD
    #             cnt = (cnt + ncnt + 1) % MOD
    #         memo[prev + 1][curr] = (total, cnt)
    #         return (total, cnt)

    #     result_sum, _ = dp(-1, 0)
    #     return result_sum

    # Approach 2: Bottom-Up Dynamic Programming (Tabulation)
    # Time: O(n)
    # Space: O(n)
    def sumOfGoodSubsequences(self, nums: List[int]) -> int:
        cnt: Dict[int, int] = defaultdict(int)
        total: Dict[int, int] = defaultdict(int)
        result = 0
        for num in nums:
            num_cnt = 1 + cnt[num - 1] + cnt[num + 1]
            num_total = num * num_cnt + total[num - 1] + total[num + 1]

            cnt[num] += num_cnt
            total[num] += num_total
            result += num_total
        return result % (10**9 + 7)

    def test(self):
        for nums, expected in [
            ([1, 2, 1], 14),
            ([3, 4, 5], 40),
        ]:
            output = self.sumOfGoodSubsequences(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
