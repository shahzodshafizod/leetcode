from typing import List
import unittest

# https://leetcode.com/problems/triangle/

# python3 -m unittest dp/0120-triangle.py


class Solution(unittest.TestCase):
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        dp = [0] * (len(triangle) + 1)
        for row in triangle[::-1]:
            for col, num in enumerate(row):
                dp[col] = num + min(dp[col], dp[col+1])
        return dp[0]

    def test(self):
        for triangle, expected in [
            ([[2],[3,4],[6,5,7],[4,1,8,3]], 11),
            ([[-10]], -10),
        ]:
            output = self.minimumTotal(triangle)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
