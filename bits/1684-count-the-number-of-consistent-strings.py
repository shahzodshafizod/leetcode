from typing import List
import unittest

# https://leetcode.com/problems/count-the-number-of-consistent-strings/

# python3 -m unittest bits/1684-count-the-number-of-consistent-strings.py

class Solution(unittest.TestCase):
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        mask = 0
        for c in allowed:
            mask |= 1 << (ord(c)-ord('a'))
        count = 0
        for word in words:
            count += 1
            for c in word:
                if (1 << (ord(c)-ord('a')))&mask == 0:
                    count -= 1
                    break
        return count

    def testCountConsistentStrings(self) -> None:
        for allowed, words, expected in [
            ("ab", ["ad","bd","aaab","baa","badab"], 2),
            ("abc", ["a","b","c","ab","ac","bc","abc"], 7),
            ("cad", ["cc","acd","b","ba","bac","bad","ac","d"], 4),
            ("xyz", ["x","xy","xyz","z","y","xxy","xxyzz"], 7),
            ("a", ["a","b","c","aa","aaa","ab","ac"], 3),
        ]:
            output = self.countConsistentStrings(allowed, words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
