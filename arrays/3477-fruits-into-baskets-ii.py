from typing import List
import unittest

# https://leetcode.com/problems/fruits-into-baskets-ii/

# python3 -m unittest arrays/3477-fruits-into-baskets-ii.py


class Solution(unittest.TestCase):
    def numOfUnplacedFruits(self, fruits: List[int], baskets: List[int]) -> int:
        unplaced, n = 0, len(baskets)
        for fruit in fruits:
            unplaced += 1
            for idx in range(n):
                if baskets[idx] >= fruit:
                    baskets[idx] = 0
                    unplaced -= 1
                    break
        return unplaced

    def test(self):
        for fruits, baskets, expected in [
            ([4, 2, 5], [3, 5, 4], 1),
            ([3, 6, 1], [6, 4, 7], 0),
        ]:
            output = self.numOfUnplacedFruits(fruits, baskets)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
