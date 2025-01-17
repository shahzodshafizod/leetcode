from typing import List
import unittest

# https://leetcode.com/problems/neighboring-bitwise-xor/

# python3 -m unittest bits/2683-neighboring-bitwise-xor.py

class Solution(unittest.TestCase):
    # Approach: Bit Manipulation
    # Time: O(n)
    # Space: O(1)
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        xorsum = 0
        for item in derived:
            xorsum ^= item
        return xorsum == 0

    def test(self):
        for derived, expected in [
            ([1,1,0], True),
            ([1,1,0], True),
            ([1,0], False),
        ]:
            output = self.doesValidArrayExist(derived)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
