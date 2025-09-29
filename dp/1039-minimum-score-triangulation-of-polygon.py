# from functools import lru_cache
from typing import List
import unittest

# https://leetcode.com/problems/minimum-score-triangulation-of-polygon/

# python3 -m unittest dp/1039-minimum-score-triangulation-of-polygon.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(nnn)
    # # Space: O(nn)
    # def minScoreTriangulation(self, values: List[int]) -> int:
    #     @lru_cache(None)
    #     def dp(i: int, j: int) -> int:
    #         if i+2 > j: return 0
    #         if i+2 == j: return values[i] * values[i+1] * values[i+2]
    #         return min(dp(i, k) + values[i] * values[k] * values[j] + dp(k, j) for k in range(i+1, j))
    #     return dp(0, len(values)-1)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(nnn)
    # Space: O(nn)
    def minScoreTriangulation(self, values: List[int]) -> int:
        n = len(values)
        dp = [[0] * n for _ in range(n)]
        for j in range(2, n):
            for i in range(j-2, -1, -1):
                dp[i][j] = (1 << 32) - 1
                for k in range(i+1, j):
                    dp[i][j] = min(dp[i][j], dp[i][k] + values[i] * values[k] * values[j] + dp[k][j])
        return dp[0][n-1]

    def test(self):
        for values, expected in [
            ([1,2,3], 6),
            ([3,7,4,5], 144),
            ([1,3,1,4,1,5], 13),
            ([4,3,4,3,5], 132),
            ([2,1,4,4], 24),
            ([1,2,3,4], 18),
            # ([35,73,90,27,71,80,21,33,33,13,48,12,68,70,80,36,66,3,70,58], 140295),
        ]:
            output = self.minScoreTriangulation(values)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
