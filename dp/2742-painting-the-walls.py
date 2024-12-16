from typing import List
import unittest

# https://leetcode.com/problems/painting-the-walls/

# python3 -m unittest dp/2742-painting-the-walls.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def paintWalls(self, cost: List[int], time: List[int]) -> int:
    #     n = len(cost)
    #     memo = [[None] * (n+1) for _ in range(n)]
    #     def dp(idx: int, walls: int) -> int:
    #         if walls <= 0: return 0
    #         if idx == n: return float("inf")
    #         if memo[idx][walls] != None: return memo[idx][walls]
    #         memo[idx][walls] = min(dp(idx+1, walls),
    #             cost[idx] + dp(idx+1, walls-1-time[idx]))
    #         return memo[idx][walls]
    #     return dp(0, n)

    # # Approach: Bottom-Up Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def paintWalls(self, cost: List[int], time: List[int]) -> int:
    #     n = len(cost)
    #     dp = [[0] * (n+1) for _ in range(n+1)]
    #     for wall in range(1, n+1):
    #         dp[n][wall] = int(1e9)
    #     for idx in range(n-1,-1,-1):
    #         for wall in range(1,n+1):
    #             dp[idx][wall] = min(
    #                 cost[idx] + dp[idx+1][max(0, wall-1-time[idx])],
    #                 dp[idx+1][wall],
    #             )
    #     return dp[0][n]

    # Approach: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(nn)
    # Space: O(n)
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        n = len(cost)
        curr, next = [int(1e9)]*(n+1), [0]*(n+1)
        curr[0] = 0
        for idx in range(n-1,-1,-1):
            curr, next = next, curr
            for wall in range(1,n+1):
                curr[wall] = min(
                    # while paid painter paints 1 wall
                    # the free one paints time[idx] walls
                    cost[idx] + next[max(0, wall-1-time[idx])],
                    next[wall],
                )
        return curr[n]

    def test(self):
        for cost, time, expected in [
            ([2,2], [1,1], 2),
            ([1,2,3,2], [1,2,3,2], 3),
            ([2,3,4,2], [1,1,1,1], 4),
            ([8,7,5,15], [1,1,2,1], 12),
            ([49,35,32,20,30,12,42], [1,1,2,2,1,1,2], 62),
            ([42,8,28,35,21,13,21,35], [2,1,1,1,2,1,1,2], 63),
            ([26,53,10,24,25,20,63,51], [1,1,1,1,2,2,2,1], 55),
            ([76,25,96,46,85,19,29,88,2,5], [1,2,1,3,1,3,3,3,2,1], 46),
        ]:
            output = self.paintWalls(cost, time)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
