from typing import List
import unittest

# https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

# python3 -m unittest matrices/1337-the-k-weakest-rows-in-a-matrix.py


class Solution(unittest.TestCase):
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        soldiers = [sum(row) for row in mat]
        indices = list(range(len(mat)))
        indices.sort(key=lambda idx: soldiers[idx])
        return indices[:k]

    def test(self):
        for mat, k, expected in [
            ([[1, 1, 0, 0, 0], [1, 1, 1, 1, 0], [1, 0, 0, 0, 0], [1, 1, 0, 0, 0], [1, 1, 1, 1, 1]], 3, [2, 0, 3]),
            ([[1, 0, 0, 0], [1, 1, 1, 1], [1, 0, 0, 0], [1, 0, 0, 0]], 2, [0, 2]),
        ]:
            output = self.kWeakestRows(mat, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
