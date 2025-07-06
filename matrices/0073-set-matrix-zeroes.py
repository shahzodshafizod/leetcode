import unittest
from typing import List

# https://leetcode.com/problems/set-matrix-zeroes/

# python3 -m unittest matrices/0073-set-matrix-zeroes.py


class Solution(unittest.TestCase):
    # # Approach #1
    # # Time: O(mn)
    # # Space: O(m+n)
    # def setZeroes(self, matrix: List[List[int]]) -> None:
    #     """
    #     Do not return anything, modify matrix in-place instead.
    #     """
    #     m, n = len(matrix), len(matrix[0])
    #     rows, cols = set(), set()
    #     for row in range(m):
    #         for col in range(n):
    #             if matrix[row][col] == 0:
    #                 rows.add(row)
    #                 cols.add(col)
    #     for row in rows:
    #         for col in range(n):
    #             matrix[row][col] = 0
    #     for col in cols:
    #         for row in range(m):
    #             matrix[row][col] = 0

    # # Approach #2
    # # Time: O(mn)
    # # Space: O(n)
    # def setZeroes(self, matrix: List[List[int]]) -> None:
    #     """
    #     Do not return anything, modify matrix in-place instead.
    #     """
    #     m, n = len(matrix), len(matrix[0])
    #     cols = set()
    #     for row in range(m):
    #         first_row = False
    #         for col in range(n):
    #             if matrix[row][col] == 0:
    #                 first_row = True
    #                 cols.add(col)
    #         if first_row:
    #             for col in range(n):
    #                 matrix[row][col] = 0
    #     for col in cols:
    #         for row in range(m):
    #             matrix[row][col] = 0

    # Approach #3
    # Time: O(mn)
    # Space: O(1)
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        m, n = len(matrix), len(matrix[0])
        first_col = False
        for row in range(m):
            if matrix[row][0] == 0:
                first_col = True
            for col in range(1, n):
                if matrix[row][col] == 0:
                    matrix[row][0] = 0
                    matrix[0][col] = 0
        for row in range(m - 1, -1, -1):
            for col in range(n - 1, 0, -1):
                if matrix[row][0] == 0 or matrix[0][col] == 0:
                    matrix[row][col] = 0
            if first_col:
                matrix[row][0] = 0

    def test(self):
        for matrix, expected in [
            ([[1, 1, 1], [1, 0, 1], [1, 1, 1]], [[1, 0, 1], [0, 0, 0], [1, 0, 1]]),
            ([[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]], [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]),
            (
                [[1, 2, 3, 4], [5, 0, 7, 8], [0, 10, 11, 12], [13, 14, 15, 0]],
                [[0, 0, 3, 0], [0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]],
            ),
            (
                [[-4, -2147483648, 6, -7, 0], [-8, 6, -8, -6, 0], [2147483647, 2, -9, -6, -10]],
                [[0, 0, 0, 0, 0], [0, 0, 0, 0, 0], [2147483647, 2, -9, -6, 0]],
            ),
        ]:
            self.setZeroes(matrix)
            output = matrix
            self.assertListEqual(output, expected, f"output: {output}, expected: {expected}")
