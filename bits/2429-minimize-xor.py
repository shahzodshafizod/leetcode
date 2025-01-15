import unittest

# https://leetcode.com/problems/minimize-xor/

# python3 -m unittest bits/2429-minimize-xor.py

class Solution(unittest.TestCase):
    # Approach: Bit Manipulation
    # Time: O(log n)
    # Space: O(1)
    def minimizeXor(self, num1: int, num2: int) -> int:
        x = num1
        diff = num2.bit_count() - num1.bit_count()
        bit = 0
        while diff > 0 and bit < 32:
            if (1 << bit) & x == 0:
                x ^= 1 << bit
                diff -= 1
            bit += 1
        while diff < 0:
            x &= x - 1
            diff += 1
        return x

    def test(self):
        for num1, num2, expected in [
            (3, 5, 3),
            (1, 12, 3),
        ]:
            output = self.minimizeXor(num1, num2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
