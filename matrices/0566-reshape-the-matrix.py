import unittest
from typing import List

# https://leetcode.com/problems/reshape-the-matrix/

# python3 -m unittest matrices/0566-reshape-the-matrix.py


class Solution(unittest.TestCase):
    # Approach: Flatten and Rebuild
    # Key insight: Treat 2D matrix as 1D array by mapping indices.
    # For position (i,j) in m×n matrix: index = i*n + j
    # For index k in r×c matrix: row = k/c, col = k%c
    #
    # Strategy:
    # 1. Check if reshape is valid: m*n must equal r*c
    # 2. Iterate through original matrix in row-major order
    # 3. Place each element in new matrix using index mapping
    #
    # Time Complexity: O(m*n) where m,n are dimensions of original matrix
    # Space Complexity: O(r*c) for the result matrix
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        if not mat or not mat[0]:
            return mat

        m, n = len(mat), len(mat[0])

        # Check if reshape is possible
        if m * n != r * c:
            return mat

        result = [[0] * c for _ in range(r)]

        # Traverse original matrix and fill result
        for i in range(m * n):
            # Map 1D index to original 2D position
            orig_row = i // n
            orig_col = i % n

            # Map 1D index to new 2D position
            new_row = i // c
            new_col = i % c

            result[new_row][new_col] = mat[orig_row][orig_col]

        return result

    def test(self):
        for mat, r, c, expected in [
            ([[1, 2], [3, 4]], 1, 4, [[1, 2, 3, 4]]),
            ([[1, 2], [3, 4]], 2, 4, [[1, 2], [3, 4]]),
            ([[1, 2, 3, 4]], 2, 2, [[1, 2], [3, 4]]),
            ([[1]], 1, 1, [[1]]),
        ]:
            output = self.matrixReshape(mat, r, c)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
