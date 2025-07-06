import unittest

# https://leetcode.com/problems/unique-length-3-palindromic-subsequences/

# python3 -m unittest prefixsums/1930-unique-length-3-palindromic-subsequences.py


class Solution(unittest.TestCase):
    def countPalindromicSubsequence(self, s: str) -> int:
        n = len(s)
        right = [0] * n
        mask = 0
        for idx in range(n - 1, -1, -1):
            mask |= 1 << (ord(s[idx]) - ord('a'))
            right[idx] = mask
        masks = [0] * 26
        left = 0
        for idx in range(n - 1):
            key = ord(s[idx]) - ord('a')
            masks[key] |= left & right[idx + 1]
            left |= 1 << key
        return sum(mask.bit_count() for mask in masks)

    def test(self):
        for s, expected in [
            ("aabca", 3),
            ("adc", 0),
            ("bbcbaba", 4),
        ]:
            output = self.countPalindromicSubsequence(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
