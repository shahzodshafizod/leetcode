import unittest

# https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/

# python3 -m unittest bits/2749-minimum-operations-to-make-the-integer-zero.py


class Solution(unittest.TestCase):
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        # num1 = (num2 + 2^i1) + (nums2 + 2^i2) + ... + (num2 + 2^ik)
        # num1 = k*num2 + (2^i1 + 2^i2 + ... + 2^ik)
        # num1 - k*num2 = 2^i1 + 2^i2 + ... + 2^ik
        for k in range(1, 61):
            target = num1 - k * num2
            # checking if target is equal to: 2^i1 + 2^i2 + ... + 2^ik
            if target.bit_count() <= k <= target:
                return k
        return -1

    def test(self):
        for num1, num2, expected in [
            (3, -2, 3),
            (5, 7, -1),
        ]:
            output = self.makeTheIntegerZero(num1, num2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
