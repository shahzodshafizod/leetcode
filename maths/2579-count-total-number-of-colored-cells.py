import unittest

# https://leetcode.com/problems/count-total-number-of-colored-cells/

# python3 -m unittest maths/2579-count-total-number-of-colored-cells.py

class Solution(unittest.TestCase):
    def coloredCells(self, n: int) -> int:
        return 1 + n * (n-1) * 2
        return 1 + sum(4*i for i in range(1, n))

    def test(self):
        for n, expected in [
            (1, 1),
            (2, 5),
            (100000, 19999800001),
        ]:
            output = self.coloredCells(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
