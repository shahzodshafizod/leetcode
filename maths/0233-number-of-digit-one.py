import unittest

# https://leetcode.com/problems/number-of-digit-one/

# python3 -m unittest maths/0233-number-of-digit-one.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # Idea: Iterate through all numbers from 0 to n and count digit '1' in each number
    # # Time Complexity: O(n * log n), where n is the input number
    # #   - We iterate through n numbers
    # #   - For each number, we count digits which takes O(log n) time
    # # Space Complexity: O(1), only using a few variables
    # def countDigitOne(self, n: int) -> int:
    #     count = 0
    #     for num in range(n + 1):
    #         # Count ones in current number
    #         temp = num
    #         while temp > 0:
    #             if temp % 10 == 1:
    #                 count += 1
    #             temp //= 10
    #     return count

    # Approach 2: Optimized (Digit DP / Mathematical Pattern)
    # Idea: Count ones at each digit position (units, tens, hundreds, etc.)
    # For each position, calculate how many times '1' appears there
    #
    # Algorithm:
    # For a number like 3141:
    # - Position 0 (units): Check digit 1
    # - Position 1 (tens): Check digit 4
    # - Position 2 (hundreds): Check digit 1
    # - Position 3 (thousands): Check digit 3
    #
    # For each position i with digit d:
    # - higher = digits before position i
    # - lower = digits after position i
    # - factor = 10^i
    #
    # Count formula:
    # - If d > 1: count = (higher + 1) * factor
    # - If d == 1: count = higher * factor + (lower + 1)
    # - If d == 0: count = higher * factor
    #
    # Time Complexity: O(log n), we process each digit position once
    # Space Complexity: O(1), only using a few variables
    def countDigitOne(self, n: int) -> int:
        if n <= 0:
            return 0

        count = 0
        factor = 1  # Current position factor (1, 10, 100, 1000, ...)

        while factor <= n:
            # Divide number into higher and lower parts
            higher = n // (factor * 10)  # Digits before current position
            current = (n // factor) % 10  # Current digit
            lower = n % factor  # Digits after current position

            # Count ones at current position
            if current > 1:
                # If current digit > 1, all combinations of higher digits
                # with 1 at current position are valid
                count += (higher + 1) * factor
            elif current == 1:
                # If current digit is 1, count depends on both higher and lower
                count += higher * factor + (lower + 1)
            else:  # current == 0
                # If current digit is 0, only higher digits contribute
                count += higher * factor

            factor *= 10

        return count

    def test(self):
        for n, expected in [
            (13, 6),  # 1,10,11,11,12,13 -> 6 ones
            (0, 0),   # No numbers with 1
            (1, 1),   # Just 1 -> 1 one
            (10, 2),  # 1,10 -> 2 ones
            (100, 21),  # Multiple positions with ones
            (824, 273),  # Larger number
        ]:
            output = self.countDigitOne(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
