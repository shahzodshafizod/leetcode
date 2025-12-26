from typing import List
import unittest

# https://leetcode.com/problems/base-7/

# python3 -m unittest maths/0504-base-7.py


class Solution(unittest.TestCase):
    # Approach 1: Using Python built-in (not recommended for interview)
    # Python has numpy.base_repr but we shouldn't use it
    # Time Complexity: O(log7(n))
    # Space Complexity: O(log7(n))

    # # Approach 2: Mathematical Conversion (Optimal)
    # # Repeatedly divide by 7 and collect remainders
    # # Build result from remainders in reverse order
    # # Time Complexity: O(log7(n)) - number of digits in base 7
    # # Space Complexity: O(log7(n)) - for result string
    # def convertToBase7(self, num: int) -> str:
    #     # Handle zero case
    #     if num == 0:
    #         return "0"

    #     # Handle negative numbers
    #     is_negative = num < 0
    #     if is_negative:
    #         num = -num

    #     result = ""

    #     # Convert to base 7
    #     while num > 0:
    #         remainder = num % 7
    #         result = str(remainder) + result
    #         num //= 7

    #     if is_negative:
    #         result = "-" + result

    #     return result

    # Alternative using list for efficiency
    def convertToBase7(self, num: int) -> str:
        if num == 0:
            return "0"

        is_negative = num < 0
        if is_negative:
            num = -num

        digits: List[str] = []

        while num > 0:
            digits.append(str(num % 7))
            num //= 7

        # Reverse to get correct order
        digits.reverse()

        result = "".join(digits)

        return "-" + result if is_negative else result

    def test(self):
        for num, expected in [
            (100, "202"),
            (-7, "-10"),
            (0, "0"),
            (7, "10"),
            (-8, "-11"),
            (49, "100"),
        ]:
            output = self.convertToBase7(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
