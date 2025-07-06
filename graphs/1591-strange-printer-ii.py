from collections import defaultdict
from itertools import product
from typing import List
import unittest

# https://leetcode.com/problems/strange-printer-ii/

# python3 -m unittest graphs/1591-strange-printer-ii.py


class Solution(unittest.TestCase):
    # Approach: Topological Sort
    # Time: O(mn), m=# of ROWS, n=# of COLS
    # Space: O(c), c=# of colors
    def isPrintable(self, targetGrid: List[List[int]]) -> bool:
        border = {}
        ROWS, COLS = len(targetGrid), len(targetGrid[0])
        for row in range(ROWS):
            for col in range(COLS):
                color = targetGrid[row][col]
                if color not in border:
                    border[color] = [row, row, col, col]
                else:
                    border[color][1] = max(border[color][1], row)
                    border[color][2] = min(border[color][2], col)
                    border[color][3] = max(border[color][3], col)
        seen = defaultdict(bool)

        def dfs(row: int, col: int) -> bool:
            color = targetGrid[row][col]
            if color == 0:
                return True
            if seen[color]:
                return False
            seen[color] = True
            for r in range(border[color][0], border[color][1] + 1):
                for c in range(border[color][2], border[color][3] + 1):
                    if targetGrid[r][c] != color and not dfs(r, c):
                        return False
                    targetGrid[r][c] = 0
            return True

        return all(dfs(row, col) for row, col in product(range(ROWS), range(COLS)))

    def test(self):
        for targetGrid, expected in [
            ([[1, 1, 1, 1], [1, 2, 2, 1], [1, 2, 2, 1], [1, 1, 1, 1]], True),
            ([[1, 1, 1, 1], [1, 1, 3, 3], [1, 1, 3, 4], [5, 5, 1, 4]], True),
            ([[1, 2, 1], [2, 1, 2], [1, 2, 1]], False),
        ]:
            output = self.isPrintable(targetGrid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
