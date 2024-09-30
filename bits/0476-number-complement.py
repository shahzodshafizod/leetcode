import unittest

# https://leetcode.com/problems/number-complement/

# python3 -m unittest bits/0476-number-complement.py

class Solution(unittest.TestCase):
    def findComplement(self, num: int) -> int:
        bit = 1
        while bit <= num:
            num ^= bit
            bit <<= 1
        return num

    def test(self) -> None:
        for num, expected in [
            (5, 2),
		    (1, 0),
            (7, 0),
            (10, 5),
            (2147483647, 0),
        ]:
            output = self.findComplement(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
