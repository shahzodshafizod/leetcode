import unittest

# https://leetcode.com/problems/maximum-score-after-splitting-a-string/

# python3 -m unittest prefixsums/1422-maximum-score-after-splitting-a-string.py


class Solution(unittest.TestCase):
    def maxScore(self, s: str) -> int:
        zeroes, ones = 0, s.count('1')
        score = 0
        for idx in range(len(s) - 1):
            if s[idx] == '0':
                zeroes += 1
            else:
                ones -= 1
            score = max(score, zeroes + ones)
        return score

    def test(self):
        for s, expected in [
            ("00", 1),
            ("01", 2),
            ("010", 2),
            ("1111", 3),
            ("0010", 3),
            ("00111", 5),
            ("01101", 4),
            ("011101", 5),
            ("1111000", 3),
        ]:
            output = self.maxScore(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
