import unittest

# https://leetcode.com/problems/sqrtx/

# python3 -m unittest binarysearch/0069-sqrtx.py


class Solution(unittest.TestCase):
    def mySqrt(self, x: int) -> int:
        left, right = 1, x
        while left <= right:
            mid = left + (right - left) // 2
            if mid * mid > x:
                right = mid - 1
            else:
                left = mid + 1
        return right

    def testMySqrt(self) -> None:
        for x, expected in [
            (0, 0),
            (1, 1),
            (4, 2),
            (8, 2),
            (2147395599, 46339),
            (2147395600, 46340),
        ]:
            output = self.mySqrt(x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
