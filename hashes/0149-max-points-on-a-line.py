from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/max-points-on-a-line/

# python3 -m unittest hashes/0149-max-points-on-a-line.py

class Solution(unittest.TestCase):
    def maxPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        if n <= 2:
            return n
        count = 1
        for i in range(n):
            x1, y1 = points[i]
            slopes = defaultdict(int)
            for j in range(i+1, n):
                x2, y2 = points[j]
                # slope for any two points = slope= (y2-y1)/(x2-x1)
                slope = float("inf") if x1 == x2 else (y2-y1) / (x2-x1)
                slopes[slope] += 1
                count = max(count, slopes[slope]+1)
        return count

    def test(self) -> None:
        for points, expected in [
            ([[1,1],[2,2],[3,3]], 3),
            ([[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]], 4),
            ([[0,0]], 1),
            ([[0,0],[94911151,94911150],[94911152,94911151]], 3),
            ([[0,0],[1,1],[0,0]], 2),
            ([[0,0],[0,0]], 2),
            ([[0,0],[1073741822,2147483645],[1073741823,2147483647]], 3),
            ([[1,1],[1,1],[2,3]], 2),
        ]:
            output = self.maxPoints(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
