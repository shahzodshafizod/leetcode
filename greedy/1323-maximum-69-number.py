import unittest

# https://leetcode.com/problems/maximum-69-number/

# python3 -m unittest greedy/1323-maximum-69-number.py


class Solution(unittest.TestCase):
    def maximum69Number(self, num: int) -> int:
        return int(str(num).replace('6', '9', 1))

    def test(self):
        for num, expected in [
            (9669, 9969),
            (9996, 9999),
            (9999, 9999),
        ]:
            output = self.maximum69Number(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
