from math import sqrt
from typing import List
import unittest

# https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/

# python3 -m unittest maths/3000-maximum-area-of-longest-diagonal-rectangle.py


class Solution(unittest.TestCase):
    def areaOfMaxDiagonal(self, dimensions: List[List[int]]) -> int:
        diag, area = 0, 0
        for l, w in dimensions:
            d = sqrt(l * l + w * w)
            a = l * w
            if d > diag or d == diag and a > area:
                diag = d
                area = a
        return area

    def test(self):
        for dimensions, expected in [
            ([[9, 3], [8, 6]], 48),
            ([[3, 4], [4, 3]], 12),
        ]:
            output = self.areaOfMaxDiagonal(dimensions)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
