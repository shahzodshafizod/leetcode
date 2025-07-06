import unittest
from typing import List

# https://leetcode.com/problems/pascals-triangle/

# python3 -m unittest dp/0118-pascals-triangle.py


class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # Time: O(nn)
    # # Space: O(nn)
    # def generate(self, numRows: int) -> List[List[int]]:
    #     if numRows == 1:
    #         return [[1]]
    #     rows = self.generate(numRows-1)
    #     curr = [1] * numRows
    #     for idx in range(1, numRows-1):
    #         curr[idx] = rows[-1][idx-1] + rows[-1][idx]
    #     rows.append(curr)
    #     return rows

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(nn)
    # Space: O(nn)
    def generate(self, numRows: int) -> List[List[int]]:
        rows = []
        for length in range(1, numRows + 1):
            row = [1] * length
            for idx in range(1, length - 1):
                row[idx] = rows[-1][idx - 1] + rows[-1][idx]
            rows.append(row)
        return rows

    def testGenerate(self) -> None:
        for numRows, expected in [
            (1, [[1]]),
            (5, [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]),
            (6, [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1], [1, 5, 10, 10, 5, 1]]),
        ]:
            output = self.generate(numRows)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
