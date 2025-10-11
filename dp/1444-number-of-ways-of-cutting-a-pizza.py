# from functools import lru_cache
from itertools import product
from typing import List
import unittest

# https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/

# python3 -m unittest dp/1444-number-of-ways-of-cutting-a-pizza.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming + Prefix Sum
    # # Time: O(mnk*(m+n)), m=len(pizza), n=len(pizza[0])
    # # Space: O(mnk)
    # def ways(self, pizza: List[str], k: int) -> int:
    #     MOD = 10**9 + 7
    #     rows, cols = len(pizza), len(pizza[0])
    #     presum = [[0] * (cols + 1) for _ in range(rows + 1)]
    #     for i, j in product(reversed(range(rows)), reversed(range(cols))):
    #         presum[i][j] = 1 if pizza[i][j] == 'A' else 0
    #         presum[i][j] += presum[i + 1][j] + presum[i][j + 1] - presum[i + 1][j + 1]

    #     lru_cache(None)

    #     def dp(row: int, col: int, k: int) -> int:
    #         if presum[row][col] == 0:
    #             return 0
    #         if k == 0:
    #             return 1
    #         cnt = 0
    #         # cutting horizontally
    #         for nrow in range(row + 1, rows):
    #             if presum[row][col] > presum[nrow][col]:
    #                 cnt = (cnt + dp(nrow, col, k - 1)) % MOD
    #         # cutting vertically
    #         for ncol in range(col + 1, cols):
    #             if presum[row][col] > presum[row][ncol]:
    #                 cnt = (cnt + dp(row, ncol, k - 1)) % MOD
    #         return cnt

    #     return dp(0, 0, k - 1)

    # Approach #2: Bottom-Up Dynamic Programming + Prefix Sum
    # Time: O(mnk*(m+n)), m=len(pizza), n=len(pizza[0])
    # Space: O(mnk)
    def ways(self, pizza: List[str], k: int) -> int:
        MOD = 10**9 + 7
        rows, cols = len(pizza), len(pizza[0])
        presum = [[0] * (cols + 1) for _ in range(rows + 1)]
        for r, c in product(reversed(range(rows)), reversed(range(cols))):
            presum[r][c] = 1 if pizza[r][c] == 'A' else 0
            presum[r][c] += presum[r + 1][c] + presum[r][c + 1] - presum[r + 1][c + 1]

        dp = [[[0] * k for _ in range(cols + 1)] for _ in range(rows + 1)]
        for r, c in product(range(rows), range(cols)):
            dp[r][c][0] = 1 if presum[r][c] > 0 else 0

        for cuts, r, c in product(range(1, k), range(rows), range(cols)):
            for nr in range(r + 1, rows):
                if presum[r][c] > presum[nr][c]:
                    dp[r][c][cuts] = (dp[r][c][cuts] + dp[nr][c][cuts - 1]) % MOD
            for nc in range(c + 1, cols):
                if presum[r][c] > presum[r][nc]:
                    dp[r][c][cuts] = (dp[r][c][cuts] + dp[r][nc][cuts - 1]) % MOD

        return dp[0][0][k - 1]

    def test(self):
        for pizza, k, expected in [
            (["A..", "AAA", "..."], 3, 3),
            (["A..", "AA.", "..."], 3, 1),
            (["A..", "A..", "..."], 1, 1),
            (["...", "...", "..."], 1, 0),
        ]:
            output = self.ways(pizza, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
