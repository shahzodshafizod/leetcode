import unittest
from typing import List

# https://leetcode.com/problems/prefix-and-suffix-search/

# python3 -m unittest design/0745-prefix-and-suffix-search.py

class TrieNode:
    def __init__(self):
        self.index = -1
        self.children = {}

class Trie:
    def __init__(self):
        self.root = TrieNode()
    
    def insert(self, word: str, index: int) -> None:
        curr = self.root
        for c in word:
            if c not in curr.children:
                curr.children[c] = TrieNode()
            curr = curr.children[c]
            curr.index = index

    def find(self, word: str) -> int:
        curr = self.root
        for c in word:
            if c not in curr.children: return -1
            curr = curr.children[c]
        return curr.index

class WordFilter:

    def __init__(self, words: List[str]):
        self.trie = Trie()
        for index, word in enumerate(words):
            word_len = len(word)
            word = word + "#" + word
            for i in range(word_len):
                self.trie.insert(word[i:], index)

    def f(self, pref: str, suff: str) -> int:
        word = suff + "#" + pref
        return self.trie.find(word)

class TestWordFilter(unittest.TestCase):
    def test(self) -> None:
        for commands, values, expected in [
            (["WordFilter", "f", "f", "f"], [["apple", "app", "ape"], ["a", "e"], ["appl", "pple"], ["apple", "apple"]], [None, 2, 0, 0]),
            (["WordFilter","f"], [["abbba","abba"],["ab","ba"]], [None,1]),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "WordFilter":
                        filter = WordFilter(values[idx])
                    case "f":
                        output = filter.f(values[idx][0], values[idx][1])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your WordFilter object will be instantiated and called as such:
# obj = WordFilter(words)
# param_1 = obj.f(pref,suff)
