import unittest
import sys
sys.setrecursionlimit(1000000)

# https://leetcode.com/problems/count-ways-to-build-good-strings/

# python3 -m unittest dp/2466-count-ways-to-build-good-strings.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(high)
    # # Space: O(high)
    # def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
    #     MOD = int(1e9 + 7)
    #     memo = [None] * (high+1)
    #     def dp(length: int) -> int:
    #         if memo[length] != None:
    #             return memo[length]
    #         count = 1 if low <= length <= high else 0
    #         if length + zero <= high:
    #             count = (count + dp(length + zero)) % MOD
    #         if length + one <= high:
    #             count = (count + dp(length + one)) % MOD
    #         memo[length] = count
    #         return count
    #     return dp(0)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(high)
    # Space: O(high)
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = int(1e9 + 7)
        dp = [1] + [0] * high
        for length in range(1, high+1):
            if length - zero >= 0:
                dp[length] = dp[length-zero]
            if length - one >= 0:
                dp[length] = (dp[length] + dp[length-one]) % MOD
        return sum(dp[low:high+1]) % MOD

    def test(self):
        for low, high, zero, one, expected in [
            (3, 3, 1, 1, 8),
            (2, 3, 1, 2, 5),
            (1, 1, 1, 1, 2),
            (100000, 100000, 100000, 100000, 2),
            # (200, 200, 10, 1, 764262396),
            # (1, 100000, 1, 1, 215447031),
            # (200, 200, 10, 1, 764262396),
            # (277, 99991, 31, 19, 271302617),
            # (2657, 9601, 1, 2, 267484946),
            # (2, 8699, 1, 1, 987490208),
        ]:
            output = self.countGoodStrings(low, high, zero, one)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
