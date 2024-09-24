from collections import defaultdict
from typing import List, Optional
import unittest

# https://leetcode.com/problems/longest-common-suffix-queries/

# python3 -m unittest tries/3093-longest-common-suffix-queries.py

class Solution(unittest.TestCase):
    # M=len(wordsContainer)
    # C=len(wordsContainer[i])
    # N=len(wordsQuery)
    # Q=len(wordsQuery[i])
    # Time: O(M*C + N*Q)
    # Space: O(M*C)
    def stringIndices(self, wordsContainer: List[str], wordsQuery: List[str]) -> List[int]:
        class Trie:
            def __init__(self) -> None:
                self.children = defaultdict(Trie)
                self.index = -1
                self.length = 5001
        def buildTrie() -> Optional[Trie]:
            root = Trie()
            for idx, word in enumerate(wordsContainer):
                curr = root
                length = len(word)
                if length < curr.length:
                    curr.length = length
                    curr.index = idx
                for c in word[::-1]:
                    if c not in curr.children:
                        curr.children[c] = Trie()
                    curr = curr.children[c]
                    if length < curr.length:
                        curr.length = length
                        curr.index = idx
            return root
        root = buildTrie()
        ans = [0] * len(wordsQuery)
        for idx, word in enumerate(wordsQuery):
            curr = root
            ans[idx] = curr.index
            for c in word[::-1]:
                if c not in curr.children:
                    break
                curr = curr.children[c]
                ans[idx] = curr.index
        return ans

    def testStringIndices(self) -> None:
        for wordsContainer, wordsQuery, expected in [
            (["abcd","bcd","xbcd"], ["cd","bcd","xyz"], [1,1,1]),
            (["abcdefgh","poiuygh","ghghgh"], ["gh","acbfgh","acbfegh"], [2,0,2]),
            (["abcde", "abcde"], ["abcde", "bcde", "cde", "de", "e"], [0, 0, 0, 0, 0]),
		    (["abcd", "bcda"], ["cdf", "afa"], [0, 1]),
        ]:
            output = self.stringIndices(wordsContainer, wordsQuery)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
