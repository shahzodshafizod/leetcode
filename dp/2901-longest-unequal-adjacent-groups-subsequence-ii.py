from typing import List
import unittest

# https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/

# python3 -m unittest dp/2901-longest-unequal-adjacent-groups-subsequence-ii.py

class Solution(unittest.TestCase):
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        def ham_dist(i: int, j: int) -> bool:
            return sum(a != b for a, b in zip(words[i], words[j]))
        n = len(words)
        dp = [0] * n
        next = [-1] * n
        best = n-1
        for i in range(n-1,-1,-1):
            for j in range(i+1,n):
                if (
                    dp[j]+1 > dp[i]
                    and groups[i] != groups[j]
                    and len(words[i]) == len(words[j])
                    and ham_dist(i, j) == 1
                ):
                    dp[i] = dp[j] + 1
                    next[i] = j
            if dp[i] > dp[best]:
                best = i
        subsequence = []
        while best != -1:
            subsequence.append(words[best])
            best = next[best]
        return subsequence

    def test(self):
        for words, groups, expected in [
            (["bab","dab","cab"], [1,2,2], ["bab","dab"]),
            (["a","b","c","d"], [1,2,3,4], ["a","b","c","d"]),
            (["bdb","aaa","ada"], [2,1,3], ["aaa","ada"]),
            # (["bad","dc","bc","ccd","dd","da","cad","dba","aba"], [9,7,1,2,6,8,3,7,2], ["dc","dd","da"]),
        ]:
            output = self.getWordsInLongestSubsequence(words, groups)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
