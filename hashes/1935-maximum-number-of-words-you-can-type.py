import unittest

# https://leetcode.com/problems/maximum-number-of-words-you-can-type/

# python3 -m unittest hashes/1935-maximum-number-of-words-you-can-type.py


class Solution(unittest.TestCase):
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        return sum(0 if set(word) & set(brokenLetters) else 1 for word in text.split())

    def test(self):
        for text, brokenLetters, expected in [
            ("hello world", "ad", 1),
            ("leet code", "lt", 1),
            ("leet code", "e", 0),
        ]:
            output = self.canBeTypedWords(text, brokenLetters)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
