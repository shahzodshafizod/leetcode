from typing import List
import unittest

# https://leetcode.com/problems/richest-customer-wealth/

# python3 -m unittest matrices/1672-richest-customer-wealth.py


class Solution(unittest.TestCase):
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        return max(sum(money) for money in accounts)

    def test(self):
        for accounts, expected in [
            ([[1, 2, 3], [3, 2, 1]], 6),
            ([[1, 5], [7, 3], [3, 5]], 10),
            ([[2, 8, 7], [7, 1, 3], [1, 9, 5]], 17),
        ]:
            output = self.maximumWealth(accounts)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
