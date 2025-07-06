from typing import List
import unittest

# https://leetcode.com/problems/string-matching-in-an-array/

# python3 -m unittest strings/1408-string-matching-in-an-array.py


class Solution(unittest.TestCase):
    # # Approach: Trie
    # # Time: O(nmm), n=len(words), m=len(words[i])
    # # Space: O(nm)
    # def stringMatching(self, words: List[str]) -> List[str]:
    #     class TrieNode:
    #         def __init__(self):
    #             self.count = 0
    #             self.children = {}
    #     root = TrieNode()
    #     n = len(words)
    #     nodes = [None] * n
    #     for idx, word in enumerate(words):
    #         for j in range(len(word)):
    #             curr = root
    #             for c in word[j:]:
    #                 if c not in curr.children:
    #                     curr.children[c] = TrieNode()
    #                 curr = curr.children[c]
    #                 curr.count += 1
    #             if nodes[idx] == None:
    #                 nodes[idx] = curr
    #     return [words[i] for i in range(n) if nodes[i].count >= 2]

    # Approach: KMP (Knutt-Morris-Pratt)
    # Time: O(nnm), n=len(words), m=len(words[i])
    # Space: O(m)
    def stringMatching(self, words: List[str]) -> List[str]:
        def getLPS(word: str) -> List[int]:
            m = len(word)
            lps = [0] * m
            p, s = 0, 1
            while s < m:
                if word[s] == word[p]:
                    p += 1
                    lps[s] = p
                    s += 1
                elif p == 0:
                    s += 1
                else:
                    p = lps[p - 1]
            return lps

        n = len(words)

        def check(curr: str, start: int) -> bool:
            m = len(curr)
            lps = getLPS(curr)
            for idx in range(start, n):
                nxt = words[idx]
                nlen = len(nxt)
                if nlen == m:
                    continue
                cid, nid = 0, 0
                while cid < m and nid < nlen:
                    if curr[cid] == nxt[nid]:
                        cid += 1
                        nid += 1
                    elif cid == 0:
                        nid += 1
                    else:
                        cid = lps[cid - 1]
                if cid == m:
                    return True
            return False

        words.sort(key=lambda word: len(word))  # pylint: disable=unnecessary-lambda
        return [w for i, w in enumerate(words) if check(w, i)]

    def test(self):
        for words, expected in [
            (["mass", "as", "hero", "superhero"], ["as", "hero"]),
            (["leetcode", "et", "code"], ["et", "code"]),
            (["blue", "green", "bu"], []),
        ]:
            output = self.stringMatching(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
