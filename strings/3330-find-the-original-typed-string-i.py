import unittest

# https://leetcode.com/problems/find-the-original-typed-string-i/

# python3 -m unittest strings/3330-find-the-original-typed-string-i.py


class Solution(unittest.TestCase):
    def possibleStringCount(self, word: str) -> int:
        return sum(word[i - 1] == word[i] for i in range(1, len(word))) + 1

    def test(self):
        for word, expected in [
            ("abbcccc", 5),
            ("abcd", 1),
            ("aaaa", 4),
        ]:
            output = self.possibleStringCount(word)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
