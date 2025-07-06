from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

# python3 -m unittest hashes/2131-longest-palindrome-by-concatenating-two-letter-words.py


class Solution(unittest.TestCase):
    def longestPalindrome(self, words: List[str]) -> int:
        counter = Counter(words)
        length, center = 0, 0
        for word, cnt in counter.items():
            revw = word[::-1]
            if word == revw:
                length += cnt // 2 * 4
                if cnt & 1:
                    center = 2
            else:
                length += min(cnt, counter[revw]) * 2
        return length + center

    def test(self):
        for words, expected in [
            (["lc", "cl", "gg"], 6),
            (["cc", "ll", "xx"], 2),
            (["lc", "cl", "gg", "ak", "kk"], 6),
            (["ab", "ty", "yt", "lc", "cl", "ab"], 8),
            (["ab", "ty", "yt", "lc", "cl", "ab", "ba"], 12),
            (["dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"], 22),
            (["qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"], 14),
        ]:
            output = self.longestPalindrome(words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
