import unittest

# https://leetcode.com/problems/sort-vowels-in-a-string/

# python3 -m unittest strings/2785-sort-vowels-in-a-string.py


class Solution(unittest.TestCase):
    def sortVowels(self, s: str) -> str:
        indices = {'A': 0, 'E': 1, 'I': 2, 'O': 3, 'U': 4, 'a': 5, 'e': 6, 'i': 7, 'o': 8, 'u': 9}
        vowels = ['A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u']
        rs = list(s)
        freq = [0] * 11
        for c in rs:
            freq[indices.get(c, 10)] += 1
        idx = 0
        for i, c in enumerate(rs):
            if c in indices:
                while freq[idx] == 0:
                    idx += 1
                rs[i] = vowels[idx]
                freq[idx] -= 1
        return "".join(rs)

    def test(self):
        for s, expected in [
            ("lEetcOde", "lEOtcede"),
            ("lYmpH", "lYmpH"),
        ]:
            output = self.sortVowels(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
