import unittest

# https://leetcode.com/problems/vowels-game-in-a-string/

# python3 -m unittest strings/3227-vowels-game-in-a-string.py


class Solution(unittest.TestCase):
    def doesAliceWin(self, s: str) -> bool:
        # vowels = set(['a', 'e', 'i', 'o', 'u'])
        # cnt = sum(1 for c in s if c in vowels)
        # return cnt > 0

        return any(c in "aeiou" for c in s)

    def test(self):
        for s, expected in [
            ("leetcoder", True),
            ("bbcd", False),
        ]:
            output = self.doesAliceWin(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
