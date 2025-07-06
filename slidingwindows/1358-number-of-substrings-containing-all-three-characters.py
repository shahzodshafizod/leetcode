import unittest

# https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

# python3 -m unittest slidingwindows/1358-number-of-substrings-containing-all-three-characters.py


class Solution(unittest.TestCase):
    def numberOfSubstrings(self, s: str) -> int:
        count = 0
        a, b, c = -1, -1, -1
        for idx, sc in enumerate(s):
            if sc == 'a':
                a = idx
            elif sc == 'b':
                b = idx
            else:
                c = idx
            count += min(a, b, c) + 1
        return count

    def test(self):
        for s, expected in [
            ("abcabc", 10),
            ("aaacb", 3),
            ("abc", 1),
        ]:
            output = self.numberOfSubstrings(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
