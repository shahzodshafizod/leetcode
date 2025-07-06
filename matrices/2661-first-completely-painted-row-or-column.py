from typing import List
import unittest

# https://leetcode.com/problems/first-completely-painted-row-or-column/

# python3 -m unittest matrices/2661-first-completely-painted-row-or-column.py


class Solution(unittest.TestCase):
    # Approach: Counting
    # Time: O(mn)
    # Space: O(mn)
    def firstCompleteIndex(self, arr: List[int], mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        rcount = [n] * m
        ccount = [m] * n
        positions = [None] * (m * n + 1)
        for row in range(m):
            for col in range(n):
                positions[mat[row][col]] = (row, col)
        idx = -1
        for idx in range(m * n):
            row, col = positions[arr[idx]]
            rcount[row] -= 1
            ccount[col] -= 1
            if not rcount[row] or not ccount[col]:
                break
        return idx

    def test(self):
        for arr, mat, expected in [
            ([1, 3, 4, 2], [[1, 4], [2, 3]], 2),
            ([2, 8, 7, 4, 1, 3, 5, 6, 9], [[3, 2, 5], [1, 4, 6], [8, 7, 9]], 3),
        ]:
            output = self.firstCompleteIndex(arr, mat)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
