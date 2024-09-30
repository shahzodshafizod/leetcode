from typing import List
import unittest

# https://leetcode.com/problems/longest-common-prefix/

# python3 -m unittest tries/0014-longest-common-prefix.py

class Solution(unittest.TestCase):
    # Frute-Force
    # N = len(strs)
    # Time: O(N x len(prefix))
    # Space: O(1)
    def longestCommonPrefix(self, strs: List[str]) -> str:
        idx = 0
        n = min([len(word) for word in strs])
        exit = False
        while not exit and idx < n:
            c = strs[0][idx]
            for word in strs:
                if word[idx] != c:
                    exit = True
                    idx -= 1
                    break
            idx += 1
        return strs[0][:idx]

    # # Trie
    # def longestCommonPrefix(self, strs: List[str]) -> str:
    #     class TrieNode:
    #         def __init__(self) -> None:
    #             self.words = 0
    #             self.children = {}
    #     # insert
    #     root = TrieNode()
    #     for word in strs:
    #         curr = root
    #         curr.words += 1
    #         for c in word:
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode()
    #             curr = curr.children[c]
    #             curr.words += 1
    #     def search(curr: TrieNode, prefix: str, n: int) -> str:
    #         if curr.words != n:
    #             return ""
    #         for c, child in curr.children.items():
    #             if not child:
    #                 continue
    #             pref = search(child, prefix+c, n)
    #             if len(pref) > len(prefix):
    #                 prefix = pref
    #         return prefix
    #     return search(root, "", len(strs))

    def test(self) -> None:
        for strs, expected in [
            (["flower","flow","flight"], "fl"),
            (["dog","racecar","car"], ""),
            (["reflower","flow","flight"], ""),
            (["",""], ""),
            (["aaaa","aa","aaa","aaaaaaaaaaaaaaaaaaaaaaaaaaa"], "aa"),
            (["flower","fkow"], "f"),
            (["spoiler","flower","flow","flight"], ""),
            (["c","acc","ccc"], ""),
        ]:
            output = self.longestCommonPrefix(strs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
