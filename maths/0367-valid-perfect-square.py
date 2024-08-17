import unittest

# https://leetcode.com/problems/valid-perfect-square/

# python3 -m unittest maths/0367-valid-perfect-square.py

class Solution(unittest.TestCase):
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True
        left, right = 2, num // 2
        while left <= right:
            mid = left + (right - left) // 2
            square = mid * mid
            if square > num:
                right = mid - 1
            elif square < num:
                left = mid + 1
            else:
                return True
        return False

    def testIsPerfectSquare(self) -> None:
        for num, expected in [
            (16, True),
            (14, False),
            (2147483647, False),
            (1, True),
            (2, False),
            (3, False),
            (4, True),
        ]:
            output = self.isPerfectSquare(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
