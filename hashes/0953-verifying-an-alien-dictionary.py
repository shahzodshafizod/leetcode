from typing import List
import unittest

# https://leetcode.com/problems/verifying-an-alien-dictionary/

# python3 -m unittest hashes/0953-verifying-an-alien-dictionary.py


class Solution(unittest.TestCase):
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        # order = {letter:index for index, letter in enumerate(order)}
        # words = [[order[c] for c in word] for word in words]
        # return all(list1 <= list2 for list1, list2 in zip(words, words[1:]))
        order = {letter: index for index, letter in enumerate(order)}
        for i in range(1, len(words)):
            is_sorted = False
            n1, n2 = len(words[i - 1]), len(words[i])
            for j in range(min(n1, n2)):
                if order[words[i - 1][j]] > order[words[i][j]]:
                    return False
                if order[words[i - 1][j]] < order[words[i][j]]:
                    is_sorted = True
                    break
            if not is_sorted and n1 > n2:
                return False
        return True

    def test(self) -> None:
        for words, order, expected in [
            (["hello", "leetcode"], "hlabcdefgijkmnopqrstuvwxyz", True),
            (["word", "world", "row"], "worldabcefghijkmnpqstuvxyz", False),
            (["apple", "app"], "abcdefghijklmnopqrstuvwxyz", False),
            (["apple", "w"], "abcdefghijklmnopqrstuvwxyz", True),
        ]:
            output = self.isAlienSorted(words, order)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
