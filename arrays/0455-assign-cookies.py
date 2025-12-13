from typing import List
import unittest

# https://leetcode.com/problems/assign-cookies/

# python3 -m unittest arrays/0455-assign-cookies.py


class Solution(unittest.TestCase):
    # Approach: Greedy with Sorting
    # Sort both greed factors and cookie sizes.
    # Assign smallest available cookie that satisfies each child.
    # Time: O(n log n + m log m) - sorting
    # Space: O(1) - excluding sort space
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort()
        s.sort()

        child = 0
        cookie = 0

        while child < len(g) and cookie < len(s):
            if s[cookie] >= g[child]:
                child += 1
            cookie += 1

        return child

    def test(self):
        for g, s, expected in [
            ([1, 2, 3], [1, 1], 1),
            ([1, 2], [1, 2, 3], 2),
            ([1, 2, 3], [3], 1),
            ([10, 9, 8, 7], [5, 6, 7, 8], 2),
        ]:
            output = self.findContentChildren(g, s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
