from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/find-lucky-integer-in-an-array/

# python3 -m unittest hashes/1394-find-lucky-integer-in-an-array.py


class Solution(unittest.TestCase):
    def findLucky(self, arr: List[int]) -> int:
        return max((num for num, cnt in Counter(arr).items() if num == cnt), default=-1)

    def test(self):
        for arr, expected in [
            ([2, 2, 3, 4], 2),
            ([1, 2, 2, 3, 3, 3], 3),
            ([2, 2, 2, 3, 3], -1),
        ]:
            output = self.findLucky(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
