import unittest

# https://leetcode.com/problems/find-all-possible-stable-binary-arrays-ii/

# python3 -m unittest dp/3130-find-all-possible-stable-binary-arrays-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (TLE + MLE)
    # # Time: O(zero * one * limit)
    # # Space: O(zero * one * limit)
    # def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
    #     MOD = 10**9 + 7
    #     memo = [[[-1] * 2 for _ in range(one+1)] for _ in range(zero+1)]
    #     def dp(zero: int, one: int, prev: int) -> int:
    #         nonlocal limit
    #         if zero == 0 and one == 0:
    #             return 1
    #         if memo[zero][one][prev] != -1:
    #             return memo[zero][one][prev]
    #         total = 0
    #         if prev == 1:
    #             for i in range(1, min(zero, limit)+1):
    #                 total += dp(zero-i, one, 0)
    #         else:
    #             for i in range(1, min(one, limit)+1):
    #                 total += dp(zero, one-i, 1)
    #         memo[zero][one][prev] = total
    #         return total
    #     return (dp(zero, one, 0)+dp(zero, one, 1))%MOD

    # # Approach #2: Bottom-Up Dynamic Programming (TLE + MLE)
    # # Time: O(zero * one * limit)
    # # Space: O(zero * one * limit)
    # def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
    #     MOD = 10**9 + 7
    #     dp = [[[0] * 2 for _ in range(one+1)] for _ in range(zero+1)]
    #     dp[0][0][0] = 1
    #     dp[0][0][1] = 1
    #     for z in range(zero+1):
    #         for o in range(one+1):
    #             if z == 0 and o == 0: continue
    #             for i in range(1, min(z, limit)+1):
    #                 dp[z][o][1] += dp[z-i][o][0]
    #             for i in range(1, min(o, limit)+1):
    #                 dp[z][o][0] += dp[z][o-i][1]
    #     return sum(dp[zero][one]) % MOD

    # Approach #3: Bottom-Up Dynamic Programming
    # Time: O(zero * one)
    # Space: O(zero * one)
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        MOD = 10**9 + 7
        dp = [[[0] * 2 for _  in range(one+1)] for _ in range(zero+1)]
        for z in range(min(zero, limit)+1):
            dp[z][0][0] = 1
        for o in range(min(one, limit)+1):
            dp[0][o][1] = 1
        for z in range(1, zero+1):
            for o in range(1, one+1):
                dp[z][o][0] = sum(dp[z-1][o])
                dp[z][o][1] = sum(dp[z][o-1])
                if z > limit:
                    dp[z][o][0] = (dp[z][o][0] - dp[z-limit-1][o][1] + MOD) % MOD
                if o > limit:
                    dp[z][o][1] = (dp[z][o][1] - dp[z][o-limit-1][0] + MOD) % MOD
        return sum(dp[zero][one]) % MOD

    def test(self):
        for zero, one, limit, expected in [
            (1, 1, 2, 2),
            (1, 2, 1, 1),
            (3, 3, 2, 14),
        ]:
            output = self.numberOfStableArrays(zero, one, limit)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
