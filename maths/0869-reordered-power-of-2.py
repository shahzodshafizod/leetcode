import unittest

# https://leetcode.com/problems/reordered-power-of-2/

# python3 -m unittest maths/0869-reordered-power-of-2.py


class Solution(unittest.TestCase):
    def reorderedPowerOf2(self, n: int) -> bool:
        sn = sorted(str(n))
        return any(sorted(str(1 << i)) == sn for i in range(32))

    def test(self):
        for n, expected in [
            (1, True),
            (10, False),
        ]:
            output = self.reorderedPowerOf2(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
