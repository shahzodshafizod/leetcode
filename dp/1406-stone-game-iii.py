from typing import List
import unittest

# https://leetcode.com/problems/stone-game-iii/

# python3 -m unittest dp/1406-stone-game-iii.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def stoneGameIII(self, stoneValue: List[int]) -> str:
    #     n = len(stoneValue)
    #     cache = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx == n: return 0
    #         if cache[idx] != None: return cache[idx]
    #         total, stones = float("-inf"), 0
    #         for x in range(idx, min(n, idx+3)):
    #             stones += stoneValue[x]
    #             total = max(total, stones - dp(x+1)) # 1-(2-3)=1-2+3
    #         cache[idx] = total
    #         return total
    #     return "Alice" if dp(0) > 0 else ("Bob" if dp(0) < 0 else "Tie")

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def stoneGameIII(self, stoneValue: List[int]) -> str:
    #     n = len(stoneValue)
    #     dp = [0]*(n+1)
    #     for idx in range(n-1,-1,-1):
    #         dp[idx] = float("-inf")
    #         stones = 0
    #         for x in range(idx, min(n, idx+3)):
    #             stones += stoneValue[x]
    #             dp[idx] = max(dp[idx], stones - dp[x+1])
    #     return "Alice" if dp[0] > 0 else ("Bob" if dp[0] < 0 else "Tie")

    # Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(n)
    # Space: O(1)
    def stoneGameIII(self, stoneValue: List[int]) -> str:
        n = len(stoneValue)
        dp = [0] * 3
        for idx in range(n - 1, -1, -1):
            best = float("-inf")
            stones = 0
            for x in range(idx, min(n, idx + 3)):
                stones += stoneValue[x]
                best = max(best, stones - dp[x - idx])
            dp[0], dp[1], dp[2] = best, dp[0], dp[1]
        return "Alice" if dp[0] > 0 else ("Bob" if dp[0] < 0 else "Tie")

    def test(self):
        for stoneValue, expected in [
            ([1, 2, 3, 7], "Bob"),
            ([1, 2, 3, -9], "Alice"),
            ([1, 2, 3, 6], "Tie"),
        ]:
            output = self.stoneGameIII(stoneValue)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
