import unittest
from typing import List

# https://leetcode.com/problems/to-lower-case/

# python3 -m unittest strings/0709-to-lower-case.py


class Solution(unittest.TestCase):
    # # Approach #1: Built-in function
    # # Use Python's built-in str.lower() method
    # # N = length of string
    # # Time: O(N) - iterate through all characters
    # # Space: O(N) - create new string
    # def toLowerCase(self, s: str) -> str:
    #     return s.lower()

    # Approach #2: ASCII manipulation
    # Manually convert each uppercase letter using ASCII math
    # Uppercase letters: A-Z have ASCII codes 65-90
    # Lowercase letters: a-z have ASCII codes 97-122
    # Difference: ord('a') - ord('A') = 32
    # N = length of string
    # Time: O(N) - iterate through all characters
    # Space: O(N) - create new string/list
    def toLowerCase(self, s: str) -> str:
        result: List[str] = []
        for char in s:
            # Check if character is uppercase (A-Z)
            if 'A' <= char <= 'Z':
                # Convert to lowercase by adding 32 to ASCII value
                result.append(chr(ord(char) + 32))
            else:
                # Keep character as is
                result.append(char)
        return ''.join(result)

    def test(self):
        for s, expected in [
            ("Hello", "hello"),
            ("here", "here"),
            ("LOVELY", "lovely"),
            ("", ""),
            ("a", "a"),
            ("A", "a"),
            ("123!@# ABCxyz", "123!@# abcxyz"),
        ]:
            output = self.toLowerCase(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
