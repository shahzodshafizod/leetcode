import unittest

# https://leetcode.com/problems/compare-version-numbers/

# python3 -m unittest strings/0165-compare-version-numbers.py


class Solution(unittest.TestCase):
    def compareVersion(self, version1: str, version2: str) -> int:
        revs1 = [int(s) for s in version1.split('.')]
        revs2 = [int(s) for s in version2.split('.')]
        n1, n2 = len(revs1), len(revs2)
        if n1 < n2:
            revs1 += [0] * (n2 - n1)
        elif n2 < n1:
            revs2 += [0] * (n1 - n2)
        return int(revs1 > revs2) - int(revs2 > revs1)

    def test(self):
        for version1, version2, expected in [
            ("1.2", "1.10", -1),
            ("1.01", "1.001", 0),
            ("1.0", "1.0.0.0", 0),
        ]:
            output = self.compareVersion(version1, version2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
