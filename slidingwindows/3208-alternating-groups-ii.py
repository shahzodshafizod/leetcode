from typing import List
import unittest

# https://leetcode.com/problems/alternating-groups-ii/

# python3 -m unittest slidingwindows/3208-alternating-groups-ii.py

class Solution(unittest.TestCase):
    # Time: O(N)
    # Space: O(1)
    def numberOfAlternatingGroups(self, colors: List[int], k: int) -> int:
        n = len(colors)
        groups, tails = 0, 1
        prev = 0
        for curr in range(1, n+k-1):
            curr %= n
            if colors[prev] == colors[curr]:
                tails = 0
            tails += 1
            if tails >= k:
                groups += 1
            prev = curr
        return groups

    def test(self):
        for colors, k, expected in [
            ([0,1,0,1,0], 3, 3),
            ([0,1,0,0,1,0,1], 6, 2),
            ([1,1,0,1], 4, 0),
        ]:
            output = self.numberOfAlternatingGroups(colors, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
