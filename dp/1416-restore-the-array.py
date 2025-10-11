import unittest

# https://leetcode.com/problems/restore-the-array/

# python3 -m unittest dp/1416-restore-the-array.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n log k)
    # # Space: O(n)
    # def numberOfArrays(self, s: str, k: int) -> int:
    #     MOD = 10**9 + 7
    #     n = len(s)
    #     memo = [-1] * n

    #     def dfs(i: int) -> int:
    #         if i == n:
    #             return 1
    #         if s[i] == '0':
    #             return 0
    #         if memo[i] != -1:
    #             return memo[i]
    #         ans, num = 0, 0
    #         for j in range(i, n):
    #             num = num * 10 + int(s[j])
    #             if num > k:
    #                 break
    #             ans = (ans + dfs(j + 1)) % MOD
    #         memo[i] = ans
    #         return ans

    #     return dfs(0)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n log k)
    # Space: O(n)
    def numberOfArrays(self, s: str, k: int) -> int:
        MOD = 10**9 + 7
        n = len(s)
        dp = [0] * (n + 1)
        dp[n] = 1

        for i in range(n - 1, -1, -1):
            if s[i] == '0':
                continue
            num = 0
            for j in range(i, n):
                num = num * 10 + int(s[j])
                if num > k:
                    break
                dp[i] = (dp[i] + dp[j + 1]) % MOD
        return dp[0]

    def test(self):
        for s, k, expected in [
            ("1000", 10000, 1),
            ("1000", 10, 0),
            ("1317", 2000, 8),
            ("1111111111111", 1000000000, 4076),
        ]:
            output = self.numberOfArrays(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
