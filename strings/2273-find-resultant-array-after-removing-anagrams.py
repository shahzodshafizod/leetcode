from itertools import groupby
from typing import List
import unittest

# https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/

# python3 -m unittest strings/2273-find-resultant-array-after-removing-anagrams.py


class Solution(unittest.TestCase):
    def removeAnagrams(self, words: List[str]) -> List[str]:
        # return [words[i] for i in range(len(words)) if i == 0 or sorted(words[i-1]) != sorted(words[i])]
        return [next(group) for _, group in groupby(words, sorted)]

    def test(self):
        for words, expected in [
            (["abba","baba","bbaa","cd","cd"], ["abba","cd"]),
            (["a","b","c","d","e"], ["a","b","c","d","e"]),
        ]:
            output = self.removeAnagrams(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
