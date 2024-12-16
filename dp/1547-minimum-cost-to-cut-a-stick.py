from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-to-cut-a-stick/

# python3 -m unittest dp/1547-minimum-cost-to-cut-a-stick.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(mmm), m=len(cuts)
    # # Space: O(mm)
    # def minCost(self, n: int, cuts: List[int]) -> int:
    #     cuts = [0] + cuts + [n]
    #     cuts.sort()
    #     m = len(cuts)
    #     memo = [[None] * m for _ in range(m)]
    #     def dp(start: int, end: int) -> int:
    #         if end-start == 1: return 0
    #         if memo[start][end] != None:
    #             return memo[start][end]
    #         best, cost = float("inf"), cuts[end] - cuts[start]
    #         for mid in range(start+1, end):
    #             best = min(best, cost + dp(start,mid) + dp(mid,end))
    #         memo[start][end] = best
    #         return best
    #     return dp(0, m-1)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(mmm), m=len(cuts)
    # Space: O(mm)
    def minCost(self, n: int, cuts: List[int]) -> int:
        cuts.sort()
        cuts = [0] + cuts + [n]
        m = len(cuts)
        dp = [[0] * m for _ in range(m)]
        for diff in range(2, m):
            for start in range(m - diff):
                end = start + diff
                dp[start][end] = float("inf")
                cost = cuts[end]-cuts[start]
                for mid in range(start+1, end):
                    dp[start][end] = min(dp[start][end],
                        cost + dp[start][mid] + dp[mid][end]
                    )
        return dp[0][m-1]
    
    def test(self):
        for n, cuts, expected in [
            (7, [1,3,4,5], 16),
            (9, [5,6,1,4,2], 22),
        ]:
            output = self.minCost(n, cuts)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
