from functools import lru_cache
import unittest

# https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/

# python3 -m unittest dp/1866-number-of-ways-to-rearrange-sticks-with-k-sticks-visible.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(NxK)
    # # Space: O(NxK)
    # def rearrangeSticks(self, n: int, k: int) -> int:
    #     @lru_cache(None)
    #     def dp(n: int, k: int) -> int:
    #         if n == k:
    #             return 1
    #         if k == 0:
    #             return 0
    #         # 1. decision to include
    #         # if we put largest (n) in the end,
    #         # it'll be visible, so we decrease k
    #         count = dp(n-1, k-1)
    #         # 2. decision NOT to include
    #         # non-largest will be hidden by the largest
    #         count += (n-1) * dp(n-1, k)
    #         return count % mod
    #     mod = int(1e9+7)
    #     return dp(n, k)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(NxK)
    # Space: O(K)
    def rearrangeSticks(self, n: int, k: int) -> int:
        mod = int(1e9+7)
        dp = [0] * (k+1)
        dp[0] = 1
        for x in range(1, n+1):
            tmp = [0] * (k+1)
            for y in range(1, k+1):
                tmp[y] = (dp[y-1] + (x-1)*dp[y]) % mod
            dp = tmp
        return dp[k]

    def test(self):
        for n, k, expected in [
            (3, 2, 3),
            (5, 5, 1),
            (20, 11, 647427950),
            (978, 575, 103376920),
        ]:
            output = self.rearrangeSticks(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
