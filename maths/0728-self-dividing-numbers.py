import unittest
from typing import List

# https://leetcode.com/problems/self-dividing-numbers/

# python3 -m unittest maths/0728-self-dividing-numbers.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force with String Conversion
    # # Idea: Iterate through range, convert number to string to check each digit
    # # Time Complexity: O((right - left) * log(max_num))
    # #   - Iterate through (right - left + 1) numbers
    # #   - For each number, convert to string and check digits (log(max_num) operations)
    # # Space Complexity: O(k), where k is the count of self-dividing numbers in result
    # def selfDividingNumbers(self, left: int, right: int) -> List[int]:
    #     result: List[int] = []
    #     for num in range(left, right + 1):
    #         # Convert to string to easily access each digit
    #         num_str = str(num)
    #         is_self_dividing = True

    #         for digit_char in num_str:
    #             digit = int(digit_char)
    #             # Check if digit is 0 or doesn't divide the number
    #             if digit == 0 or num % digit != 0:
    #                 is_self_dividing = False
    #                 break

    #         if is_self_dividing:
    #             result.append(num)

    #     return result

    # Approach 2: Optimized (Mathematical Digit Extraction)
    # Idea: Extract digits mathematically without string conversion
    # Time Complexity: O((right - left) * log(max_num))
    #   - Same as brute force but more efficient constant factor
    # Space Complexity: O(k), where k is the count of self-dividing numbers
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        def is_self_dividing(num: int) -> bool:
            original = num
            while num > 0:
                digit = num % 10
                # If digit is 0 or doesn't divide original number
                if digit == 0 or original % digit != 0:
                    return False
                num //= 10
            return True

        return [num for num in range(left, right + 1) if is_self_dividing(num)]

    def test(self):
        for left, right, expected in [
            (1, 22, [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]),
            (47, 85, [48, 55, 66, 77]),
            (1, 1, [1]),
            (10, 20, [11, 12, 15]),
            (21, 30, [22, 24]),
        ]:
            output = self.selfDividingNumbers(left, right)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
