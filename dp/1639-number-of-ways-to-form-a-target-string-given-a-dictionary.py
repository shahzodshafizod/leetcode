from typing import List
import unittest

# https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/

# python3 -m unittest dp/1639-number-of-ways-to-form-a-target-string-given-a-dictionary.py


class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
    # # Space: O(W*T)
    # def numWays(self, words: List[str], target: str) -> int:
    #     MOD = int(1e9+7)
    #     wlen, tlen = len(words[0]), len(target)
    #     count = [[0] * 26 for _ in range(wlen)]
    #     for word in words:
    #         for idx in range(wlen):
    #             count[idx][ord(word[idx])-ord('a')] += 1
    #     memo = [[None] * tlen for _ in range(wlen)]
    #     def dp(wid: int, tid: int) -> int:
    #         if tid == tlen: return 1
    #         if wid == wlen: return 0
    #         if memo[wid][tid] != None:
    #             return memo[wid][tid]
    #         # decision to skip
    #         ways = dp(wid+1, tid)
    #         # decision to include
    #         if count[wid][ord(target[tid])-ord('a')] > 0:
    #             ways = (ways +
    #                     count[wid][ord(target[tid])-ord('a')] * dp(wid+1, tid+1)
    #             ) % MOD
    #         memo[wid][tid] = ways
    #         return ways
    #     return dp(0, 0)

    # # Approach: Bottom-Up Dynamic Programming
    # # Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
    # # Space: O(W*T)
    # def numWays(self, words: List[str], target: str) -> int:
    #     MOD = int(1e9+7)
    #     tlen, wlen = len(target), len(words[0])
    #     count = [[0] * 26 for _ in range(wlen)]
    #     for word in words:
    #         for wid in range(wlen):
    #             count[wid][ord(word[wid])-ord('a')] += 1
    #     dp = [[0]*tlen + [1] for _ in range (wlen+1)]
    #     for wid in range(wlen-1, -1, -1):
    #         for tid in range(tlen-1, -1, -1):
    #             dp[wid][tid] = (
    #                 dp[wid+1][tid] +
    #                 count[wid][ord(target[tid])-ord('a')] * dp[wid+1][tid+1]
    #             ) % MOD
    #     return dp[0][0]

    # Approach: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
    # Space: O(T)
    def numWays(self, words: List[str], target: str) -> int:
        MOD = int(1e9 + 7)
        tlen, wlen = len(target), len(words[0])
        count = [[0] * 26 for _ in range(wlen)]
        for word in words:
            for wid in range(wlen):
                count[wid][ord(word[wid]) - ord('a')] += 1
        curr, nxt = [0] * tlen + [1], [0] * tlen + [1]
        for wid in range(wlen - 1, -1, -1):
            curr, nxt = nxt, curr
            for tid in range(tlen - 1, -1, -1):
                curr[tid] = (nxt[tid] + count[wid][ord(target[tid]) - ord('a')] * nxt[tid + 1]) % MOD
        return curr[0]

    def test(self):
        for words, target, expected in [
            (["acca", "bbbb", "caca"], "aba", 6),
            (["abba", "baab"], "bab", 4),
            # (["cbabddddbc","addbaacbbd","cccbacdccd","cdcaccacac","dddbacabbd","bdbdadbccb","ddadbacddd","bbccdddadd","dcabaccbbd","ddddcddadc","bdcaaaabdd","adacdcdcdd","cbaaadbdbb","bccbabcbab","accbdccadd","dcccaaddbc","cccccacabd","acacdbcbbc","dbbdbaccca","bdbddbddda","daabadbacb","baccdbaada","ccbabaabcb","dcaabccbbb","bcadddaacc","acddbbdccb","adbddbadab","dbbcdcbcdd","ddbabbadbb","bccbcbbbab","dabbbdbbcb","dacdabadbb","addcbbabab","bcbbccadda","abbcacadac","ccdadcaada","bcacdbccdb"], "bcbbcccc", 677452090),
        ]:
            output = self.numWays(words, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
