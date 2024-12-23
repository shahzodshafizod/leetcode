import unittest

# https://leetcode.com/problems/hamming-distance/

# python3 -m unittest bits/0461-hamming-distance.py

# Note: This question is the same as 2220: Minimum Bit Flips to Convert Number.

class Solution(unittest.TestCase):
    # # Approach #1: Recursive
    # # Time: O(b), b=# of bits of max(start, goal)
    # # Space: O(b)
    # def hammingDistance(self, x: int, y: int) -> int:
    #     if x == 0 and y == 0: return 0
    #     return (x&1 ^ y&1) + self.hammingDistance(x>>1, y>>1)

    # # Approach #2: Iterative
    # # Time: O(b), b=# of bits of max(start, goal)
    # # Space: O(1)
    # def hammingDistance(self, x: int, y: int) -> int:
    #     distance = 0
    #     while x > 0 or y > 0:
    #         distance += x&1 ^ y&1
    #         x >>= 1
    #         y >>= 1
    #     return distance

    # Approach #3: Bit Manipulation
    # Time: O(b), b=# of bits of max(start, goal)
    # Space: O(1)
    def hammingDistance(self, x: int, y: int) -> int:
        return str(bin(x^y)).count("1")

    def testHammingDistance(self) -> None:
        for x, y, expected in [
            (1, 4, 2),
            (3, 1, 1),
        ]:
            output = self.hammingDistance(x, y)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
