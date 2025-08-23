from typing import List
import unittest

# https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-ii/

# python3 -m unittest matrices/3197-find-the-minimum-area-to-cover-all-ones-ii.py


class Solution(unittest.TestCase):
    def minimumSum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])

        def area(lb: int, rb: int, tb: int, db: int) -> int:  # lb is left-bound
            left, right = rb + 1, lb - 1
            top, down = db + 1, tb - 1
            for row in range(tb, db + 1):
                for col in range(lb, rb + 1):
                    if grid[row][col] == 1:
                        left = min(left, col)
                        right = max(right, col)
                        top = min(top, row)
                        down = max(down, row)
            if left > rb:
                return m * n + 1  # MAX
            return (right - left + 1) * (down - top + 1)

        lb, rb = n, 0
        tb, db = m, 0
        for row in range(m):
            for col in range(n):
                if grid[row][col] == 1:
                    lb = min(lb, col)
                    rb = max(rb, col)
                    tb = min(tb, row)
                    db = max(db, row)

        res = m * n
        for row in range(tb, db + 1):
            for col in range(lb, rb + 1):
                # horizontal-up: divide horizontally, up is also divided
                res = min(res, area(lb, col, tb, row) + area(col + 1, rb, tb, row) + area(lb, rb, row + 1, db))
                # horizontal-down: divide horizontally, down is also divided
                res = min(res, area(lb, rb, 0, row) + area(lb, col, row + 1, db) + area(col + 1, rb, row + 1, db))
                # vertical-left: divide vertically, left is also divided
                res = min(res, area(lb, col, tb, row) + area(lb, col, row + 1, db) + area(col + 1, rb, tb, db))
                # vertical-right: divide vertically, right is also divided
                res = min(res, area(lb, col, tb, db) + area(col + 1, rb, tb, row) + area(col + 1, rb, row + 1, db))
        # two horizontal cuts
        for row1 in range(m - 2):
            for row2 in range(row1, m - 1):
                res = min(res, area(lb, rb, tb, row1) + area(lb, rb, row1 + 1, row2) + area(lb, rb, row2 + 1, db))
        # two vertical cuts
        for col1 in range(n - 2):
            for col2 in range(col1, n - 1):
                res = min(res, area(lb, col1, tb, db) + area(col1 + 1, col2, tb, db) + area(col2 + 1, rb, tb, db))
        return res

    def test(self):
        for grid, expected in [
            ([[1, 0, 1], [1, 1, 1]], 5),
            ([[1, 0, 1, 0], [0, 1, 0, 1]], 5),
        ]:
            output = self.minimumSum(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
