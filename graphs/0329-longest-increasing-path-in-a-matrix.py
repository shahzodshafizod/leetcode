from typing import List
import unittest

# https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

# python3 -m unittest graphs/0329-longest-increasing-path-in-a-matrix.py

class Solution(unittest.TestCase):
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        m, n = len(matrix), len(matrix[0])
        dirs = [-1,0,1,0,-1]
        dp = {} # (row, col)
        def dfs(row: int, col: int) -> int:
            if (row, col) in dp:
                return dp[(row, col)]
            length = 1
            for d in range(4):
                r, c = row+dirs[d], col+dirs[d+1]
                if 0<=r<m and 0<=c<n and matrix[r][c] > matrix[row][col]:
                    length = max(length, 1 + dfs(r, c))
            dp[(row, col)] = length
            return length
        length = 1
        for row in range(m):
            for col in range(n):
                length = max(length, dfs(row, col))
        return length

    def test(self) -> None:
        for matrix, expected in [
            ([[9,9,4],[6,6,8],[2,1,1]], 4),
            ([[3,4,5],[3,2,6],[2,2,1]], 4),
            ([[1]], 1),
        ]:
            output = self.longestIncreasingPath(matrix)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
