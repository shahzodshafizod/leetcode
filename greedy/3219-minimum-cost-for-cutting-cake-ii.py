from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-for-cutting-cake-ii/

# python3 -m unittest greedy/3219-minimum-cost-for-cutting-cake-ii.py


class Solution(unittest.TestCase):
    def minimumCost(self, m: int, n: int, horizontalCut: List[int], verticalCut: List[int]) -> int:
        horizontalCut.sort(reverse=True)
        verticalCut.sort(reverse=True)
        
        m, n = m-1, n-1
        total, hid, vid = 0, 0, 0
        
        while hid < m and vid < n:
            if horizontalCut[hid] >= verticalCut[vid]:
                total += horizontalCut[hid] * (vid+1)
                hid += 1
            else:
                total += verticalCut[vid] * (hid+1)
                vid += 1

        if hid < m:
            total += sum(horizontalCut[hid:]) * (vid+1)
        elif vid < n:
            total += sum(verticalCut[vid:]) * (hid+1)

        return total

    def test(self):
        for m, n, horizontalCut, verticalCut, expected in [
            (3, 2, [1,3], [5], 13),
            (2, 2, [7], [4], 15),
        ]:
            output = self.minimumCost(m, n, horizontalCut, verticalCut)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
