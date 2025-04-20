from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/rabbits-in-forest/

# python3 -m unittest greedy/0781-rabbits-in-forest.py

class Solution(unittest.TestCase):
    def numRabbits(self, answers: List[int]) -> int:
        return sum((total+others)//(others+1)*(others+1) for others, total in Counter(answers).items())

    def test(self):
        for answers, expected in [
            ([1,1,2], 5),
            ([10,10,10], 11),
            ([0,0,1,1,1], 6),
            ([24,24,24,24,2,2,1,2,8,8,2,8,8], 42),
            ([2,1,0,0,0,0,0], 10),
            ([5,1,5,3,5,2,12,4,20,20], 54),
            ([33,25,22,2,39,15,50,7,5,12,12,12,12,3,4], 229),
        ]:
            output = self.numRabbits(answers)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
