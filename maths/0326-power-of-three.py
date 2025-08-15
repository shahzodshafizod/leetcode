import unittest

# https://leetcode.com/problems/power-of-three/

# python3 -m unittest maths/0326-power-of-three.py


class Solution(unittest.TestCase):
    def isPowerOfThree(self, n: int) -> bool:
        # if n <= 0:
        #     return False
        # while n % 3 == 0:
        #     n //= 3
        # return n == 1

        # every power of 3 divides to the earliest powers of 3.
        # in 32 bit numbers, 3^20 is the largest power of three
        # 3^20 = 3486784401
        return n > 0 and 3486784401 % n == 0

    def test(self):
        for n, expected in [
            (27, True),
            (0, False),
            (-1, False),
        ]:
            output = self.isPowerOfThree(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
