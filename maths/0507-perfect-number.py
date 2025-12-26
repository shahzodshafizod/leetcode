import unittest

# https://leetcode.com/problems/perfect-number/

# python3 -m unittest maths/0507-perfect-number.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Check All Divisors
    # # Iterate from 1 to n-1 and sum all divisors
    # # Time Complexity: O(n) - too slow for large numbers
    # # Space Complexity: O(1)
    # def checkPerfectNumber(self, num: int) -> bool:
    #     if num <= 1:
    #         return False

    #     divisor_sum = 0
    #     for i in range(1, num):
    #         if num % i == 0:
    #             divisor_sum += i

    #     return divisor_sum == num

    # # Approach 2: Optimized - Check up to sqrt(n)
    # # Key insight: divisors come in pairs (i, n/i)
    # # Only need to check up to sqrt(n)
    # # Time Complexity: O(sqrt(n))
    # # Space Complexity: O(1)
    # def checkPerfectNumber(self, num: int) -> bool:
    #     if num <= 1:
    #         return False

    #     divisor_sum = 1  # 1 is always a divisor
    #     i = 2

    #     # Check divisors up to sqrt(num)
    #     while i * i <= num:
    #         if num % i == 0:
    #             divisor_sum += i
    #             # Add paired divisor if different
    #             if i * i != num:
    #                 divisor_sum += num // i
    #         i += 1

    #     return divisor_sum == num

    # Alternative with early termination
    def checkPerfectNumber(self, num: int) -> bool:
        if num <= 1:
            return False

        divisor_sum = 1
        i = 2

        while i * i <= num:
            if num % i == 0:
                divisor_sum += i
                if i * i != num:
                    divisor_sum += num // i
                # Early termination if sum exceeds num
                if divisor_sum > num:
                    return False
            i += 1

        return divisor_sum == num

    def test(self):
        for num, expected in [
            (28, True),  # 1 + 2 + 4 + 7 + 14 = 28
            (7, False),
            (6, True),  # 1 + 2 + 3 = 6
            (496, True),  # Perfect number
            (8128, True),  # Perfect number
            (1, False),
            (2, False),
        ]:
            output = self.checkPerfectNumber(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
