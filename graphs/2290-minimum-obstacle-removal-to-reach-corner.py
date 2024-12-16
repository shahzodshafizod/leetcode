import heapq
from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/

# python3 -m unittest graphs/2290-minimum-obstacle-removal-to-reach-corner.py

class Solution(unittest.TestCase):
    # # Approach #1: Depth-First Search + Backtracking (Time Limit Exceeded)
    # # Time: O(4^(m*n)), m=# of rows, n=# of cols
    # # Space: O(m*n)
    # def minimumObstacles(self, grid: List[List[int]]) -> int:
    #     MAX = int(1e5)
    #     directions = [-1,0,1,0,-1]
    #     m, n = len(grid), len(grid[0])
    #     seen = [[False] * n for _ in range(m)]
    #     def dfs(row: int, col: int) -> int:
    #         if row == m-1 and col == n-1:
    #             return 0
    #         result = MAX
    #         for dir in range(1, 5):
    #             r, c = row+directions[dir-1], col+directions[dir]
    #             if min(r, c) >= 0 and r < m and c < n and not seen[r][c]:
    #                 seen[row][col] = True
    #                 result = min(result, grid[row][col] + dfs(r, c))
    #                 seen[row][col] = False
    #         return result
    #     return dfs(0, 0)

    # # Approach #2: Dijkstra's
    # # Time: O(mnlog(mn)), m=# of rows, n=# of cols
    # # Space: O(mn)
    # def minimumObstacles(self, grid: List[List[int]]) -> int:
    #     m, n = len(grid), len(grid[0])
    #     directions = [-1,0,1,0,-1]
    #     min_heap = [(0, 0, 0)] # (obstacles, row, col)
    #     seen = [[False] * n for _ in range(m)]
    #     while min_heap:
    #         obstacles, row, col = heapq.heappop(min_heap)
    #         if (row, col) == (m-1, n-1):
    #             return obstacles
    #         for dir in range(1, 5):
    #             r, c = row+directions[dir-1], col+directions[dir]
    #             if min(r, c) >= 0 and r < m and c < n and not seen[r][c]:
    #                 heapq.heappush(min_heap, (obstacles+grid[r][c], r, c))
    #                 seen[r][c] = True
    #     return m+n-1

    # Approach #3: Greedy (Breadth-First Search + Deque)
    # Time: O(mn), m=# of rows, n=# of cols
    # Space: O(mn)
    def minimumObstacles(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        directions = [-1,0,1,0,-1]
        queue = deque([(0, 0, 0)]) # (obstacles, row, col)
        seen = [[False] * n for _ in range(m)]
        seen[0][0] = True
        while queue:
            obstacles, row, col = queue.popleft()
            if (row, col) == (m-1, n-1):
                return obstacles
            for dir in range(1, 5):
                r, c = row+directions[dir-1], col+directions[dir]
                if min(r, c) >= 0 and r < m and c < n and not seen[r][c]:
                    if grid[r][c]:
                        queue.append((obstacles+1, r, c))
                    else:
                        queue.appendleft((obstacles, r, c))
                    seen[r][c] = True
        return m+n-1

    def test(self):
        for grid, expected in [
            ([[0,0],[0,0]], 0),
            ([[0,1],[1,0]], 1),
            ([[0,0],[0,1],[1,1],[1,0]], 2),
            ([[0,1,1],[1,1,0],[1,1,0]], 2),
            ([[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], 0),
            ([[0,1,1,1,1,0],[1,1,0,0,1,0],[0,0,1,1,0,0],[0,1,1,1,0,1],[1,0,0,1,1,1],[1,1,0,1,0,0],[1,1,0,0,1,0]], 3),
            ([[0,1,1,1,1,1,0,1],[0,1,1,1,1,1,0,1],[1,1,1,1,0,0,0,1],[1,1,1,1,1,1,1,1],[0,1,1,0,0,0,0,1],[1,1,1,1,1,1,0,0]], 4),
            ([[0,0,1,1,1,1,0,0,0,1],[0,1,1,1,1,1,1,0,1,1],[1,1,0,1,1,1,1,0,1,0],[0,0,1,1,1,1,0,0,1,1],[1,0,1,0,0,0,1,1,1,0]], 5),
            # ([[0],[1],[1],[1],[1],[1],[0],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[0]], 454),
        ]:
            output = self.minimumObstacles(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
