from typing import List
import unittest
import math

# https://leetcode.com/problems/construct-the-rectangle/

# python3 -m unittest maths/0492-construct-the-rectangle.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check all divisors
    # # Time: O(n) - check all numbers
    # # Space: O(1)
    # def constructRectangle(self, area: int) -> List[int]:
    #     result: List[int] = []
    #     for width in range(1, area + 1):
    #         if area % width == 0:
    #             length = area // width
    #             if length >= width:
    #                 result = [length, width]
    #     return result

    # Approach #2: Find factors starting from sqrt
    # Time: O(sqrt(n)) - check factors from sqrt down to 1
    # Space: O(1) - constant space
    def constructRectangle(self, area: int) -> List[int]:
        width = int(math.sqrt(area))
        while area % width != 0:
            width -= 1
        length = area // width
        return [length, width]

    def test(self):
        for area, expected in [
            (4, [2, 2]),
            (37, [37, 1]),
            (122122, [427, 286]),
            (1, [1, 1]),
            (10, [5, 2]),
        ]:
            output = self.constructRectangle(area)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
