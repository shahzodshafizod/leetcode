import unittest

# https://leetcode.com/problems/power-of-four/

# python3 -m unittest maths/0342-power-of-four.py


class Solution(unittest.TestCase):
    # # Approach #1: Recursive
    # def isPowerOfFour(self, n: int) -> bool:
    #     if n <= 1: return n == 1
    #     return self.isPowerOfFour(n / 4)

    # # Approach #2: Iterative
    # def isPowerOfFour(self, n: int) -> bool:
    #     while n and n%4 == 0:
    #         n //= 4
    #     return n == 1

    # Approach #3: Math
    # Time: O(1)
    # Space: O(1)
    def isPowerOfFour(self, n: int) -> bool:
        # All power of 4 numbers have 4 common features:
        # 1. greater than 0.
        one = n > 0
        # 2. only have one set bit ('1') in their binary representation.
        two = n & (n - 1) == 0
        # 3. the set bit is located in odd position from right to left.
        # The maximum number in binary representation with a length of 32 is:
        # bin(0b01010101010101010101010101010101) = hex(0x55555555) = dec(1431655765).
        three = n & 0b01010101010101010101010101010101 != 0
        # 4. subtracting 1 makes the number divisable to 3.
        four = (n - 1) % 3 == 0
        return one and two and three and four

    def testIsPowerOfFour(self) -> None:
        for n, expected in [
            (0, False),
            (1, True),
            (16, True),
            (5, False),
            (512, False),
            (1024, True),
            (2048, False),
            # (1073741824, True),
        ]:
            output = self.isPowerOfFour(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
