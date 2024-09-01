from typing import List
import unittest

# https://leetcode.com/problems/convert-1d-array-into-2d-array/

# python3 -m unittest matrices/2022-convert-1d-array-into-2d-array.py

class Solution(unittest.TestCase):
    def construct2DArray(self, original: List[int], m: int, n: int) -> List[List[int]]:
        # return [original[n*i:n*(i+1)] for i in range(m)] if m*n == len(original) else []
        if m*n != len(original):
            return []
        matrix = [[0]*n for _ in range(m)]
        idx = 0
        for row in range(m):
            for col in range(n):
                matrix[row][col] = original[idx]
                idx += 1
        # for idx in range(len(original)):
        #     matrix[idx//n][idx%n] = original[idx]
        return matrix

    def testConstruct2DArray(self) -> None:
        for original, m, n, expected in [
            ([1,2,3,4], 2, 2, [[1,2],[3,4]]),
            ([1,2,3], 1, 3, [[1,2,3]]),
            ([1,2], 1, 1, []),
        ]:
            output = self.construct2DArray(original, m, n)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
