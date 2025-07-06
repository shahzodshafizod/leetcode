from functools import cache
import unittest

# https://leetcode.com/problems/painting-a-grid-with-three-different-colors/

# python3 -m unittest dp/1931-painting-a-grid-with-three-different-colors.py


class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming
    # Time: O(m*n*2^10)
    # Space: O(m*n*2^10)
    def colorTheGrid(self, m: int, n: int) -> int:
        shift_pos = (m - 1) * 2

        @cache
        def dp(mask: int, row: int, col: int) -> int:
            if row == m:
                return dp(mask, 0, col + 1)
            if col == n:
                return 1
            up = mask >> shift_pos if row > 0 else 0
            left = mask & 0b11 if col > 0 else 0
            mask >>= 2
            res = 0
            for color in (1, 2, 3):
                if color == up or color == left:
                    continue
                color <<= shift_pos
                res += dp(mask | color, row + 1, col)
            return res % int(1e9 + 7)

        return dp(0, 0, 0)

    def test(self):
        for m, n, expected in [
            (1, 1, 3),
            (1, 2, 6),
            (5, 5, 580986),
        ]:
            output = self.colorTheGrid(m, n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
