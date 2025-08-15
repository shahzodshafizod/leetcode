import unittest

# https://leetcode.com/problems/largest-3-same-digit-number-in-string/

# python3 -m unittest strings/2264-largest-3-same-digit-number-in-string.py


class Solution(unittest.TestCase):
    def largestGoodInteger(self, num: str) -> str:
        good = ""
        for i in range(2, len(num)):
            if num[i - 2] == num[i - 1] == num[i]:
                good = max(good, num[i - 2 : i + 1])
        return good

    def test(self):
        for num, expected in [
            ("6777133339", "777"),
            ("2300019", "000"),
            ("42352338", ""),
        ]:
            output = self.largestGoodInteger(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
