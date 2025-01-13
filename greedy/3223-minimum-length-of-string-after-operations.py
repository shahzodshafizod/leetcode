from collections import Counter
import unittest

# https://leetcode.com/problems/minimum-length-of-string-after-operations/

# python3 -m unittest greedy/3223-minimum-length-of-string-after-operations.py

class Solution(unittest.TestCase):
    def minimumLength(self, s: str) -> int:
        return len(s) - sum(cnt - 2 + (cnt&1) for cnt in Counter(s).values())

    def test(self):
        for s, expected in [
            ("abaacbcbb", 5),
            ("aa", 2),
        ]:
            output = self.minimumLength(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
