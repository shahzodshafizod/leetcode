from typing import List
import unittest

# https://leetcode.com/problems/dungeon-game/

# python3 -m unittest dp/0174-dungeon-game.py

class Solution(unittest.TestCase):
    # # Approach: Dynamic Programming (Top-Down)
    # # Time: O(MxN)
    # # Space: O(MxN)
    # def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
    #     m, n = len(dungeon), len(dungeon[0])
    #     dp = [[0] * n for _ in range(m)]
    #     def dfs(row: int, col: int) -> int:
    #         if row == m or col == n:
    #             return 1e9
    #         if row == m-1 and col == n-1:
    #             return max(1, 1 - dungeon[row][col])
    #         if dp[row][col] != 0:
    #             return dp[row][col]
    #         right = dfs(row, col+1)
    #         down = dfs(row+1, col)
    #         dp[row][col] = max(1, min(right, down) - dungeon[row][col])
    #         return dp[row][col]
    #     return dfs(0,0)

    # Approach: Dynamic Programming (Bottom-Up)
    # Time: O(MxN)
    # Space: O(MxN)
    def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
        m, n = len(dungeon), len(dungeon[0])
        dp = [[float("inf")] * (n+1) for _ in range(m+1)]
        dp[m-1][n] = dp[m][n-1] = 1
        for row in range(m-1,-1,-1):
            for col in range(n-1,-1,-1):
                dp[row][col] = max(1, min(dp[row][col+1], dp[row+1][col]) - dungeon[row][col])
        return dp[0][0]

    def test(self) -> None:
        for dungeon, expected in [
            ([[-2,-3,3],[-5,-10,1],[10,30,-5]], 7),
		    ([[0]], 1),
            ([[0,0]], 1),
            ([[0,100,-98]], 1),
            ([[100]], 1),
            ([[10]], 1),
            ([[-3,5]], 4),
            ([[-10]], 11),
            ([[-2,-3,3,-5,-10]], 18),
            ([[2,3,3,5,10]], 1),
        ]:
            output = self.calculateMinimumHP(dungeon)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
