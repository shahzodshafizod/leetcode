import unittest

# https://leetcode.com/problems/minimum-bit-flips-to-convert-number/

# python3 -m unittest bits/2220-minimum-bit-flips-to-convert-number.py

class Solution(unittest.TestCase):
    # # Approach: Iterative
    # # Time: O(max bits)
    # # Space: O(1)
    # def minBitFlips(self, start: int, goal: int) -> int:
    #     flips = 0
    #     xor_result = start ^ goal
    #     while xor_result:
    #         flips += xor_result&1
    #         xor_result >>= 1
    #     return flips

    # # Approach: Recursive
    # # Time: O(max bits)
    # # Space: O(max bits)
    # def minBitFlips(self, start: int, goal: int) -> int:
    #     if not start and not goal:
    #         return 0
    #     flips = (start&1) ^ (goal&1)
    #     return flips + self.minBitFlips(start>>1, goal>>1)

    # Approach: Brian Kernighan’s Algorithm
    # Time: O(number of set bits)
    # Space: O(1)
    def minBitFlips(self, start: int, goal: int) -> int:
        # return bin(start^goal).count('1')
        # return (start^goal).bit_count()
        flips = 0
        xor_result = start ^ goal
        while xor_result:
            xor_result &= xor_result-1
            flips += 1
        return flips

    def test(self):
        for start, goal, expected in [
            (10, 7, 3),
            (3, 4, 3),
        ]:
            output = self.minBitFlips(start, goal)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
