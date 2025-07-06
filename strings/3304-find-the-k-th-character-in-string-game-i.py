import unittest

# https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/

# python3 -m unittest strings/3304-find-the-k-th-character-in-string-game-i.py


class Solution(unittest.TestCase):
    def kthCharacter(self, k: int) -> str:
        word, length = [0], 1
        while length < k:
            for i in range(length):
                word.append((word[i] + 1) % 26)
            length *= 2
        return chr(ord('a') + word[k - 1])

    def test(self):
        for k, expected in [
            (5, "b"),
            (10, "c"),
        ]:
            output = self.kthCharacter(k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
