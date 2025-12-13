from typing import List
import unittest

# https://leetcode.com/problems/1-bit-and-2-bit-characters/

# python3 -m unittest arrays/0717-1-bit-and-2-bit-characters.py


class Solution(unittest.TestCase):
    # # Approach #1: Single Pass Greedy
    # # Time: O(n)
    # # Space: O(1)
    # def isOneBitCharacter(self, bits: List[int]) -> bool:
    #     i, n = 0, len(bits)
    #     while i < n - 1:  # Stop before last bit
    #         if bits[i] == 1:  # skip two bits: Two-bit character (10 or 11)
    #             i += 2
    #         else:  # skip one bit: One-bit character (0)
    #             i += 1
    #     return i == n - 1

    # Approach #2: Greedy
    # Time: O(n)
    # Space: O(1)
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        i, n = 0, len(bits)
        while i < n - 1:  # Stop before last bit
            i += bits[i] + 1
        return i == n - 1

    def test(self):
        for bits, expected in [
            ([1, 0, 0], True),
            ([1, 1, 1, 0], False),
            ([0], True),
            ([1, 0], False),
            ([0, 0], True),
            ([1, 1, 0], True),
            ([0, 1, 0], False),
        ]:
            output = self.isOneBitCharacter(bits)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
