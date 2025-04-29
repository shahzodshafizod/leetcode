import unittest

# https://leetcode.com/problems/number-of-segments-in-a-string/

# python3 -m unittest strings/0434-number-of-segments-in-a-string.py

class Solution(unittest.TestCase):
    def countSegments(self, s: str) -> int:
        return len(s.split())

    def test(self):
        for s, expected in [
            ("Hello, my name is John", 5),
            ("Hello", 1),
            (", , , , a, eaefa", 6),
            (" ", 0),
            ("", 0),
            (", , , , a, e a e f a", 10),
        ]:
            output = self.countSegments(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
