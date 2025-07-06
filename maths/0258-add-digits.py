import unittest

# https://leetcode.com/problems/add-digits/

# python3 -m unittest maths/0258-add-digits.py


class Solution(unittest.TestCase):
    # # Approach #1: Recursive
    # # Time: O(D), D=# of digits
    # # Space: O(D), for recursion stack
    # def addDigits(self, num: int) -> int:
    #     if num < 10: return num
    #     sum = 0
    #     while num > 0:
    #         num, digit = divmod(num, 10)
    #         sum += digit
    #     return self.addDigits(sum)

    # # Approach #2: Iterative
    # # Time: O(D), D=# of digits
    # # Space: O(1)
    # def addDigits(self, num: int) -> int:
    #     while num > 9:
    #         sum = 0
    #         while num > 0:
    #             num, digit = divmod(num, 10)
    #             sum += digit
    #         num = sum
    #     return num

    # # Approach #3: Digital Root
    # # Time: O(D), D=# of digits
    # # Space: O(1)
    # def addDigits(self, num: int) -> int:
    #     while num > 9:
    #         num = num//10 + num%10
    #     return num

    # Follow up: Could you do it without any loop/recursion in O(1) runtime?
    # Approach #4: Digital Root (Time-Optimized)
    # Time: O(1)
    # Space: O(1)
    def addDigits(self, num: int) -> int:
        # A simple idea why digital root equals to mod 9 if we've got an ABCD number
        # ABCD = 1000A+100B+10*C+D = (A + B + C + D) + 9 * (111 * A + 11 * B + C)
        # this equals (mod 9) to A + B + C + D.
        return 1 + (num - 1) % 9 if num else 0

    def testAddDigits(self) -> None:
        for num, expected in [
            (38, 2),
            (0, 0),
            (121, 4),
            (10, 1),
        ]:
            output = self.addDigits(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
