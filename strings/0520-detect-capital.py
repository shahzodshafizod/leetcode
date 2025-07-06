import unittest

# https://leetcode.com/problems/detect-capital/

# python3 -m unittest strings/0520-detect-capital.py


class Solution(unittest.TestCase):
    def detectCapitalUse(self, word: str) -> bool:
        return word.isupper() or word.islower() or word.istitle()

    def test(self):
        for word, expected in [
            ("USA", True),
            ("uSA", False),
            ("FlaG", False),
            ("LeetCode", False),
            ("g", True),
        ]:
            output = self.detectCapitalUse(word)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
