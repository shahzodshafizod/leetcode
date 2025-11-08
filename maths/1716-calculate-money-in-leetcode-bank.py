import unittest

# https://leetcode.com/problems/calculate-money-in-leetcode-bank/

# python3 -m unittest maths/1716-calculate-money-in-leetcode-bank.py


class Solution(unittest.TestCase):
    def totalMoney(self, n: int) -> int:
        total, money = 0, 0
        for day in range(n):
            if day % 7 == 0:
                money = day // 7 + 1
            total += money
            money += 1
        return total

    def test(self):
        for n, expected in [
            (4, 10),
            (10, 37),
            (20, 96),
        ]:
            output = self.totalMoney(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
