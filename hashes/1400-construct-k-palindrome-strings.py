from collections import Counter
import unittest

# https://leetcode.com/problems/construct-k-palindrome-strings/

# python3 -m unittest hashes/1400-construct-k-palindrome-strings.py

class Solution(unittest.TestCase):
    def canConstruct(self, s: str, k: int) -> bool:
        return len(s) >= k and sum(1 for cnt in Counter(s).values() if cnt&1 == 1) <= k

    def test(self):
        for s, k, expected in [
            ("annabelle", 2, True),
            ("leetcode", 3, False),
            ("true", 4, True),
        ]:
            output = self.canConstruct(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
