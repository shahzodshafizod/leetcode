from typing import List
import unittest

# https://leetcode.com/problems/making-a-large-island/

# python3 -m unittest graphs/0827-making-a-large-island.py

class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(nn)
    # Space: O(nn)
    def largestIsland(self, grid: List[List[int]]) -> int:
        n = len(grid)
        ids = [[-1] * n for _ in range(n)]
        directions = [1, 0, -1, 0, 1]
        def dfs(row: int, col: int, id: int) -> int:
            ids[row][col] = id
            area = 1
            for dir in range(1, 5):
                r = row + directions[dir-1]
                c = col + directions[dir]
                if min(r, c) < 0 or max(r, c) == n or \
                    not grid[r][c] or ids[r][c] != -1:
                    continue
                area += dfs(r, c, id)
            return area
        id = 0
        areas, zeroes = [], []
        max_area = 0
        for row in range(n):
            for col in range(n):
                if grid[row][col] == 0:
                    zeroes.append((row, col))
                elif ids[row][col] == -1:
                    areas.append(dfs(row, col, id))
                    max_area = max(max_area, areas[-1])
                    id += 1
        if len(areas) == 0:
            return 1
        for row, col in zeroes:
            neighbors = set()
            for dir in range(1, 5):
                r = row + directions[dir-1]
                c = col + directions[dir]
                if min(r, c) >= 0 and max(r, c) < n and grid[r][c]:
                    neighbors.add(ids[r][c])
            max_area = max(max_area, 1 + sum(areas[nei] for nei in neighbors))
        return max_area

    def test(self):
        for grid, expected in [
            ([[1,0],[0,1]], 3),
            ([[1,1],[1,0]], 4),
            ([[1,1],[1,1]], 4),
            ([[0,0],[0,0]], 1),
            ([[1,1,0,1],[1,0,0,1],[1,0,0,1],[1,0,0,1]], 10),
            ([[1,0,1,0],[0,0,0,1],[1,0,1,0],[1,0,1,0]], 5),
            ([[0,0,0,1,1],[1,0,0,1,0],[1,1,1,1,1],[1,1,1,0,1],[0,0,0,1,0]], 15),
        ]:
            output = self.largestIsland(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
