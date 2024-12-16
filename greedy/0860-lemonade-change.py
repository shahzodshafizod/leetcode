from typing import List
import unittest

# https://leetcode.com/problems/lemonade-change/

# python3 -m unittest greedy/0860-lemonade-change.py

class Solution(unittest.TestCase):
    def lemonadeChange(self, bills: List[int]) -> bool:
        tenStack, fiveStack = 0, 0
        for bill in bills:
            if bill == 5:
                fiveStack += 1
            elif bill == 10:
                tenStack += 1
                fiveStack -= 1
            else:
                if tenStack > 0:
                    tenStack -= 1
                    fiveStack -= 1
                else:
                    fiveStack -= 3
            if fiveStack < 0:
                return False
        return True

    def test(self):
        for bills, expected in [
            ([5, 5, 5, 10, 20], True),
            ([5, 5, 10, 10, 20], False),
            ([5, 5, 5, 5, 10, 20, 10], True),
            ([5, 5, 5, 20], True),
            ([5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5], True),
            ([5, 5, 5, 5, 20, 20, 5, 5, 20, 5], False),
            ([5, 5, 10, 10, 5, 20, 5, 10, 5, 5], True),
        ]:
            output = self.lemonadeChange(bills)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
