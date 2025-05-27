import unittest

# https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/

# python3 -m unittest maths/2894-divisible-and-non-divisible-sums-difference.py

"""
num2 = 1m + 2m + 3m + 4m
num2 = (1 + 2 + 3 + 4) * m
num2 = (k * (k+1) // 2) * m
num1 = n * (n+1) // 2 - num2
res = num1 - num2
res = n * (n+1) // 2 - num2 - num2
res = n * (n+1) // 2 - 2 * num2
res = n * (n+1) // 2 - [2] * (k * (k+1) // [2]) * m
res = n * (n+1) // 2 - k * (k+1) * m
"""

class Solution(unittest.TestCase):
    def differenceOfSums(self, n: int, m: int) -> int:
        k = n // m
        return n * (n+1) // 2 - k * (k+1) * m

    def test(self):
        for n, m, expected in [
            (10, 3, 19),
            (5, 6, 15),
            (5, 1, -15),
        ]:
            output = self.differenceOfSums(n, m)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
