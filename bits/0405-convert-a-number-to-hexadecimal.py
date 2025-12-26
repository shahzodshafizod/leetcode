import unittest

# https://leetcode.com/problems/convert-a-number-to-hexadecimal/

# python3 -m unittest bits/0405-convert-a-number-to-hexadecimal.py


class Solution(unittest.TestCase):
    # Approach 1: Using Python built-in (not allowed by problem)
    # hex(num) would work but problem says don't use built-ins
    # Time Complexity: O(1)
    # Space Complexity: O(1)

    # # Approach 2: Bit Manipulation (Optimal)
    # # Extract 4 bits at a time from right to left
    # # Each 4 bits represents one hex digit
    # # Handle negative with two's complement (32-bit unsigned)
    # # Time Complexity: O(8) = O(1) - at most 8 hex digits
    # # Space Complexity: O(8) = O(1)
    # def toHex(self, num: int) -> str:
    #     if num == 0:
    #         return "0"

    #     # Convert negative to 32-bit unsigned representation
    #     # In Python, use & 0xFFFFFFFF to get 32-bit unsigned value
    #     if num < 0:
    #         num = num & 0xFFFFFFFF

    #     hex_chars = "0123456789abcdef"
    #     result = ""

    #     # Extract 4 bits at a time
    #     while num > 0:
    #         digit = num & 0xF  # Get last 4 bits
    #         result = hex_chars[digit] + result
    #         num >>= 4  # Shift right by 4 bits

    #     return result

    # Alternative using list for efficiency
    def toHex(self, num: int) -> str:
        if num == 0:
            return "0"

        if num < 0:
            num = num & 0xFFFFFFFF

        hex_chars = "0123456789abcdef"
        digits: list[str] = []

        while num > 0:
            digits.append(hex_chars[num & 0xF])
            num >>= 4

        # Reverse to get correct order
        digits.reverse()

        return "".join(digits)

    def test(self):
        for num, expected in [
            (26, "1a"),
            (-1, "ffffffff"),
            (0, "0"),
            (1, "1"),
            (15, "f"),
            (16, "10"),
            (255, "ff"),
        ]:
            output = self.toHex(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
