import unittest

# https://leetcode.com/problems/longest-uncommon-subsequence-i/

# python3 -m unittest strings/0521-longest-uncommon-subsequence-i.py


class Solution(unittest.TestCase):
    # Approach 1: Brute Force - Generate All Subsequences
    # Generate all subsequences of both strings
    # Find longest that appears in only one string
    # Time Complexity: O(2^n * 2^m) - exponential
    # Space Complexity: O(2^n + 2^m)
    # Not practical to implement

    # Approach 2: Brain Teaser / Logic (Optimal)
    # Key insight: If strings differ, the longer one (or either if same length)
    # is an uncommon subsequence. If identical, no uncommon subsequence exists.
    #
    # Why? If a != b, then "a" itself is subsequence of "a" but not "b".
    # If a == b, every subsequence of a is also subsequence of b.
    #
    # Time Complexity: O(min(len(a), len(b))) - string comparison
    # Space Complexity: O(1)
    def findLUSlength(self, a: str, b: str) -> int:
        # If strings are identical, no uncommon subsequence
        if a == b:
            return -1

        # If different, return the maximum length
        return max(len(a), len(b))

    def test(self):
        for a, b, expected in [
            ("aba", "cdc", 3),
            ("aaa", "bbb", 3),
            ("aaa", "aaa", -1),
            ("abc", "ab", 3),
            ("a", "b", 1),
        ]:
            output = self.findLUSlength(a, b)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
