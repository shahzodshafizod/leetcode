import unittest

# https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/

# python3 -m unittest strings/1790-check-if-one-string-swap-can-make-strings-equal.py


class Solution(unittest.TestCase):
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        diff_count = 0
        index1, index2 = 0, 0
        for idx, (c1, c2) in enumerate(zip(s1, s2)):
            if c1 == c2:
                continue
            diff_count += 1
            match diff_count:
                case 1:
                    index1 = idx
                case 2:
                    index2 = idx
                case 3:
                    return False
        return s1[index1] == s2[index2] and s1[index2] == s2[index1]

    def test(self):
        for s1, s2, expected in [
            ("bank", "kanb", True),
            ("attack", "defend", False),
            ("kelb", "kelb", True),
            ("xkkkkkkkx", "fkkkkkkkf", False),
        ]:
            output = self.areAlmostEqual(s1, s2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
