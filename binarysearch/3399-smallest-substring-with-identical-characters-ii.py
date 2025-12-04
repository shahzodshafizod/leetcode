from bisect import bisect_left
from itertools import groupby
import unittest

# https://leetcode.com/problems/smallest-substring-with-identical-characters-ii/

# python3 -m unittest binarysearch/3399-smallest-substring-with-identical-characters-ii.py


class Solution(unittest.TestCase):
    # Approach: Binary Search + Greedy Verification
    # Time: O(n * log n)
    # Space: O(n)
    def minLength(self, s: str, numOps: int) -> int:
        n = len(s)

        # check k=1: 010101... or 101010...
        cnt = sum(s[i] == str(i % 2) for i in range(n))
        if min(cnt, n - cnt) <= numOps:
            return 1

        lens = [len(list(g)) for _, g in groupby(s)]
        return bisect_left(range(2, max(lens) + 1), True, key=lambda k: sum(l // (k + 1) for l in lens) <= numOps) + 2

    def test(self):
        for s, numOps, expected in [
            ("000001", 1, 2),
            ("0000", 2, 1),
            ("0101", 0, 1),
            ("000000", 2, 2),
            ("111111", 1, 3),
            ("00", 0, 2),
            ("0", 0, 1),
            ("1111100000", 2, 2),
            ("0110", 1, 2),
        ]:
            output = self.minLength(s, numOps)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
