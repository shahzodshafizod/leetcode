from typing import List
import unittest

# https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/

# python3 -m unittest tries/3045-count-prefix-and-suffix-pairs-ii.py


class Solution(unittest.TestCase):
    def countPrefixSuffixPairs(self, words: List[str]) -> int:
        class TrieNode:
            def __init__(self):
                self.count = 0
                self.children = {}

        count = 0
        root = TrieNode()
        for word in words[::-1]:
            curr = root
            for pair in zip(word, word[::-1]):
                if pair not in curr.children:
                    curr.children[pair] = TrieNode()
                curr = curr.children[pair]
                curr.count += 1
            count += curr.count - 1
        return count

    def test(self):
        for words, expected in [
            (["a", "aba", "ababa", "aa"], 4),
            (["pa", "papa", "ma", "mama"], 2),
            (["abab", "ab"], 0),
        ]:
            output = self.countPrefixSuffixPairs(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
