import math
import unittest

# https://leetcode.com/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/

# python3 -m unittest maths/3405-count-the-number-of-arrays-with-k-matching-adjacent-elements.py


class Solution(unittest.TestCase):
    def countGoodArrays(self, n: int, m: int, k: int) -> int:
        MOD = 10**9 + 7
        return m * pow(m - 1, n - 1 - k, MOD) * math.comb(n - 1, k) % MOD

    def test(self):
        for n, m, k, expected in [
            (3, 2, 1, 4),
            (4, 2, 2, 6),
            (5, 2, 0, 2),
        ]:
            output = self.countGoodArrays(n, m, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
