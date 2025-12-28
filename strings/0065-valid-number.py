import unittest

# https://leetcode.com/problems/valid-number/

# python3 -m unittest strings/0065-valid-number.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Try/Except parsing
    # # Try to convert string to float and catch exceptions
    # # Time: O(N) - parsing the string
    # # Space: O(1) - no extra space
    # def isNumber(self, s: str) -> bool:
    #     try:
    #         float(s)
    #         return True
    #     except ValueError:
    #         return False
    # # Note: This doesn't work because Python's float() is more permissive
    # # than the problem requirements (e.g., accepts "inf", "nan")

    # Approach #2: Optimized - State Machine / Character-by-character parsing
    # Parse string while tracking: signs, digits, decimal point, exponent
    # Valid number = (integer OR decimal) + optional exponent
    # Time: O(N) - single pass through string
    # Space: O(1) - only boolean flags
    def isNumber(self, s: str) -> bool:
        seen_digit = False
        seen_exponent = False
        seen_dot = False

        for i, char in enumerate(s):
            if char.isdigit():
                seen_digit = True
            elif char in ['+', '-']:
                # Sign is valid at start OR right after 'e'/'E'
                if i > 0 and s[i - 1] not in ['e', 'E']:
                    return False
            elif char == '.':
                # Dot is invalid after exponent or another dot
                if seen_exponent or seen_dot:
                    return False
                seen_dot = True
            elif char in ['e', 'E']:
                # Exponent requires at least one digit before it
                # and cannot appear twice
                if seen_exponent or not seen_digit:
                    return False
                seen_exponent = True
                seen_digit = False  # Must have digit after exponent
            else:
                # Invalid character
                return False

        return seen_digit

    def test(self):
        for s, expected in [
            # Valid numbers
            ("0", True),
            ("2", True),
            ("0089", True),
            ("-0.1", True),
            ("+3.14", True),
            ("4.", True),
            ("-.9", True),
            ("2e10", True),
            ("-90E3", True),
            ("3e+7", True),
            ("+6e-1", True),
            ("53.5e93", True),
            ("-123.456e789", True),
            # Invalid numbers
            ("abc", False),
            ("1a", False),
            ("1e", False),
            ("e3", False),
            ("99e2.5", False),
            ("--6", False),
            ("-+3", False),
            ("95a54e53", False),
            ("e", False),
            (".", False),
            ("+", False),
            ("6e6.5", False),
            (".e1", False),
            ("3.", True),
            (".1", True),
        ]:
            output = self.isNumber(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
