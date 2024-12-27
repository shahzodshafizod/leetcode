from typing import List
import unittest

# https://leetcode.com/problems/best-sightseeing-pair/

# python3 -m unittest dp/1014-best-sightseeing-pair.py

class Solution(unittest.TestCase):
    def maxScoreSightseeingPair(self, values: List[int]) -> int:
        score = 0
        n = len(values)
        maxval = values[n-1] - n+1
        for idx in range(n-2, -1, -1):
            score = max(score, values[idx] + idx + maxval)
            maxval = max(maxval, values[idx]-idx)
        return score

    def test(self):
        for values, expected in [
            ([8,1,5,2,6], 11),
            ([1,2], 2),
        ]:
            output = self.maxScoreSightseeingPair(values)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
