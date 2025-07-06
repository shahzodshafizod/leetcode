import unittest

# https://leetcode.com/problems/sum-of-digits-of-string-after-convert/

# python3 -m unittest strings/1945-sum-of-digits-of-string-after-convert.py


class Solution(unittest.TestCase):
    def getLucky(self, s: str, k: int) -> int:
        number = 0
        for c in s:
            order = ord(c) - ord('a') + 1
            number += order // 10 + order % 10
        for _ in range(k - 1):
            digit_sum = 0
            while number:
                number, digit = divmod(number, 10)
                digit_sum += digit
            number = digit_sum
        return number

    def test(self):
        for s, k, expected in [
            ("iiii", 1, 36),
            ("leetcode", 2, 6),
            ("zbax", 2, 8),
            ("vbyytoijnbgtyrjlsc", 2, 15),
            ("hvmhoasabaymnmsd", 1, 79),
        ]:
            output = self.getLucky(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
