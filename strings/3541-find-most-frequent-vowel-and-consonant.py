from collections import Counter
import unittest

# https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

# python3 -m unittest strings/3541-find-most-frequent-vowel-and-consonant.py


class Solution(unittest.TestCase):
    def maxFreqSum(self, s: str) -> int:
        freq = Counter(s)
        vowels = set(['a', 'e', 'i', 'o', 'u'])
        vmx = max((f for c, f in freq.items() if c in vowels), default=0)
        cmx = max((f for c, f in freq.items() if c not in vowels), default=0)
        return vmx + cmx

    def test(self):
        for s, expected in [
            ("successes", 6),
            ("aeiaeia", 3),
        ]:
            output = self.maxFreqSum(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
