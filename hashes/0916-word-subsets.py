from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/word-subsets/

# python3 -m unittest hashes/0916-word-subsets.py

class Solution(unittest.TestCase):
    def wordSubsets(self, words1: List[str], words2: List[str]) -> List[str]:
        freq = Counter()
        for word in words2:
            for ch, cnt in Counter(word).items():
                freq[ch] = max(freq[ch], cnt)
        return [word for word in words1 if Counter(word) >= freq]

    def test(self):
        for words1, words2, expected in [
            (["amazon","apple","facebook","google","leetcode"], ["e","o"], ["facebook","google","leetcode"]),
            (["amazon","apple","facebook","google","leetcode"], ["l","e"], ["apple","google","leetcode"])
        ]:
            output = self.wordSubsets(words1, words2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
