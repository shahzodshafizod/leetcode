from typing import List
import unittest

# https://leetcode.com/problems/new-21-game/

# python3 -m unittest dp/0837-new-21-game.py

# In short, you're tasked with finding the likelihood of her final score being between k and n.


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(k * maxPts)
    # # Space: O(k)
    # def new21Game(self, n: int, k: int, maxPts: int) -> float:
    #     memo: List[float] = [-1] * k

    #     def dp(points: int) -> float:
    #         if points >= k:
    #             return 1 if points <= n else 0
    #         if memo[points] != -1:
    #             return memo[points]
    #         score = 0
    #         for point in range(1, maxPts + 1):
    #             score += dp(points + point)
    #         memo[points] = score / maxPts
    #         return memo[points]

    #     return dp(0)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(n * maxPts)
    # # Space: O(n)
    # def new21Game(self, n: int, k: int, maxPts: int) -> float:
    #     dp: List[float] = [0] * (n + 1)
    #     dp[0] = 1
    #     for i in range(1, n + 1):
    #         for j in range(1, maxPts + 1):
    #             if i - j >= 0 and i - j < k:
    #                 dp[i] += dp[i - j] / maxPts
    #     return sum(dp[k:])

    # Approach #3: Bottom-Up Dynamic Programming (Optimized)
    # Time: O(n)
    # Space: O(n)
    def new21Game(self, n: int, k: int, maxPts: int) -> float:
        dp: List[float] = [0] * (n + 1)
        dp[0] = 1
        s = 1 if k > 0 else 0
        for i in range(1, n + 1):
            dp[i] = s / maxPts
            if i < k:
                s += dp[i]
            if i - maxPts >= 0 and i - maxPts < k:
                s -= dp[i - maxPts]
        return sum(dp[k:])

    def test(self):
        for n, k, maxPts, expected in [
            (10, 1, 10, 1.00000),
            (6, 1, 10, 0.60000),
            (21, 17, 10, 0.73278),
        ]:
            output = self.new21Game(n, k, maxPts)
            output = round(output, 5)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
