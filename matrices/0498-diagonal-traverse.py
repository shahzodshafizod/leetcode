from typing import List
import unittest

# https://leetcode.com/problems/diagonal-traverse/

# python3 -m unittest matrices/0498-diagonal-traverse.py


class Solution(unittest.TestCase):
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        m, n = len(mat), len(mat[0])
        diag: List[List[int]] = [[] for _ in range(m * n)]
        for row in range(m):
            for col in range(n):
                diag[row + col].append(mat[row][col])
        result: List[int] = []
        for idx, line in enumerate(diag):
            if idx & 1:
                result.extend(line)
            else:
                result.extend(line[::-1])
        return result

    def test(self):
        for mat, expected in [
            ([[1, 2, 3], [4, 5, 6], [7, 8, 9]], [1, 2, 4, 7, 5, 3, 6, 8, 9]),
            ([[1, 2], [3, 4]], [1, 2, 3, 4]),
        ]:
            output = self.findDiagonalOrder(mat)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
