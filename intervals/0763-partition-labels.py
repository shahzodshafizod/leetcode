from typing import List
import unittest

# https://leetcode.com/problems/partition-labels/

# python3 -m unittest intervals/0763-partition-labels.py


class Solution(unittest.TestCase):
    def partitionLabels(self, s: str) -> List[int]:
        last = {}
        for idx, c in enumerate(s):
            last[c] = idx
        sizes = []
        start, end = 0, -1
        for idx, c in enumerate(s):
            end = max(end, last[c])
            if idx == end:
                sizes.append(end - start + 1)
                start = idx + 1
        return sizes

    def test(self):
        for s, expected in [
            ("ababcbacadefegdehijhklij", [9, 7, 8]),
            ("eccbbbbdec", [10]),
        ]:
            output = self.partitionLabels(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
