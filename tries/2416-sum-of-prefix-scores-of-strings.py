from typing import Any, Dict, List
import unittest

# https://leetcode.com/problems/sum-of-prefix-scores-of-strings/

# python3 -m unittest tries/2416-sum-of-prefix-scores-of-strings.py


class Solution(unittest.TestCase):
    # # Approach: Hashmap+Counting
    # # N=len(words)
    # # L=len(words[i])
    # # Time: O(N*L^2)
    # # Space: O(N*L)
    # def sumPrefixScores(self, words: List[str]) -> List[int]:
    #     pref_count = defaultdict(int)
    #     for word in words: # O(N)
    #         for idx in range(len(word)): # O(L)
    #             pref_count[word[:idx+1]] += 1 # O(L)
    #     answer = [0] * len(words)
    #     for i, word in enumerate(words):
    #         for idx in range(len(word)):
    #             answer[i] += pref_count[word[:idx+1]]
    #     return answer

    # Approach: Trie
    # N=len(words)
    # L=len(words[i])
    # Time: O(N*L)
    # Space: O(N*L)
    def sumPrefixScores(self, words: List[str]) -> List[int]:
        root: Dict[str, Any] = {}
        for word in words:
            curr = root
            for c in word:
                if c not in curr:
                    curr[c] = {'#': 0}
                curr = curr[c]
                curr['#'] += 1
        answer = [0] * len(words)
        for idx, word in enumerate(words):
            curr = root
            for c in word:
                curr = curr[c]
                answer[idx] += curr['#']
        return answer

    def test(self):
        for words, expected in [
            (["abc", "ab", "bc", "b"], [5, 4, 3, 2]),
            (["abcd"], [4]),
            (["a", "ab", "abc", "cab"], [3, 5, 6, 3]),
            (["a"], [1]),
            (["a", "b", "aa", "ab"], [3, 1, 4, 4]),
            (["abababab", "abab", "ab", "a", "babababa", "baba", "ba", "b"], [15, 11, 7, 4, 15, 11, 7, 4]),
        ]:
            output = self.sumPrefixScores(words)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
