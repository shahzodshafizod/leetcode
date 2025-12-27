import unittest

# https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/

# python3 -m unittest hashes/0828-count-unique-characters-of-all-substrings-of-a-given-string.py


class Solution(unittest.TestCase):
    # Approach: Contribution of Each Character
    # Key insight: Instead of checking all substrings, calculate how many substrings
    # contain each character occurrence as unique (appears exactly once).
    #
    # Strategy:
    # For each occurrence of character at position i:
    # 1. Find previous occurrence of same character (prev)
    # 2. Find next occurrence of same character (next)
    # 3. Character at i is unique in substrings that:
    #    - Start after prev and at or before i
    #    - End at or after i and before next
    # 4. Count = (i - prev) * (next - i)
    #
    # Example: "ABA"
    # - A at index 0: prev=-1, next=2, count=(0-(-1))*(2-0)=2 substrings: "A", "AB"
    # - B at index 1: prev=-1, next=3, count=(1-(-1))*(3-1)=4 substrings: "B", "AB", "BA", "ABA"
    # - A at index 2: prev=0, next=3, count=(2-0)*(3-2)=2 substrings: "BA", "A"
    # Total = 2 + 4 + 2 = 8
    #
    # Time Complexity: O(n) where n is length of string
    # Space Complexity: O(26) = O(1) for storing last occurrence of each character
    def uniqueLetterString(self, s: str) -> int:
        n = len(s)
        result = 0

        # Store last occurrence index of each character (A-Z)
        last = [-1] * 26

        # Store second-to-last occurrence for each character
        second_last = [-1] * 26

        for i in range(n):
            c = ord(s[i]) - ord('A')

            # Previous occurrence of character c
            prev = second_last[c]
            # Current occurrence
            curr = last[c]
            # Next occurrence (we don't know yet, so use i)
            next_pos = i

            # If there was a previous occurrence at curr,
            # calculate its contribution now that we know the next occurrence is at i
            if curr != -1:
                result += (curr - prev) * (next_pos - curr)

            # Update occurrences
            second_last[c] = last[c]
            last[c] = i

        # Process remaining characters (no next occurrence)
        for c in range(26):
            prev = second_last[c]
            curr = last[c]
            next_pos = n

            if curr != -1:
                result += (curr - prev) * (next_pos - curr)

        return result

    def test(self):
        for s, expected in [
            ("ABC", 10),
            ("ABA", 8),
            ("LEETCODE", 92),
            ("A", 1),
            ("AA", 2),
        ]:
            output = self.uniqueLetterString(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
