from typing import List
import unittest

# https://leetcode.com/problems/add-strings/

# python3 -m unittest strings/0415-add-strings.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Convert to integers and add
    # # Time: O(n + m) - string to int conversion
    # # Space: O(1) - excluding output
    # def addStringsBruteForce(self, num1: str, num2: str) -> str:
    #     return str(int(num1) + int(num2))

    # Approach #2: Simulation with Two Pointers
    # Time: O(max(n, m)) where n, m are lengths of num1, num2
    # Space: O(max(n, m)) for result string
    def addStrings(self, num1: str, num2: str) -> str:
        i, j = len(num1) - 1, len(num2) - 1
        carry = 0
        result: List[str] = []

        while i >= 0 or j >= 0 or carry:
            # Get current digits (0 if index out of bounds)
            digit1 = int(num1[i]) if i >= 0 else 0
            digit2 = int(num2[j]) if j >= 0 else 0

            # Calculate sum with carry
            total = digit1 + digit2 + carry
            carry = total // 10
            result.append(str(total % 10))

            i -= 1
            j -= 1

        # Result is built backwards, so reverse it
        return "".join(reversed(result))

    def test(self):
        for num1, num2, expected in [
            ("11", "123", "134"),
            ("456", "77", "533"),
            ("0", "0", "0"),
            ("1", "9", "10"),
            ("999", "1", "1000"),
            ("123456789", "987654321", "1111111110"),
        ]:
            output = self.addStrings(num1, num2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
