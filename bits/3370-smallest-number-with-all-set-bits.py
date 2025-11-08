import unittest

# https://leetcode.com/problems/smallest-number-with-all-set-bits/

# python3 -m unittest bits/3370-smallest-number-with-all-set-bits.py


class Solution(unittest.TestCase):
    def smallestNumber(self, n: int) -> int:
        return (1 << n.bit_length()) - 1

    def test(self):
        for n, expected in [
            (5, 7),
            (10, 15),
            (3, 3),
        ]:
            output = self.smallestNumber(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
