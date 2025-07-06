from collections import defaultdict, deque
from typing import List
import unittest

# https://leetcode.com/problems/word-ladder/

# python3 -m unittest graphs/0127-word-ladder.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # N = len(wordList)
    # M = len(wordList[i])
    # Time: O(N * M^2)
    # Space: O(N*M)
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        if endWord not in wordList:
            return 0
        adjList = defaultdict(list)
        word_len = len(beginWord)
        for word in wordList:
            for idx in range(word_len):
                pattern = word[:idx] + '*' + word[idx + 1 :]
                adjList[pattern].append(word)
        queue = deque([beginWord])
        visited = set([beginWord])
        steps = 1
        while queue:
            for _ in range(len(queue)):
                word = queue.popleft()
                if word == endWord:
                    return steps
                for idx in range(word_len):
                    pattern = word[:idx] + '*' + word[idx + 1 :]
                    for nxt in adjList[pattern]:
                        if nxt not in visited:
                            visited.add(nxt)
                            queue.append(nxt)
            steps += 1
        return 0

    def test(self):
        for beginWord, endWord, wordList, expected in [
            ("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"], 5),
            ("hit", "cog", ["hot", "dot", "dog", "lot", "log"], 0),
            ("a", "c", ["a", "b", "c"], 2),
            ("leet", "code", ["lest", "leet", "lose", "code", "lode", "robe", "lost"], 6),
            ("hot", "dog", ["hot", "dog", "dot"], 3),
            ("hot", "dog", ["hot", "dog", "cog", "pot", "dot"], 3),
            ("hot", "dot", ["hot", "dot", "dog"], 2),
            ("hot", "dog", ["hot", "dog"], 0),
            (
                "hbo",
                "qbx",
                ["abo", "hco", "hbw", "ado", "abq", "hcd", "hcj", "hww", "qbq", "qby", "qbz", "qbx", "qbw"],
                4,
            ),
            ("ab", "ac", ["xx", "ee", "ac"], 2),
            ("be", "ko", ["ce", "mo", "ko", "me", "co"], 4),
        ]:
            output = self.ladderLength(beginWord, endWord, wordList)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
