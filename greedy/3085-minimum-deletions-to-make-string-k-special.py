from collections import Counter
import unittest

# https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/

# python3 -m unittest greedy/3085-minimum-deletions-to-make-string-k-special.py

class Solution(unittest.TestCase):
    def minimumDeletions(self, word: str, k: int) -> int:
        freq = list(Counter(word).values())
        res = len(word)
        for low in freq:
            high = low + k
            dels = 0
            for cnt in freq:
                if cnt < low:
                    dels += cnt
                elif cnt > high:
                    dels += cnt - high
            res = min(res, dels)
        return res

    def test(self):
        for word, k, expected in [
            ("aabcaba", 0, 3),
            ("dabdcbdcdcd", 2, 2),
            ("aaabaaa", 2, 1),
        ]:
            output = self.minimumDeletions(word, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
