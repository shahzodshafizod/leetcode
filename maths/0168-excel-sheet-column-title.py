from typing import List
import unittest

# https://leetcode.com/problems/excel-sheet-column-title/

# python3 -m unittest maths/0168-excel-sheet-column-title.py


class Solution(unittest.TestCase):
    def convertToTitle(self, columnNumber: int) -> str:
        ans: List[str] = []
        while columnNumber > 0:
            columnNumber -= 1
            ans.append(chr(ord('A') + columnNumber % 26))
            columnNumber //= 26
        return "".join(ans[::-1])

    def test(self):
        for columnNumber, expected in [
            (1, "A"),
            (28, "AB"),
            (701, "ZY"),
        ]:
            output = self.convertToTitle(columnNumber)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
