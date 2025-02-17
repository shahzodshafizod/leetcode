from collections import Counter
import unittest

# https://leetcode.com/problems/letter-tile-possibilities/

# python3 -m unittest backtracking/1079-letter-tile-possibilities.py

class Solution(unittest.TestCase):
    def numTilePossibilities(self, tiles: str) -> int:
        counter = Counter(tiles)
        def backtrack() -> int:
            res = 0
            for letter, cnt in counter.items():
                if cnt == 0:
                    continue
                counter[letter] -= 1
                res += 1 + backtrack()
                counter[letter] += 1
            return res
        return backtrack()

    def test(self):
        for tiles, expected in [
            ("AAB", 8),
            ("AAABBC", 188),
            ("V", 1),
        ]:
            output = self.numTilePossibilities(tiles)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
