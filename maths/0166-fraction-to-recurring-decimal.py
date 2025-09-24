from typing import List, Dict
import unittest

# https://leetcode.com/problems/fraction-to-recurring-decimal/

# python3 -m unittest maths/0166-fraction-to-recurring-decimal.py


class Solution(unittest.TestCase):
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        res: List[str] = []
        if numerator * denominator < 0:
            res.append('-')
        numerator = abs(numerator)
        denominator = abs(denominator)
        res.append(str(numerator // denominator))
        remainder = numerator % denominator
        if remainder == 0:
            return "".join(res)
        res.append('.')
        fractionals: Dict[int, int] = {}
        while remainder != 0:
            index = fractionals.get(remainder, -1)
            if index >= 0:
                res.insert(index, '(')
                res.append(')')
                break
            fractionals[remainder] = len(res)
            remainder *= 10
            res.append(str(remainder // denominator))
            remainder %= denominator
        return "".join(res)

    def test(self):
        for numerator, denominator, expected in [
            (1, 2, "0.5"),
            (2, 1, "2"),
            (4, 333, "0.(012)"),
            (-50, 8, "-6.25"),
            (0, -5, "0"),
            (-1, -2147483648, "0.0000000004656612873077392578125"),
            (1, 6, "0.1(6)"),
        ]:
            output = self.fractionToDecimal(numerator, denominator)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
