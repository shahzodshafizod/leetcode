from typing import List
import unittest

# https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/

# python3 -m unittest tries/3042-count-prefix-and-suffix-pairs-i.py

class Solution(unittest.TestCase):
    # def countPrefixSuffixPairs(self, words: List[str]) -> int:
    #     n = len(words)
    #     count = 0
    #     for i in range(n):
    #         for j in range(i+1, n):
    #             if len(words[i]) <= len(words[j]) and words[j].startswith(words[i]) and  words[j].endswith(words[i]):
    #                 count += 1
    #     return count

    def countPrefixSuffixPairs(self, words: List[str]) -> int:
        class TrieNode:
            def __init__(self):
                self.prefix = 0
                self.suffix = 0
                self.children = {}
        root = TrieNode()
        count = 0
        for wid, word in enumerate(words[::-1]):
            prefix = True
            currp = root
            for c in word:
                if c not in currp.children:
                    prefix = False
                    currp.children[c] = TrieNode()
                currp = currp.children[c]
                currp.prefix |= 1 << wid
            currs = root
            suffix = True
            for c in word[::-1]:
                if c not in currs.children:
                    suffix = False
                    currs.children[c] = TrieNode()
                currs = currs.children[c]
                currs.suffix |= 1 << wid
            if prefix and suffix:
                count += (currp.prefix & currs.suffix).bit_count() - 1
        return count

    def test(self):
        for words, expected in [
            (["a","aba","ababa","aa"], 4),
            (["pa","papa","ma","mama"], 2),
            (["abab","ab"], 0),
        ]:
            output = self.countPrefixSuffixPairs(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
