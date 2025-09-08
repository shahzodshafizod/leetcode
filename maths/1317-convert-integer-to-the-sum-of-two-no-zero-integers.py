from typing import List
import unittest

# https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

# python3 -m unittest maths/1317-convert-integer-to-the-sum-of-two-no-zero-integers.py


class Solution(unittest.TestCase):
    def getNoZeroIntegers(self, n: int) -> List[int]:
        for a in range(1, n // 2 + 1):
            b = n - a
            if '0' not in f"{a}{b}":
                return [a, b]
        return [-1, -1]

    def test(self):
        for n, expected in [
            (2, [1, 1]),
            (11, [2, 9]),
        ]:
            output = self.getNoZeroIntegers(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
