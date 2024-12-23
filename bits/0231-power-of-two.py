import unittest

# https://leetcode.com/problems/power-of-two/

# python3 -m unittest bits/0231-power-of-two.py

class Solution(unittest.TestCase):
    # # Approach #1: Recursive
    # # Time: O(logn)
    # # Space: O(logn)
    # def isPowerOfTwo(self, n: int) -> bool:
    #     if n <= 1:
    #         return n == 1
    #     return self.isPowerOfTwo(n/2)

    # # Approach #2: Iterative
    # # Time: O(logn)
    # # Space: O(1)
    # def isPowerOfTwo(self, n: int) -> bool:
    #     count = 0
    #     while n > 0:
    #         count += n&1
    #         n >>= 1
    #     return count == 1

    # Approach #3: Bit Manipulation
    # Time: O(1)
    # Space: O(1)
    def isPowerOfTwo(self, n: int) -> bool:
        # return n > 0 and str(bin(n)).count('1') == 1
        return n > 0 and (n & (n-1)) == 0

    def test(self):
        for n, expected in [
            (1, True),
            (16, True),
            (3, False),
            (4, True),
            (0, False),
            (-16, False),
        ]:
            output = self.isPowerOfTwo(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
