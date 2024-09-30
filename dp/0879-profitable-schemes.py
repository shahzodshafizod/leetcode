from functools import lru_cache
from typing import List
import unittest

# https://leetcode.com/problems/profitable-schemes/

# python3 -m unittest dp/0879-profitable-schemes.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(NGP), G=len(group), P=minProfit
    # # Space: O(NGP)
    # def profitableSchemes(self, n: int, minProfit: int, group: List[int], profit: List[int]) -> int:
    #     mod = int(1e9+7)
    #     g = len(group)
    #     @lru_cache(None)
    #     def dfs(idx: int, members: int, profits: int) -> int:
    #         if idx == g:
    #             return 1 if profits == minProfit else 0
    #         # decision NOT to include
    #         count = dfs(idx+1, members, profits)
    #         # decision to include
    #         if members+group[idx] <= n:
    #             count = (count + dfs(
    #                 idx+1,
    #                 members+group[idx],
    #                 min(minProfit, profits+profit[idx])
    #             )) % mod
    #         return count
    #     return dfs(0, 0, 0)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(NGP), G=len(group), P=minProfit
    # Space: O(NP)
    def profitableSchemes(self, n: int, minProfit: int, group: List[int], profit: List[int]) -> int:
        mod = int(1e9+7)
        dp = [[0] * (n+1) for _ in range(minProfit+1)]
        dp[0][0] = 1
        for p, g in zip(profit, group):
            for i in range(minProfit, -1, -1):
                for j in range(n-g, -1, -1):
                    dp[min(minProfit, i+p)][j+g] += dp[i][j]
                    dp[min(minProfit, i+p)][j+g] %= mod
        return sum(dp[minProfit]) % mod

    def test(self) -> None:
        for n, minProfit, group, profit, expected in [
            (5, 3, [2,2], [2,3], 2),
            (10, 5, [2,3,5], [6,7,8], 7),
            (64, 0, [80, 40], [88, 88], 2),
            # (100, 100, [24,23,7,4,26,3,7,11,1,7,1,3,5,26,26,1,13,12,2,1,7,4,1,27,13,16,26,18,6,1,1,7,16,1,6,2,5,9,19,28,1,23,2,1,3,4,4,3,22,1,1,3,5,34,2,1,22,16,8,5,3,21,1,8,14,2,1,3,8,12,40,6,4,2,2,14,1,11,9,1,7,1,1,1,6,6,4,1,1,7,8,10,20,2,14,31,1,13,1,9], [5,2,38,25,4,17,5,1,4,0,0,8,13,0,20,0,28,1,22,7,10,32,6,37,0,11,6,11,23,20,13,13,6,2,36,1,0,9,4,5,6,14,20,1,13,6,33,0,22,1,17,12,10,1,19,13,8,1,0,17,20,9,8,6,2,2,1,4,22,11,3,2,6,0,40,0,0,7,1,0,25,5,12,7,19,4,12,7,4,4,1,15,33,14,2,1,1,61,4,5], 653387801),
        ]:
            output = self.profitableSchemes(n, minProfit, group, profit)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
