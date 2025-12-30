import unittest

# https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/

# python3 -m unittest bits/0762-prime-number-of-set-bits-in-binary-representation.py


class Solution(unittest.TestCase):
    # # Approach: Bit Counting with Prime Check Optimization
    # # Key insight: Since right <= 10^6 < 2^20, the maximum number of set bits is 20
    # # Prime numbers up to 20: {2, 3, 5, 7, 11, 13, 17, 19}
    # #
    # # Algorithm:
    # # 1. For each number in [left, right], count the number of set bits
    # # 2. Check if the count is a prime number using a precomputed set
    # # 3. Count how many numbers have a prime number of set bits
    # #
    # # Optimization: Use a bitmask to represent primes up to 20
    # # primes = {2, 3, 5, 7, 11, 13, 17, 19}
    # # bitmask = 0b10100010100010101100 (bits 2,3,5,7,11,13,17,19 are set)
    # # bitmask = 665772 in decimal
    # #
    # # Time: O((right - left) * log(right)) - count bits for each number
    # # Space: O(1) - only using constant space
    # def countPrimeSetBits(self, left: int, right: int) -> int:
    #     # Primes up to 20: {2, 3, 5, 7, 11, 13, 17, 19}
    #     # Binary representation with these bits set: 0b10100010100010101100
    #     primes_mask = 0b10100010100010101100  # = 665772

    #     count = 0
    #     for num in range(left, right + 1):
    #         # Count set bits using Brian Kernighan's algorithm
    #         set_bits = 0
    #         n = num
    #         while n:
    #             n &= n - 1  # Remove the rightmost set bit
    #             set_bits += 1

    #         # Check if set_bits is prime using the bitmask
    #         if primes_mask & (1 << set_bits):
    #             count += 1

    #     return count

    # Alternative: Using built-in bit_count (Python 3.10+)
    def countPrimeSetBits(self, left: int, right: int) -> int:
        primes = {2, 3, 5, 7, 11, 13, 17, 19}
        count = 0
        for num in range(left, right + 1):
            if num.bit_count() in primes:
                count += 1
        return count

    def test(self):
        for left, right, expected in [
            (6, 10, 4),
            (10, 15, 5),
            (1, 1, 0),  # 1 has 1 set bit, 1 is not prime
            (2, 2, 0),  # 2 has 1 set bit, 1 is not prime
            (3, 3, 1),  # 3 has 2 set bits, 2 is prime
            (4, 4, 0),  # 4 has 1 set bit, 1 is not prime
            (5, 5, 1),  # 5 has 2 set bits, 2 is prime
            (1, 10, 6),  # 3,5,6,7,9,10 have prime set bits
            (244, 269, 16),
        ]:
            output = self.countPrimeSetBits(left, right)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
