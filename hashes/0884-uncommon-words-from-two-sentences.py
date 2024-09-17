from typing import Counter, List
import unittest

# https://leetcode.com/problems/uncommon-words-from-two-sentences/

# python3 -m unittest hashes/0884-uncommon-words-from-two-sentences.py

class Solution(unittest.TestCase):
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        return [word for word,count in Counter((s1+" "+s2).split()).items() if count == 1]

    def testUncommonFromSentences(self) -> None:
        for s1, s2, expected in [
            ("this apple is sweet", "this apple is sour", ["sweet","sour"]),
            ("apple apple", "banana", ["banana"]),
            ("this apple is sweet sweet sweet", "this apple is sour", ["sour"]),
            ("this apple is sweet sweet sweet and good", "this apple is sour", ["and","good","sour"]),
            ("s z z z s", "s z ejt", ["ejt"]),
        ]:
            output = self.uncommonFromSentences(s1, s2)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
