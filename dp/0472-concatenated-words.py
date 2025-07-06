from typing import List
import unittest

# https://leetcode.com/problems/concatenated-words/

# python3 -m unittest dp/0472-concatenated-words.py


class Solution(unittest.TestCase):
    # # Approach #1: Depth-First Search + Trie
    # # Time: O(n*www), n=len(words), w=max(len(words[i]))
    # # Space: O(n*w)
    # def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
    #     class TrieNode:
    #         def __init__(self):
    #             self.children = {}
    #             self.end = False
    #     root = TrieNode()
    #     def insert(word: str) -> None:
    #         curr = root
    #         for c in word:
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode()
    #             curr = curr.children[c]
    #         curr.end = True
    #     max_len = max(len(word) for word in words)
    #     memo = [[None] * max_len for _ in range(len(words))]
    #     def dfs(row: int, col: int) -> bool:
    #         n = len(words[row])
    #         if col == n: return True
    #         if memo[row][col] != None:
    #             return memo[row][col]
    #         curr = root
    #         result = False
    #         for idx in range(col, n):
    #             if words[row][idx] not in curr.children:
    #                 break
    #             curr = curr.children[words[row][idx]]
    #             if curr.end and dfs(row, idx+1):
    #                 result = True
    #                 break
    #         memo[row][col] = result
    #         return result
    #     words.sort(key=len)
    #     concatenated = []
    #     for row, word in enumerate(words):
    #         if dfs(row, 0):
    #             concatenated.append(word)
    #         insert(word)
    #     return concatenated

    # # Approach #2: Depth-First Search + Hash Set
    # # Time: O(n*www), n=len(words), w=max(len(words[i]))
    # # Space: O(n*w)
    # def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
    #     word_set = set(words)
    #     memo = {}
    #     def dfs(word: str) -> bool:
    #         if word in memo: return memo[word]
    #         memo[word] = False
    #         for idx in range(1, len(word)):
    #             prefix = word[:idx]
    #             suffix = word[idx:]
    #             if prefix not in word_set:
    #                 continue
    #             if suffix in word_set or dfs(suffix):
    #                 memo[word] = True
    #                 break
    #         return memo[word]
    #     concatenated = []
    #     for word in words:
    #         if dfs(word):
    #             concatenated.append(word)
    #     return concatenated

    # Approach #3: Bottom-Up Dynamic Programming + Hash Set
    # Time: O(NLogN+N∗M^2), N=len(words), M=max(len(words[i]))
    # Space: O(N+M)
    def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
        words.sort(key=len)
        word_set = set()
        concatenated = []
        for word in words:
            n = len(word)
            dp = [False] * (n + 1)
            dp[0] = True
            for length in range(1, n + 1):
                for j in range(length - 1, -1, -1):
                    if dp[j] and word[j:length] in word_set:
                        dp[length] = True
                        break
            if dp[n]:
                concatenated.append(word)
            word_set.add(word)
        return concatenated

    def test(self):
        for words, expected in [
            (["cat", "dog", "catdog"], ["catdog"]),
            (
                ["cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"],
                ["catsdogcats", "dogcatsdog", "ratcatdogcat"],
            ),
        ]:
            output = self.findAllConcatenatedWordsInADict(words)
            expected.sort()
            output.sort()
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
