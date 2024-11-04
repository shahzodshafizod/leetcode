from typing import List
import unittest

# https://leetcode.com/problems/stone-game-viii/

# python3 -m unittest dp/1872-stone-game-viii.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
    # # Time: O(nn)
    # # Space: O(n)
    # def stoneGameVIII(self, stones: List[int]) -> int:
    #     n = len(stones)
    #     presum = [0] * n
    #     presum[0] = stones[0]
    #     for idx in range(1,n):
    #         presum[idx] = presum[idx-1] + stones[idx]
    #     memo = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx == n: return 0
    #         if memo[idx] != None: return memo[idx]
    #         memo[idx] = float("-inf")
    #         for x in range(idx, n):
    #             memo[idx] = max(memo[idx], presum[x] - dp(x+1))
    #         return memo[idx]
    #     return dp(1)

    # # Approach #2: Top-Down Dynamic Programming (Time-Optimized)
    # # Time: O(n)
    # # Space: O(n)
    # def stoneGameVIII(self, stones: List[int]) -> int:
    #     n = len(stones)
    #     presum = [0] * n
    #     presum[0] = stones[0]
    #     for idx in range(1,n):
    #         presum[idx] = presum[idx-1] + stones[idx]
    #     memo = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx == n-1: return presum[idx]
    #         if memo[idx] != None: return memo[idx]
    #         memo[idx] = max(dp(idx+1), presum[idx] - dp(idx+1))
    #         return memo[idx]
    #     return dp(1)

    # # Approach #3: Bottom-Up Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def stoneGameVIII(self, stones: List[int]) -> int:
    #     n = len(stones)
    #     presum = [0] * n
    #     presum[0] = stones[0]
    #     for idx in range(1, n):
    #         presum[idx] = presum[idx-1] + stones[idx]
    #     dp = [0] * n
    #     dp[-1] = presum[-1]
    #     for idx in range(n-2,0,-1):
    #         dp[idx] = max(dp[idx+1], presum[idx] - dp[idx+1])
    #     return dp[1]

    # Approach #4: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(n)
    # Space: O(n)
    def stoneGameVIII(self, stones: List[int]) -> int:
        n = len(stones)
        presum = [0] * n
        presum[0] = stones[0]
        for idx in range(1, n):
            presum[idx] = presum[idx-1] + stones[idx]
        dp = presum[-1]
        for idx in range(n-2,0,-1):
            dp = max(dp, presum[idx] - dp)
        return dp

    def test(self) -> None:
        for stones, expected in [
            ([-10,-12], -22),
            ([-1,2,-3,4,-5], 5),
            ([7,-6,5,10,5,-2,-6], 13),
            ([25,-35,-37,4,34,43,16,-33,0,-17,-31,-42,-42,38,12,-5,-43,-10,-37,12], 38),
        ]:
            output = self.stoneGameVIII(stones)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
