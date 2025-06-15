import unittest

# https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/

# python3 -m unittest greedy/2566-maximum-difference-by-remapping-a-digit.py

class Solution(unittest.TestCase):
    # Time: O(log10(num))
    # Space: O(1)
    def minMaxDifference(self, num: int) -> int:
        num, nine = str(num), '9'
        for d in num:
            if d != '9':
                nine = d
                break
        return int(num.replace(nine, '9')) - int(num.replace(num[0], '0'))

    def test(self):
        for num, expected in [
            (11891, 99009),
            (90, 99),
        ]:
            output = self.minMaxDifference(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
