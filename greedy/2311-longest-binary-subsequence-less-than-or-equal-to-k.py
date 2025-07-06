import unittest

# https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/

# python3 -m unittest greedy/2311-longest-binary-subsequence-less-than-or-equal-to-k.py


class Solution(unittest.TestCase):
    def longestSubsequence(self, s: str, k: int) -> int:
        count, n = 0, len(s)
        for i in range(n):
            digit = int(s[n - 1 - i])
            if k >= (digit << i):
                k -= digit << i
                count += 1
        return count

    def test(self):
        for s, k, expected in [
            ("0", 1, 1),
            ("1001010", 5, 5),
            ("00101001", 1, 6),
            ("101011010", 15, 6),
            ("1111111111", 1, 1),
            ("0001101010", 5, 7),
            ("11001101000111010111", 15, 11),
            ("00101000001101011111111110110111101011010011100100", 5, 22),
        ]:
            output = self.longestSubsequence(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
