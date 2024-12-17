from typing import Counter
import unittest

# https://leetcode.com/problems/valid-anagram/

# python3 -m unittest hashes/0242-valid-anagram.py

class Solution(unittest.TestCase):
    def isAnagram(self, s: str, t: str) -> bool:
        # return sorted(s) == sorted(t)
        return Counter(s) == Counter(t)

    def testIsAnagram(self) -> None:
        for s, t, expected in [
            ("anagram", "nagaram", True),
            ("rat", "car", False),
            ("rrat", "tar", False),
            ("a", "abb", False), # a clear example why bit manipulation doesn't work
        ]:
            output = self.isAnagram(s, t)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
