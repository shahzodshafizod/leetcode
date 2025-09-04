from typing import List
import unittest

# https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/

# python3 -m unittest maths/3025-find-the-number-of-ways-to-place-people-i.py


class Solution(unittest.TestCase):
    def numberOfPairs(self, points: List[List[int]]) -> int:
        cnt, n = 0, len(points)
        for i in range(n):
            xi, yi = points[i]
            for j in range(n):
                xj, yj = points[j]
                # we ensure points[i] is on the upper left side of points[j]
                if i == j or xi > xj or yi < yj:
                    continue
                cnt += 1
                for k in range(n):
                    if k in (i, j):
                        continue
                    xk, yk = points[k]
                    if xi <= xk <= xj and yi >= yk >= yj:
                        cnt -= 1
                        break
        return cnt

    def test(self):
        for points, expected in [
            ([[1, 1], [2, 2], [3, 3]], 0),
            ([[6, 2], [4, 4], [2, 6]], 2),
            ([[3, 1], [1, 3], [1, 1]], 2),
        ]:
            output = self.numberOfPairs(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
