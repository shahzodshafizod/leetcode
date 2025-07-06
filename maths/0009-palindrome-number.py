import unittest

# https://leetcode.com/problems/palindrome-number/

# python3 -m unittest maths/0009-palindrome-number.py


class Solution(unittest.TestCase):
    def isPalindrome(self, x: int) -> bool:
        if x <= 0 or x % 10 == 0:
            return False
        half = 0
        while x > half:
            x, d = divmod(x, 10)
            half = half * 10 + d
        return x in (half, half // 10)

    def testIsPalindrome(self) -> None:
        for x, expected in [
            (121, True),
            (-121, False),
            (10, False),
            (1554, False),
            (1551, True),
        ]:
            output = self.isPalindrome(x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
