import unittest

# https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/

# python3 -m unittest strings/3461-check-if-digits-are-equal-in-string-after-operations-i.py


class Solution(unittest.TestCase):
    def hasSameDigits(self, s: str) -> bool:
        stack = list(map(int, s))
        for size in range(len(stack), 2, -1):
            for i in range(size - 1):
                stack[i] = (stack[i] + stack[i + 1]) % 10
        return stack[0] == stack[1]

    def test(self):
        for s, expected in [
            ("3902", True),
            ("34789", False),
        ]:
            output = self.hasSameDigits(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
