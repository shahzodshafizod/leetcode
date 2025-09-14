from typing import List, Set, Dict
import unittest

# https://leetcode.com/problems/vowel-spellchecker/

# python3 -m unittest hashes/0966-vowel-spellchecker.py


class Solution(unittest.TestCase):
    # # Approach #1: Tries
    # # Time: O(n*m)
    # # Space: O(n*m)
    # def spellchecker(self, wordlist: List[str], queries: List[str]) -> List[str]:
    #     class Trie:
    #         def __init__(self):
    #             self.root: Dict[str, Any] = {}

    #         def insert(self, word: str, orig: str):
    #             curr = self.root
    #             for c in word:
    #                 if c not in curr:
    #                     curr[c] = {}
    #                 curr = curr[c]
    #             curr['$'] = orig

    #         def find(self, word: str) -> str:
    #             curr = self.root
    #             for c in word:
    #                 if c not in curr:
    #                     return ""
    #                 curr = curr[c]
    #             return curr.get('$', "")

    #     vowels = set(['a', 'e', 'i', 'o', 'u'])

    #     def devowel(word: str) -> str:
    #         return ''.join('*' if c in vowels else c for c in word.lower())

    #     indices = {word: idx for idx, word in enumerate(wordlist)}
    #     indices[''] = len(wordlist)

    #     orig, lower, vowel = Trie(), Trie(), Trie()
    #     for word in wordlist[::-1]:
    #         orig.insert(word, word)
    #         lower.insert(word.lower(), word)
    #         vowel.insert(devowel(word), word)

    #     answer = [''] * len(queries)
    #     for i, word in enumerate(queries):
    #         for find, w in [(orig.find, word), (lower.find, word.lower()), (vowel.find, devowel(word))]:
    #             ans = find(w)
    #             if ans:
    #                 answer[i] = ans
    #                 break

    #     return answer

    # Approach #2: Hash Table
    # Time: O(n)
    # Space: O(n)
    def spellchecker(self, wordlist: List[str], queries: List[str]) -> List[str]:
        vowels = set(['a', 'e', 'i', 'o', 'u'])

        def devowel(word: str):
            return ''.join('*' if c in vowels else c for c in word.lower())

        orig: Set[str] = set()
        lower: Dict[str, str] = {}
        vowel: Dict[str, str] = {}
        for word in wordlist[::-1]:
            orig.add(word)
            lower[word.lower()] = word
            vowel[devowel(word)] = word

        answer: List[str] = [''] * len(queries)
        for i, word in enumerate(queries):
            if word in orig:
                answer[i] = word
            elif word.lower() in lower:
                answer[i] = lower[word.lower()]
            elif devowel(word) in vowel:
                answer[i] = vowel[devowel(word)]

        return answer

    def test(self):
        for wordlist, queries, expected in [
            (
                ["KiTe", "kite", "hare", "Hare"],
                ["kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"],
                ["kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"],
            ),
            (["yellow"], ["YellOw"], ["yellow"]),
        ]:
            output = self.spellchecker(wordlist, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
