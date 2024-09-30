from typing import List
import unittest

# https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/

# python3 -m unittest graphs/2328-number-of-increasing-paths-in-a-grid.py

class Solution(unittest.TestCase):
    def countPaths(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dirs = [-1,0,1,0,-1]
        dp = [[0] * n for _ in range(m)]
        def dfs(row: int, col: int) -> int:
            if dp[row][col] != 0:
                return dp[row][col]
            count = 1
            for d in range(4):
                r, c = row+dirs[d], col+dirs[d+1]
                if min(r,c) >= 0 and r<m and c<n and grid[row][col] < grid[r][c]:
                    count += dfs(r, c)
            dp[row][col] = count
            return count
        return sum([dfs(row, col) for row in range(m) for col in range(n)]) % int(1e9+7)

    def test(self) -> None:
        for grid, expected in [
            ([[1,1],[3,4]], 8),
		    ([[1],[2]], 3),
        ]:
            output = self.countPaths(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
