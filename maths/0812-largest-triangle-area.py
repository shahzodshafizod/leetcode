from typing import List
import unittest

# https://leetcode.com/problems/largest-triangle-area/

# python3 -m unittest maths/0812-largest-triangle-area.py

# Heron's formula:
# Area = √(s(s-a)(s-b)(s-c))
# s = (a+b+c)/2
# a = √(x2-x1)^2 + (y2-y1)^2)

# Shoelace formula for a triangle
# Area = ½ |(x₁y₂ + x₂y₃ + x₃y₁) - (y₁x₂ + y₂x₃ + y₃x₁)|


class Solution(unittest.TestCase):
    def largestTriangleArea(self, points: List[List[int]]) -> float:
        n = len(points)
        larea = 0
        for i in range(n - 1):
            x1, y1 = points[i]
            for j in range(i + 1, n - 1):
                x2, y2 = points[j]
                for k in range(j + 1, n):
                    x3, y3 = points[k]
                    area = 0.5 * abs((x1 * y2 + x2 * y3 + x3 * y1) - (y1 * x2 + y2 * x3 + y3 * x1))
                    larea = max(larea, area)
        return larea

    def test(self):
        for points, expected in [
            ([[0, 0], [0, 1], [1, 0], [0, 2], [2, 0]], 2.00000),
            ([[1, 0], [0, 0], [0, 1]], 0.50000),
        ]:
            output = self.largestTriangleArea(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
