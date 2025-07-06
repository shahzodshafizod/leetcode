from typing import List
import unittest

# https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/

# python3 -m unittest intervals/3394-check-if-grid-can-be-cut-into-sections.py


class Solution(unittest.TestCase):
    def checkValidCuts(self, n: int, rectangles: List[List[int]]) -> bool:  # pylint: disable=unused-argument
        def check(dim: int) -> bool:
            rectangles.sort(key=lambda rect: rect[dim])
            count, last = 0, 0
            for rect in rectangles:
                if last <= rect[dim]:
                    count += 1
                    if count == 3:
                        break
                last = max(last, rect[dim + 2])
            return count == 3

        return check(0) or check(1)

    def test(self):
        for n, rectangles, expected in [
            (5, [[1, 0, 5, 2], [0, 2, 2, 4], [3, 2, 5, 3], [0, 4, 4, 5]], True),
            (4, [[0, 0, 1, 1], [2, 0, 3, 4], [0, 2, 2, 3], [3, 0, 4, 3]], True),
            (4, [[0, 2, 2, 4], [1, 0, 3, 2], [2, 2, 3, 4], [3, 0, 4, 2], [3, 2, 4, 4]], False),
        ]:
            output = self.checkValidCuts(n, rectangles)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
