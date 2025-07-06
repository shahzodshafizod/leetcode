import unittest

# https://leetcode.com/problems/roman-to-integer/

# python3 -m unittest hashes/0013-roman-to-integer.py


class Solution(unittest.TestCase):
    def romanToInt(self, s: str) -> int:
        table = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }
        num = 0
        nxt = 'I'
        for idx in range(len(s) - 1, -1, -1):
            curr = s[idx]
            if table[curr] < table[nxt]:
                num -= table[curr]
            else:
                num += table[curr]
            nxt = curr
        return num

    def test(self):
        for s, expected in [
            ("III", 3),
            ("LVIII", 58),
            ("MCMXCIV", 1994),
            ("II", 2),
            ("XII", 12),
            ("XXVII", 27),
            ("IV", 4),
            ("IX", 9),
        ]:
            output = self.romanToInt(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
