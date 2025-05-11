from typing import List
import unittest

# https://leetcode.com/problems/three-consecutive-odds/

# python3 -m unittest arrays/1550-three-consecutive-odds.py

class Solution(unittest.TestCase):
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        count = 0
        for num in arr:
            count += 1 if num&1 else -count
            if count == 3:
                return True
        return False

    def test(self):
        for arr, expected in [
            ([2,6,4,1], False),
            ([1,2,34,3,4,5,7,23,12], True),
            ([1,3,4,6,7,45], False),
        ]:
            output = self.threeConsecutiveOdds(arr)
            self.assertEqual(output, expected, f"output: {output}, expected: {expected}")
