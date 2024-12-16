from typing import List
import unittest

# https://leetcode.com/problems/cherry-pickup-ii/

# python3 -m unittest dp/1463-cherry-pickup-ii.py

class Solution(unittest.TestCase):
    # # Approach: Dynamic Programming (Top-Down)
    # # Time: O(m*n^2)
    # # Space: O(m*n^2)
    # def cherryPickup(self, grid: List[List[int]]) -> int:
    #     m, n = len(grid), len(grid[0])
    #     memo = [[[None] * n for _ in range(n)] for _ in range(m)]
    #     def dfs(row: int, col1: int, col2: int) -> int:
    #         # row1 & row2 are equal
    #         if memo[row][col1][col2] != None:
    #             return memo[row][col1][col2]
    #         if row == m-1:
    #             return grid[row][col1] + grid[row][col2]
    #         count = 0
    #         for nc1 in range(max(0, col1-1), min(n-1, col1+1)+1):
    #             for nc2 in range(max(0, nc1+1, col2-1), min(n-1, col2+1)+1):
    #                 count = max(count, dfs(row+1, nc1, nc2))
    #         count += grid[row][col1] + grid[row][col2]
    #         memo[row][col1][col2] = count
    #         return count
    #     return dfs(0, 0, n-1)

    # Approach: Dynamic Programming (Bottom-Up)
    # Time: O(m*n^2)
    # Space: O(n^2)
    def cherryPickup(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[[0] * n for _ in range(n)] for _ in range(2)]
        next, curr = 0, 1
        for row in range(m-1, -1, -1):
            next, curr = curr, next
            for col1 in range(n):
                for col2 in range(col1+1, n):
                    for nc1 in range(max(0, col1-1), min(n-1, col1+1)+1):
                        for nc2 in range(max(0, nc1+1, col2-1), min(n-1, col2+1)+1):
                            dp[curr][col1][col2] = max( dp[curr][col1][col2], dp[next][nc1][nc2] )
                    dp[curr][col1][col2] += grid[row][col1] + grid[row][col2]
        return dp[curr][0][n-1]

    def test(self):
        for grid, expected in [
            ([[3,1,1],[2,5,1],[1,5,5],[2,1,1]], 24),
            ([[13,14,37,49,64,98,4,11,47,81],[71,46,50,50,10,14,35,35,52,69]], 234),
		    ([[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]], 28),
            ([[9,79],[49,85],[54,48],[37,72],[68,14],[53,30],[65,80],[94,18],[89,46],[7,93]], 1090),
            ([[1,0,0,0,0,0,0,0,0,2],[0,6,0,0,0,0,0,0,6,0],[0,0,9,0,0,0,0,3,0,0],[0,0,0,8,0,0,5,0,0,0],[100,0,0,0,2,3,0,0,0,100]], 227),
            ([[4,0,0,0,0,0,0,0,0,86],[0,52,0,0,0,0,0,0,48,0],[0,0,94,0,0,0,0,74,0,0],[0,0,0,98,0,0,25,0,0,0],[0,0,0,0,70,66,0,0,0,0]], 617),
            # ([[0,40,0,0,0,0,0,0,93,0],[0,0,41,0,0,0,0,59,0,0],[89,0,0,0,0,0,0,0,0,28],[0,32,0,0,0,0,0,0,80,0],[0,0,40,0,0,0,0,0,0,0],[50,0,0,0,0,0,0,0,0,66],[0,90,0,0,0,0,0,0,11,0],[0,0,62,0,0,0,0,12,0,0],[95,0,0,0,0,0,0,0,0,88],[0,31,0,0,0,0,0,0,26,0]], 686),
            # ([[0,100,100,100,100,100,100,100,100,0],[0,0,100,100,100,100,100,100,0,0],[0,0,0,100,100,100,100,0,0,0],[0,0,0,0,100,100,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0]], 0),
            # ([[63,56,38,100,40,50,44,98,27,20],[13,98,92,31,46,29,81,2,37,3],[75,5,46,15,74,17,34,60,100,44],[8,4,17,14,60,77,23,0,93,83],[41,40,5,2,73,80,71,100,82,39],[96,76,56,42,65,65,22,11,85,80],[64,71,27,78,85,15,2,28,75,31],[51,16,30,65,54,68,12,5,48,1],[100,57,93,43,40,51,3,82,46,91],[81,63,20,12,83,59,59,46,67,66]], 1400),
        ]:
            output = self.cherryPickup(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
