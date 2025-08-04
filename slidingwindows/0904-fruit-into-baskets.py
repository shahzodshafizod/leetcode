from typing import List
import unittest

# https://leetcode.com/problems/fruit-into-baskets/

# python3 -m unittest slidingwindows/0904-fruit-into-baskets.py


class Solution(unittest.TestCase):
    def totalFruit(self, fruits: List[int]) -> int:
        a, b = 0, 0
        totallength, lenb = 0, 0
        res = 0
        for c in fruits:
            if c in (a, b):
                totallength += 1
            else:
                totallength = lenb + 1
            if c != b:
                a, b = b, c
                lenb = 0
            lenb += 1
            res = max(res, totallength)
        return res

    def test(self):
        for fruits, expected in [
            ([1, 2, 1], 3),
            ([0, 1, 2, 2], 3),
            ([1, 2, 3, 2, 2], 4),
        ]:
            output = self.totalFruit(fruits)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
