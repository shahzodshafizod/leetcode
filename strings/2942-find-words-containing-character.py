from typing import List
import unittest

# https://leetcode.com/problems/find-words-containing-character/

# python3 -m unittest strings/2942-find-words-containing-character.py

class Solution(unittest.TestCase):
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
        return [idx for idx, word in enumerate(words) if x in word]

    def test(self):
        for words, x, expected in [
            (["leet","code"], "e", [0,1]),
            (["abc","bcd","aaaa","cbc"], "a", [0,2]),
            (["abc","bcd","aaaa","cbc"], "z", []),
        ]:
            output = self.findWordsContaining(words, x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
