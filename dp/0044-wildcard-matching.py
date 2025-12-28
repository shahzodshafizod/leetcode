from typing import List
import unittest

# https://leetcode.com/problems/wildcard-matching/

# python3 -m unittest dp/0044-wildcard-matching.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Recursion
    # # Try all possible matches recursively
    # # For '*', try matching 0, 1, 2, ... characters
    # # N = length of s, M = length of p
    # # Time: O(2^(N+M)) - exponential branching at each '*'
    # # Space: O(N+M) - recursion stack depth
    # def isMatch(self, s: str, p: str) -> bool:
    #     def backtrack(i: int, j: int) -> bool:
    #         # Base cases
    #         if j == len(p):
    #             return i == len(s)
    #         if i == len(s):
    #             # Pattern can only match if remaining is all '*'
    #             return all(c == '*' for c in p[j:])
    
    #         # Match current characters
    #         if p[j] == '*':
    #             # Try matching 0 characters (skip *) or 1+ characters
    #             return backtrack(i, j + 1) or backtrack(i + 1, j)
    #         elif p[j] == '?' or p[j] == s[i]:
    #             return backtrack(i + 1, j + 1)
    #         else:
    #             return False
    
    #     return backtrack(0, 0)

    # Approach #2: Optimized - DP with memoization
    # Use 2D DP table where dp[i][j] = whether s[0:i] matches p[0:j]
    # Handle '?' and '*' with specific transitions
    # N = length of s, M = length of p
    # Time: O(N * M) - fill DP table
    # Space: O(N * M) - DP table
    def isMatch(self, s: str, p: str) -> bool:
        n, m = len(s), len(p)

        # dp[i][j] = True if s[0:i] matches p[0:j]
        dp: List[List[bool]] = [[False] * (m + 1) for _ in range(n + 1)]

        # Empty string matches empty pattern
        dp[0][0] = True

        # Empty string can match pattern with only '*'
        for j in range(1, m + 1):
            if p[j - 1] == '*':
                dp[0][j] = dp[0][j - 1]

        # Fill DP table
        for i in range(1, n + 1):
            for j in range(1, m + 1):
                if p[j - 1] == '*':
                    # '*' matches empty sequence or one+ characters
                    dp[i][j] = dp[i][j - 1] or dp[i - 1][j]
                elif p[j - 1] == '?' or p[j - 1] == s[i - 1]:
                    # '?' or exact character match
                    dp[i][j] = dp[i - 1][j - 1]
                # else: no match, remains False

        return dp[n][m]

    def test(self):
        for s, p, expected in [
            ("aa", "a", False),
            ("aa", "*", True),
            ("cb", "?a", False),
            ("adceb", "*a*b", True),
            ("acdcb", "a*c?b", False),
            ("", "*", True),
            ("", "?", False),
            ("a", "", False),
            ("", "", True),
            ("abc", "abc", True),
            ("abc", "a*c", True),
            ("abc", "a?c", True),
            ("abcdefg", "a*g", True),
            ("ho", "ho**", True),
            ("mississippi", "m*issi*ippi", True),
        ]:
            output = self.isMatch(s, p)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
