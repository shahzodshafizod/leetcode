from typing import List
import unittest

# https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/

# python3 -m unittest matrices/2658-maximum-number-of-fish-in-a-grid.py

class Solution(unittest.TestCase):
    # # Approach: Depth-First Search (Immutable Input)
    # # Time: O(mn)
    # # Space: O(mn)
    # def findMaxFish(self, grid: List[List[int]]) -> int:
    #     directions = [1, 0, -1, 0, 1]
    #     m, n = len(grid), len(grid[0])
    #     seen = [[False] * n for _ in range(m)]
    #     def dfs(row: int, col: int) -> int:
    #         fish = grid[row][col]
    #         seen[row][col] = True
    #         for dir in range(1, 5):
    #             r = row + directions[dir-1]
    #             c = col + directions[dir]
    #             if min(r, c) < 0 or r == m or c == n or grid[r][c] == 0 or seen[r][c]:
    #                 continue
    #             fish += dfs(r, c)
    #         return fish
    #     maxfish = 0
    #     for row in range(m):
    #         for col in range(n):
    #             if grid[row][col] > 0 and not seen[row][col]:
    #                 maxfish = max(maxfish, dfs(row, col))
    #     return maxfish

    # Approach: Depth-First Search
    # Time: O(mn)
    # Space: O(1)
    def findMaxFish(self, grid: List[List[int]]) -> int:
        directions = [1, 0, -1, 0, 1]
        m, n = len(grid), len(grid[0])
        def dfs(row: int, col: int) -> int:
            fish = grid[row][col]
            grid[row][col] = 0
            for dir in range(1, 5):
                r = row + directions[dir-1]
                c = col + directions[dir]
                if min(r, c) < 0 or r == m or c == n or grid[r][c] == 0:
                    continue
                fish += dfs(r, c)
            return fish
        maxfish = 0
        for row in range(m):
            for col in range(n):
                if grid[row][col] > 0:
                    maxfish = max(maxfish, dfs(row, col))
        return maxfish

    def test(self):
        for grid, expected in [
            ([[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]], 7),
            ([[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]], 1),
            ([[0]], 0),
            ([[4]], 4),
            ([[0,0]], 0),
            ([[0,6,0,9]], 9),
            ([[4,5,5],[0,10,0]], 24),
            ([[3],[4],[0],[4],[8],[9],[0]], 21),
        ]:
            output = self.findMaxFish(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
