from typing import List
import unittest

# https://leetcode.com/problems/count-good-triplets/

# python3 -m unittest arrays/1534-count-good-triplets.py


class Solution(unittest.TestCase):
    # Approach: Enumeration
    # Time: O(nnn)
    # Space: O(1)
    def countGoodTriplets(self, arr: List[int], a: int, b: int, c: int) -> int:
        return sum(
            abs(arr[i] - arr[j]) <= a and abs(arr[j] - arr[k]) <= b and abs(arr[k] - arr[i]) <= c
            for i in range(len(arr))
            for j in range(i + 1, len(arr))
            for k in range(j + 1, len(arr))
        )

    def test(self):
        for arr, a, b, c, expected in [
            ([3, 0, 1, 1, 9, 7], 7, 2, 3, 4),
            ([1, 1, 2, 2, 3], 0, 0, 1, 0),
        ]:
            output = self.countGoodTriplets(arr, a, b, c)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
