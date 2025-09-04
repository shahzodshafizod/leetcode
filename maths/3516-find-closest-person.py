import unittest

# https://leetcode.com/problems/find-closest-person/

# python3 -m unittest maths/3516-find-closest-person.py


class Solution(unittest.TestCase):
    def findClosest(self, x: int, y: int, z: int) -> int:
        dxz = abs(z - x)
        dyz = abs(z - y)
        return 1 if dxz < dyz else 2 if dyz < dxz else 0

    def test(self):
        for x, y, z, expected in [
            (2, 7, 4, 1),
            (2, 5, 6, 2),
            (1, 5, 3, 0),
        ]:
            output = self.findClosest(x, y, z)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
