from typing import List
import unittest
import bisect

# https://leetcode.com/problems/russian-doll-envelopes/

# python3 -m unittest dp/0354-russian-doll-envelopes.py


class Solution(unittest.TestCase):
    # Approach: Sort + LIS
    # Time: O(n log n), Space: O(n)
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:
        if not envelopes:
            return 0

        # Sort: width asc, height desc
        envelopes.sort(key=lambda x: (x[0], -x[1]))

        # Find LIS on heights
        lis: List[int] = []
        for _, h in envelopes:
            pos = bisect.bisect_left(lis, h)
            if pos == len(lis):
                lis.append(h)
            else:
                lis[pos] = h

        return len(lis)

    def test(self):
        for envelopes, expected in [
            ([[5, 4], [6, 4], [6, 7], [2, 3]], 3),
            ([[1, 1], [1, 1], [1, 1]], 1),
        ]:
            output = self.maxEnvelopes(envelopes)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
