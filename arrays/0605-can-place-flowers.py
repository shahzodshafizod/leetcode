from typing import List
import unittest

# https://leetcode.com/problems/can-place-flowers/

# python3 -m unittest arrays/0605-can-place-flowers.py


class Solution(unittest.TestCase):
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        flowerbed = [0] + flowerbed + [0]
        for idx in range(1, len(flowerbed) - 1):
            if n == 0:
                break
            if sum(flowerbed[idx - 1 : idx + 2]) == 0:
                flowerbed[idx] = 1
                n -= 1
        return n == 0

    def testCanPlaceFlowers(self) -> None:
        for flowerbed, n, expected in [
            ([1, 0, 0, 0, 1], 1, True),
            ([1, 0, 0, 0, 1], 2, False),
            ([0, 0, 1, 0, 0], 1, True),
            ([0, 0, 1, 0, 0], 2, True),
            ([0, 0, 1, 0, 0], 3, False),
            ([1, 0, 0, 0, 1], 1, True),
            ([1, 0, 0, 0, 1], 2, False),
            ([1, 0, 0, 0, 1], 1, True),
            ([0, 0, 1, 0, 0], 1, True),
            ([0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0], 4, True),
            ([1, 0, 0, 0, 0, 0, 1], 2, True),
            ([0], 1, True),
            ([1, 0, 0, 0, 0, 1], 2, False),
            ([1, 0, 0, 1, 0, 0, 0, 1], 1, True),
        ]:
            output = self.canPlaceFlowers(flowerbed, n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
