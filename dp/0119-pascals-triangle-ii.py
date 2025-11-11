from typing import List
import unittest

# https://leetcode.com/problems/pascals-triangle-ii/

# python3 -m unittest dp/0119-pascals-triangle-ii.py


class Solution(unittest.TestCase):
    def getRow(self, rowIndex: int) -> List[int]:
        row = [1]
        for n in range(rowIndex):
            prev = 0
            for i in range(n + 1):
                prev, row[i] = row[i], prev + row[i]
            row.append(1)
        return row

    def test(self):
        for rowIndex, expected in [
            (3, [1, 3, 3, 1]),
            (0, [1]),
            (1, [1, 1]),
        ]:
            output = self.getRow(rowIndex)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
