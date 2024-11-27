from typing import List
import unittest

# https://leetcode.com/problems/check-if-n-and-its-double-exist/

# python3 -m unittest hashes/1346-check-if-n-and-its-double-exist.py

class Solution(unittest.TestCase):
    def checkIfExist(self, arr: List[int]) -> bool:
        seen = set()
        for num in arr:
            if num*2 in seen or num/2 in seen:
                return True
            seen.add(num)
        return False

    def testCheckIfExist(self) -> None:
        for arr, expected in [
            ([10,2,5,3], True),
            ([7,1,14,11], True),
            ([3,1,7,11], False),
            ([-2,0,10,-19,4,6,-8], False),
            ([0], False),
            ([0,0], True),
        ]:
            output = self.checkIfExist(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
