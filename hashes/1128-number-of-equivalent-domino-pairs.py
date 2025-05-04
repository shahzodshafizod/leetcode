from typing import List
import unittest

# https://leetcode.com/problems/number-of-equivalent-domino-pairs/

# python3 -m unittest hashes/1128-number-of-equivalent-domino-pairs.py

class Solution(unittest.TestCase):
    def numEquivDominoPairs(self, dominoes: List[List[int]]) -> int:
        count, total = {}, 0
        for a, b in dominoes:
            key = a*10 + b if a <= b else b*10 + a
            count[key] = count.get(key, 0) + 1
            total += count[key] - 1
        return total

    def test(self):
        for dominoes, expected in [
            ([[8,8]], 0),
            ([[1,8]], 0),
            ([[8,1],[1,8]], 1),
            ([[1,2],[2,1],[3,4],[5,6]], 1),
            # ([[1,1],[2,2],[1,1],[1,2],[1,2],[1,1]], 4),
            # ([[1,2],[1,3],[1,4],[1,5],[1,6],[1,7],[1,8],[1,9]], 0),
            # ([[1,8],[8,1],[1,8],[1,8],[1,8],[8,1],[8,2],[8,1],[8,1]], 28),
            # ([[8,8],[8,8],[8,8],[8,8],[8,8],[8,8],[8,8],[8,8],[8,8]], 36),
        ]:
            output = self.numEquivDominoPairs(dominoes)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
