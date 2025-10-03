from typing import List
import unittest

# https://leetcode.com/problems/last-day-where-you-can-still-cross/

# python3 -m unittest unionfinds/1970-last-day-where-you-can-still-cross.py


class Solution(unittest.TestCase):
    # Approach #1: Binary Search, DFS
    # Time: O(row * col * log(len(cells)))
    # Space: O(row * col)
    def latestDayToCross(self, row: int, col: int, cells: List[List[int]]) -> int:
        directions = [(0, -1), (-1, 0), (0, 1), (1, 0)]
        def dfs(seen: List[List[bool]], r: int, c: int) -> bool:
            if min(r, c) < 0 or r == row or c == col or seen[r][c]:
                return False
            if r == row-1:
                return True
            seen[r][c] = True
            for dr, dc in directions:
                if dfs(seen, r+dr, c+dc):
                    return True
            return False
        left, right = 0, len(cells)
        while left <= right:
            mid = (left + right) // 2
            seen = [[False] * col for _ in range(row)]
            for i in range(mid):
                r, c = cells[i]
                seen[r-1][c-1] = True
            if any(dfs(seen, 0, c) for c in range(col)):
                left = mid+1
            else:
                right = mid-1
        return right

    # Approach #2: Union Find
    # Time: O(row * col)
    # Space: O(row * col)
    def latestDayToCross(self, row: int, col: int, cells: List[List[int]]) -> int:
        n = row * col + 2
        root = list(range(n))
        def find(x: int) -> int:
            if root[x] != x:
                root[x] = find(root[x])
            return root[x]
        def union(x: int, y: int):
            rx, ry = find(x), find(y)
            if rx < ry:
                root[ry] = rx
            elif ry < rx:
                root[rx] = ry
        directions = [(-1,-1),(-1,0),(-1,1),(0,1),(1,1),(1,0),(1,-1),(0,-1)]
        grid = [[0] * col for _ in range(row)]
        left, right = 0, n-1
        day = 0
        for r, c in cells:
            r, c = r-1, c-1
            grid[r][c] = 1
            curr = r*col+c+1
            for dr, dc in directions:
                nr, nc = r + dr, c + dc
                if 0 <= nr < row and 0 <= nc < col and grid[nr][nc] == 1:
                    union(curr, nr*col+nc+1)
            if c == 0:
                union(curr, left)
            elif c == col-1:
                union(curr, right)
            if find(left) == find(right):
                break
            day += 1
        return day

    def test(self):
        for row, col, cells, expected in [
            (2, 2, [[1,1],[2,1],[1,2],[2,2]], 2),
            (2, 2, [[1,1],[1,2],[2,1],[2,2]], 1),
            (3, 3, [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]], 3),
        ]:
            output = self.latestDayToCross(row, col, cells)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
