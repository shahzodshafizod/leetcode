import unittest

# https://leetcode.com/problems/count-binary-substrings/

# python3 -m unittest twopointers/0696-count-binary-substrings.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check all substrings
    # # For each possible substring, check if valid (equal consecutive 0s and 1s)
    # # N = length of string
    # # Time: O(N^2) - check all substrings
    # # Space: O(1) - no extra space
    # def countBinarySubstrings(self, s: str) -> int:
    #     count = 0
    #     n = len(s)
    
    #     for i in range(n):
    #         # Count consecutive characters starting at i
    #         zeros = 0
    #         ones = 0
    #         char = s[i]
    
    #         # Count first group
    #         j = i
    #         while j < n and s[j] == char:
    #             if char == '0':
    #                 zeros += 1
    #             else:
    #                 ones += 1
    #             j += 1
    
    #         # Count second group (opposite character)
    #         if j < n:
    #             next_char = s[j]
    #             while j < n and s[j] == next_char:
    #                 if next_char == '0':
    #                     zeros += 1
    #                 else:
    #                     ones += 1
    #                 j += 1
    #                 # Check if we have equal counts
    #                 if zeros == ones:
    #                     count += 1
    
    #     return count

    # Approach #2: Optimized - Group counting with two pointers
    # Count consecutive groups, then calculate valid substrings between adjacent groups
    # For groups of size m and n, can form min(m, n) valid substrings
    # Example: "000111" -> groups [3, 3] -> min(3, 3) = 3 substrings
    # N = length of string
    # Time: O(N) - single pass to count groups
    # Space: O(1) - only track previous and current group counts
    def countBinarySubstrings(self, s: str) -> int:
        prev_count = 0
        curr_count = 1
        result = 0

        for i in range(1, len(s)):
            if s[i] == s[i - 1]:
                # Same character, extend current group
                curr_count += 1
            else:
                # Different character, start new group
                # Add min(prev, curr) valid substrings
                result += min(prev_count, curr_count)
                prev_count = curr_count
                curr_count = 1

        # Don't forget the last group
        result += min(prev_count, curr_count)

        return result

    def test(self):
        for s, expected in [
            ("00110011", 6),
            ("10101", 4),
            ("00110", 3),
            ("000111", 3),
            ("01", 1),
            ("0011", 2),
            ("00100", 2),
            ("001101", 4),
        ]:
            output = self.countBinarySubstrings(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
