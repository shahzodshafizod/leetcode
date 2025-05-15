from typing import List
import unittest

# https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/

# python3 -m unittest greedy/2900-longest-unequal-adjacent-groups-subsequence-i.py

class Solution(unittest.TestCase):
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        return [words[i] for i in range(len(words)) if i == 0 or groups[i] != groups[i-1]]

    def test(self):
        for words, groups, expected in [
            (["e","a","b"], [0,0,1], ["e","b"]),
            (["a","b","c","d"], [1,0,1,1], ["a","b","c"]),
        ]:
            output = self.getLongestSubsequence(words, groups)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
