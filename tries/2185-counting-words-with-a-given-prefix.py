from typing import List
import unittest

# https://leetcode.com/problems/counting-words-with-a-given-prefix/

# python3 -m unittest tries/2185-counting-words-with-a-given-prefix.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(m*n), m=len(pref), n=len(words)
    # # Space: O(1)
    # def prefixCount(self, words: List[str], pref: str) -> int:
    #     plen = len(pref)
    #     count = 0
    #     for word in words:
    #         pid, wid, wlen = 0, 0, len(word)
    #         while wid < wlen and pid < plen:
    #             if wlen - wid < plen - pid or word[wid] != pref[pid]:
    #                 break
    #             wid += 1
    #             pid += 1
    #         if pid == plen:
    #             count += 1
    #     return count

    # # Approach: Tries
    # # Time: O(m + n*l), m=len(pref), n=len(words), l=len(words[i])
    # # Space: O(n*l)
    # def prefixCount(self, words: List[str], pref: str) -> int:
    #     class TrieNode:
    #         def __init__(self):
    #             self.count = 0
    #             self.children = {}
    #     root = TrieNode()
    #     for word in words:
    #         curr = root
    #         for c in word:
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode()
    #             curr = curr.children[c]
    #             curr.count += 1
    #     count = 0
    #     curr = root
    #     for c in pref:
    #         if c not in curr.children:
    #             count = 0
    #             break
    #         curr = curr.children[c]
    #         count = curr.count
    #     return count

    # Approach: Built-In Methods
    # Time: O(m*n), m=len(pref), n=len(words)
    # Space: O(1)
    def prefixCount(self, words: List[str], pref: str) -> int:
        return sum(word.startswith(pref) for word in words)

    def test(self):
        for words, pref, expected in [
            (["pay", "attention", "practice", "attend"], "at", 2),
            (["leetcode", "win", "loops", "success"], "code", 0),
        ]:
            output = self.prefixCount(words, pref)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
