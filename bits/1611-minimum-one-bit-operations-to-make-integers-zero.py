import unittest

# https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/

"""
n = 6
Output: 4
bin(6) = 110

[1]10 -> 010
01[0] -> 011
0[1]1 -> 001
00[1] -> 000
"""

# python3 -m unittest bits/1611-minimum-one-bit-operations-to-make-integers-zero.py

class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # Time: O(O(log^2(n)))
    # # Space: O(logn)
    # def minimumOneBitOperations(self, n: int) -> int:
    #     if n == 0:
    #         return 0
    #     power2 = 1
    #     while power2 <= n:
    #         power2 *= 2
    #     return power2 - 1 - self.minimumOneBitOperations(power2//2 ^ n)

    # # Approach: Iterative
    # # Time: O(O(log^2(n)))
    # # Space: O(1)
    # def minimumOneBitOperations(self, n: int) -> int:
    #     mask, power2 = 1, 1
    #     operations = 0
    #     while mask <= n:
    #         power2 *= 2
    #         if n & mask:
    #             operations = power2 - 1 - operations
    #         mask <<= 1
    #     return operations

    # Approach: Gray Code
    # Time: O(logn)
    # Space: O(1)
    def minimumOneBitOperations(self, n: int) -> int:
        operations = 0
        while n > 0:
            operations ^= n
            n //= 2
        return operations

    def test(self) -> None:
        for n, expected in [
            (3, 2),
		    (6, 4),
            (333, 393),
        ]:
            output = self.minimumOneBitOperations(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
