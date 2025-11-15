import unittest

# https://leetcode.com/problems/excel-sheet-column-number/

# python3 -m unittest maths/0171-excel-sheet-column-number.py


class Solution(unittest.TestCase):
    def titleToNumber(self, columnTitle: str) -> int:
        ordA = ord('A')
        columnNumber, base = 0, 1
        for c in columnTitle[::-1]:
            columnNumber += (ord(c) - ordA + 1) * base
            base *= 26
        return columnNumber

    def test(self):
        for columnTitle, expected in [
            ("A", 1),
            ("AB", 28),
            ("ZY", 701),
        ]:
            output = self.titleToNumber(columnTitle)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
