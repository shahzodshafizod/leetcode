from typing import List
import unittest

# https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/

# python3 -m unittest greedy/1007-minimum-domino-rotations-for-equal-row.py

class Solution(unittest.TestCase):
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        # If one side can be made equal,
        # another side can be made equal with at least the same effort,
        # not more optimal!!!
        # So the first found one is the optimal one.
        n = len(tops)
        for target in [tops[0], bottoms[0]]:
            t_swaps, b_swaps = 0, 0
            for idx, (top, bottom) in enumerate(zip(tops, bottoms)):
                if top != target and bottom != target:
                    break
                elif top != target:
                    t_swaps += 1
                elif bottom != target:
                    b_swaps += 1
                if idx == n-1:
                    return min(t_swaps, b_swaps)
        return -1

    def test(self):
        for tops, bottoms, expected in [
            ([2,1,2,4,2,2], [5,2,6,2,3,2], 2),
            ([3,5,1,2,3], [3,6,3,3,4], -1),
            ([1,1,1,2,2,2], [5,5,5,1,1,1], 3),
        ]:
            output = self.minDominoRotations(tops, bottoms)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
