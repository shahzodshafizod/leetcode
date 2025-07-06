from itertools import zip_longest
import unittest

# https://leetcode.com/problems/word-pattern/

# python3 -m unittest hashes/0290-word-pattern.py


class Solution(unittest.TestCase):
    def wordPattern(self, pattern: str, s: str) -> bool:
        # The counts of distinct elements in two groups
        # and the count of distinct mappings are all equal.
        s = s.split(' ')
        return len(set(s)) == len(set(pattern)) == len(set(zip_longest(pattern, s)))

    def test(self):
        for pattern, s, expected in [
            ("abba", "dog cat cat dog", True),
            ("abba", "dog cat cat fish", False),
            ("aaaa", "dog cat cat dog", False),
        ]:
            output = self.wordPattern(pattern, s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
