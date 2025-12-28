import unittest

# https://leetcode.com/problems/binary-number-with-alternating-bits/

# python3 -m unittest bits/0693-binary-number-with-alternating-bits.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check each adjacent pair
    # # Convert to binary string and check consecutive characters
    # # Time: O(log N) - binary representation has log(N) bits
    # # Space: O(log N) - string storage for binary representation
    # def hasAlternatingBits(self, n: int) -> bool:
    #     binary = bin(n)[2:]  # Remove '0b' prefix
    
    #     for i in range(1, len(binary)):
    #         if binary[i] == binary[i - 1]:
    #             return False
    
    #     return True

    # # Approach #2: Bit Manipulation - Check each bit iteratively
    # # Extract bits one by one and compare with previous
    # # Time: O(log N) - iterate through all bits
    # # Space: O(1) - only store previous bit
    # def hasAlternatingBits(self, n: int) -> bool:
    #     prev_bit = n & 1
    #     n >>= 1
    
    #     while n > 0:
    #         curr_bit = n & 1
    #         if curr_bit == prev_bit:
    #             return False
    #         prev_bit = curr_bit
    #         n >>= 1
    
    #     return True

    # Approach #3: Optimized - Bit Manipulation with XOR trick
    # If bits alternate: XOR with (n >> 1) gives all 1s
    # Then (all_ones + 1) should be power of 2
    # Example: n=5 (101) -> 101 ^ 010 = 111 -> 111 + 1 = 1000 (power of 2)
    # Time: O(1) - constant time bit operations
    # Space: O(1) - only variables
    def hasAlternatingBits(self, n: int) -> bool:
        # XOR with right-shifted version
        xor = n ^ (n >> 1)
        # If alternating, xor should be all 1s (like 111, 11111, etc.)
        # all_ones + 1 should be power of 2 (1000, 100000, etc.)
        # Power of 2 check: (x & (x+1)) == 0
        return (xor & (xor + 1)) == 0

    def test(self):
        for n, expected in [
            (5, True),    # 101
            (7, False),   # 111
            (11, False),  # 1011
            (10, True),   # 1010
            (1, True),    # 1
            (2, True),    # 10
            (3, False),   # 11
            (4, False),   # 100
            (21, True),   # 10101
            (42, True),   # 101010
            (85, True),   # 1010101
        ]:
            output = self.hasAlternatingBits(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
