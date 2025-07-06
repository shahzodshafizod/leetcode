from typing import List
import unittest

# https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/

# python3 -m unittest strings/2138-divide-a-string-into-groups-of-size-k.py


class Solution(unittest.TestCase):
    def divideString(self, s: str, k: int, fill: str) -> List[str]:
        # return [s[i : i + k].ljust(k, fill) for i in range(0, len(s), k)]
        parts = []
        for start in range(0, len(s), k):
            parts.append(s[start : start + k])
        parts[-1] += fill * (k - len(parts[-1]))
        return parts

    def test(self):
        for s, k, fill, expected in [
            ("abcdefghi", 3, "x", ["abc", "def", "ghi"]),
            ("abcdefghij", 3, "x", ["abc", "def", "ghi", "jxx"]),
        ]:
            output = self.divideString(s, k, fill)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
