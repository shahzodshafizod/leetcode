import unittest

# https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/

# python3 -m unittest dp/2787-ways-to-express-an-integer-as-sum-of-powers.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def numberOfWays(self, n: int, x: int) -> int:
    #     MOD = 10**9 + 7
    #     memo = [[-1] * (n + 1) for _ in range(n + 1)]

    #     def dp(k: int, n: int) -> int:
    #         if n <= 0:
    #             return 1 if n == 0 else 0
    #         if k == 0:
    #             return 0
    #         if memo[k][n] != -1:
    #             return memo[k][n]
    #         ways = dp(k - 1, n)
    #         if k**x <= n:
    #             ways = (ways + dp(k - 1, n - k**x)) % MOD
    #         memo[k][n] = ways
    #         return ways

    #     return dp(n, n)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def numberOfWays(self, n: int, x: int) -> int:
    #     MOD = 10**9 + 7
    #     dp = [[0] * (n + 1) for _ in range(n + 1)]
    #     dp[0][0] = 1
    #     max_i = 0
    #     for i in range(1, n + 1):
    #         power = i**x
    #         if power > n:
    #             break
    #         max_i = i
    #         for j in range(n + 1):
    #             dp[i][j] = dp[i - 1][j]
    #             if j >= power:
    #                 dp[i][j] = (dp[i][j] + dp[i - 1][j - power]) % MOD
    #     return dp[max_i][n]

    # Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(nn)
    # Space: O(n)
    def numberOfWays(self, n: int, x: int) -> int:
        MOD = 10**9 + 7
        dp = [0] * (n + 1)
        dp[0] = 1
        for i in range(1, n + 1):
            power = i**x
            for j in range(n, power - 1, -1):
                dp[j] = (dp[j] + dp[j - power]) % MOD
        return dp[n]

    def test(self):
        for n, x, expected in [
            (10, 2, 1),
            (4, 1, 2),
        ]:
            output = self.numberOfWays(n, x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
