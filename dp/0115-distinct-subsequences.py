import unittest

# https://leetcode.com/problems/distinct-subsequences/

# python3 -m unittest dp/0115-distinct-subsequences.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(ns * nt), ns=len(s), nt=len(t)
    # # Space: O(ns * nt)
    # def numDistinct(self, s: str, t: str) -> int:
    #     ns, nt = len(s), len(t)
    #     dp = [[None] * nt for _ in range(ns)]
    #     def dfs(i: int, j: int) -> int:
    #         if j == nt: return 1
    #         if i == ns: return 0
    #         if dp[i][j] != None:
    #             return dp[i][j]
    #         if s[i] == t[j]:
    #             dp[i][j] = dfs(i+1, j) + dfs(i+1, j+1)
    #         else:
    #             dp[i][j] = dfs(i+1, j)
    #         return dp[i][j]
    #     return dfs(0, 0)

    # # Approach: Bottom-UP Dynamic Programming
    # # Time: O(ns * nt), ns=len(s), nt=len(t)
    # # Space: O(ns * nt)
    # def numDistinct(self, s: str, t: str) -> int:
    #     ns, nt = len(s), len(t)
    #     dp = [[0] * (nt+1) for _ in range(ns+1)]
    #     for i in range(ns+1): dp[i][0] = 1
    #     for i in range(1, ns+1):
    #         for j in range(1, nt+1):
    #             if s[i-1] == t[j-1]:
    #                 dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
    #             else:
    #                 dp[i][j] = dp[i-1][j]
    #     return dp[ns][nt]

    # Approach: Bottom-UP Dynamic Programming (Space Optimized)
    # Time: O(ns * nt), ns=len(s), nt=len(t)
    # Space: O(nt)
    def numDistinct(self, s: str, t: str) -> int:
        ns, nt = len(s), len(t)
        dp = [0] * (nt+1)
        dp[0] = 1
        for i in range(1, ns+1):
            for j in range(nt, 0, -1):
                if s[i-1] == t[j-1]:
                    dp[j] += dp[j-1]
        return dp[nt]

    def test(self):
        for s, t, expected in [
            ("rabbbit", "rabbit", 3),
            ("babgbag", "bag", 5),
            ("a", "", 1),
            ("", "a", 0),
        ]:
            output = self.numDistinct(s, t)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
