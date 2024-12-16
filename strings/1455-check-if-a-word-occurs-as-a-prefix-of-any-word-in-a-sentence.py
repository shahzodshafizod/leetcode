import unittest

# https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/

# python3 -m unittest strings/1455-check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence.py

class Solution(unittest.TestCase):
    # # Approach #1: Trie
    # # Time: O(n+m), n=len(sentence), m=len(searchWord)
    # # Space: O(n)
    # def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
    #     class TrieNode:
    #         def __init__(self, index: int):
    #             self.children = {}
    #             self.index = index
    #     root = TrieNode(-1)
    #     for idx, word in enumerate(sentence.split()):
    #         curr = root
    #         for c in word:
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode(idx+1)
    #             curr = curr.children[c]
    #     curr = root
    #     index = -1
    #     for c in searchWord:
    #         if c not in curr.children:
    #             index = -1
    #             break
    #         curr = curr.children[c]
    #         index = curr.index
    #     return index

    # # Approach #2: Brute-Force
    # # Time: O(n), n=len(sentence), m=len(searchWord)
    # # Space: O(n)
    # def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
    #     m = len(searchWord)
    #     index = -1
    #     for idx, word in enumerate(sentence.split()):
    #         if len(word) < m: continue
    #         index = -1
    #         for j in range(m):
    #             if searchWord[j] != word[j]:
    #                 index = -1
    #                 break
    #             index = idx+1
    #         if index != -1:
    #             break
    #     return index

    # Approach #3: Two Pointers
    # Time: O(n), n=len(sentence), m=len(searchWord)
    # Space: O(1)
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        idx, m = 0, len(searchWord)
        words_index = 0
        block = False
        for i in range(len(sentence)):
            if sentence[i] == ' ':
                words_index += 1
                idx = 0
                block = False
                continue
            if block:
                continue
            if searchWord[idx] != sentence[i]:
                block = True
                continue
            idx += 1
            if idx == m:
                return words_index+1
        return -1

    def test(self) -> None:
        for sentence, searchWord, expected in [
            ("love", "lov", 1),
            ("i am tired", "you", -1),
            ("helloh hellohe", "hellohe", 2),
            ("i love eating burger", "burg", 4),
            ("hello from the other side", "they", -1),
            ("hellohello hellohellohello", "ello", -1),
            ("i love eating broadburgers", "burg", -1),
            ("this problem is an easy problem", "pro", 2),
            ("love i love eating bunburger burger", "i", 2),
            ("i love eating bunburger burger", "burg", 5),
            # ("winstontang love they i pillow jonathan you winstontang pillow dream you duck", "p", 5),
        ]:
            output = self.isPrefixOfWord(sentence, searchWord)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
