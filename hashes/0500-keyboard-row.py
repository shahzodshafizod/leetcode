from typing import List
import unittest

# https://leetcode.com/problems/keyboard-row/

# python3 -m unittest hashes/0500-keyboard-row.py

class Solution(unittest.TestCase):
    def findWords(self, words: List[str]) -> List[str]:
        codes = {c:idx for idx, row in enumerate(["qwertyuiop", "asdfghjkl", "zxcvbnm"]) for c in row}
        res = []
        for word in words:
            code = codes[word[0].lower()]
            ok = True
            for c in word.lower():
                if codes[c] != code:
                    ok = False
                    break
            if ok:
                res.append(word)
        return res

    def test(self):
        for words, expected in [
            (["Hello","Alaska","Dad","Peace"], ["Alaska","Dad"]),
            (["omk"], []),
            (["adsdf","sfd"], ["adsdf","sfd"]),
        ]:
            output = self.findWords(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
