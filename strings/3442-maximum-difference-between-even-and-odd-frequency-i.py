import unittest

# https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/

# python3 -m unittest strings/3442-maximum-difference-between-even-and-odd-frequency-i.py

class Solution(unittest.TestCase):
    def maxDifference(self, s: str) -> int:
        count = [0] * 26
        orda = ord('a')
        for c in s:
            count[ord(c)-orda] += 1
        a1, a2 = 0, len(s)
        for cnt in count:
            if cnt&1 == 1: a1 = max(a1, cnt)
            elif cnt > 0: a2 = min(a2, cnt)
        return a1 - a2

    def test(self):
        for s, expected in [
            ("aaaaabbc", 3),
            ("abcabcab", 1),
        ]:
            output = self.maxDifference(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
