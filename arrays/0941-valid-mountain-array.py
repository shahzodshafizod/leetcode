from typing import List
import unittest

# https://leetcode.com/problems/valid-mountain-array/

# python3 -m unittest arrays/0941-valid-mountain-array.py


class Solution(unittest.TestCase):
    def validMountainArray(self, arr: List[int]) -> bool:
        n = len(arr)
        if n < 3:  # Recall that arr is a mountain array if and only if: arr.length >= 3
            return False
        # walk up
        idx = 0
        while idx + 1 < n and arr[idx] < arr[idx + 1]:
            idx += 1
        # peak can't be first or last
        if idx == 0 or idx == n - 1:
            return False
        # walk down
        while idx + 1 < n and arr[idx] > arr[idx + 1]:
            idx += 1
        return idx == n - 1

    def test(self):
        for arr, expected in [
            ([2, 1], False),
            ([3, 5, 5], False),
            ([0, 3, 2, 1], True),
            ([0, 1, 2, 3, 4, 5, 6, 7, 8, 9], False),
            ([9, 8, 7, 6, 5, 4, 3, 2, 1, 0], False),
        ]:
            output = self.validMountainArray(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
