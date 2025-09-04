from typing import List
import unittest

# https://leetcode.com/problems/find-the-number-of-ways-to-place-people-ii/

# python3 -m unittest maths/3027-find-the-number-of-ways-to-place-people-ii.py


class Solution(unittest.TestCase):
    def numberOfPairs(self, points: List[List[int]]) -> int:
        points.sort(key=lambda x: (x[0], -x[1]))
        cnt, n = 0, len(points)
        for i in range(n - 1):
            mxy, mny = points[i][1], -1e9 - 1
            for j in range(i + 1, n):
                _, y = points[j]
                if mny < y <= mxy:
                    cnt += 1
                    mny = y
        return cnt

    def test(self):
        for points, expected in [
            ([[1, 1], [2, 2], [3, 3]], 0),
            ([[6, 2], [4, 4], [2, 6]], 2),
            ([[3, 1], [1, 3], [1, 1]], 2),
            ([[1, 1], [2, -1000000000]], 1),
        ]:
            output = self.numberOfPairs(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
