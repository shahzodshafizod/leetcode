from typing import List
import unittest

# https://leetcode.com/problems/toeplitz-matrix/

# python3 -m unittest matrices/0766-toeplitz-matrix.py


class Solution(unittest.TestCase):
    # Approach: Compare Each Element with Top-Left Neighbor
    # A Toeplitz matrix has the property that all diagonals from top-left to bottom-right
    # contain the same element. This means for every position (i, j) where i > 0 and j > 0,
    # matrix[i][j] should equal matrix[i-1][j-1].
    #
    # Algorithm:
    # 1. Iterate through the matrix starting from row 1, column 1
    # 2. For each element, check if it equals its top-left neighbor
    # 3. If any element differs from its top-left neighbor, return False
    # 4. If all elements match their top-left neighbors, return True
    #
    # Time: O(m * n) where m is rows, n is columns - visit each element once
    # Space: O(1) - only using constant extra space
    def isToeplitzMatrix(self, matrix: List[List[int]]) -> bool:
        m, n = len(matrix), len(matrix[0])

        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j] != matrix[i - 1][j - 1]:
                    return False

        return True

    def test(self):
        for matrix, expected in [
            ([[1, 2, 3, 4], [5, 1, 2, 3], [9, 5, 1, 2]], True),
            ([[1, 2], [2, 2]], False),
            ([[1]], True),
            ([[1, 2], [3, 1]], True),  # Diagonals: [3], [1,1], [2]
            ([[11, 74, 0, 93], [40, 11, 74, 0], [13, 40, 11, 74]], True),
        ]:
            output = self.isToeplitzMatrix([row[:] for row in matrix])
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
