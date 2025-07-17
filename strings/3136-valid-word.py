import unittest

# https://leetcode.com/problems/valid-word/

# python3 -m unittest strings/3136-valid-word.py


class Solution(unittest.TestCase):
    def isValid(self, word: str) -> bool:
        has_vowel, has_consonent = False, False
        for c in word.lower():
            if c in "@#$":
                return False
            if c in "aeiou":
                has_vowel = True
            else:
                has_consonent = True
        return has_vowel and has_consonent

    def test(self):
        for word, expected in [
            ("234Adas", True),
            ("b3", False),
            ("a3$e", False),
        ]:
            output = self.isValid(word)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
