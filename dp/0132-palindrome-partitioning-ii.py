from typing import List
import unittest

# https://leetcode.com/problems/palindrome-partitioning-ii/

# python3 -m unittest dp/0132-palindrome-partitioning-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Recursion with backtracking
    # # Try all possible partitions, check if each is palindrome
    # # For each position, try all possible cuts and recurse
    # # N = length of string
    # # Time: O(2^N) - exponential, try all possible partitions
    # # Space: O(N) - recursion stack depth
    # def minCut(self, s: str) -> int:
    #     n = len(s)
    
    #     def is_palindrome(start: int, end: int) -> bool:
    #         while start < end:
    #             if s[start] != s[end]:
    #                 return False
    #             start += 1
    #             end -= 1
    #         return True
    
    #     def backtrack(start: int) -> int:
    #         if start >= n:
    #             return 0
    
    #         # Try all possible end positions
    #         min_cuts = n  # worst case: n-1 cuts
    #         for end in range(start, n):
    #             if is_palindrome(start, end):
    #                 # If substring is palindrome, recurse on remaining
    #                 cuts = 1 + backtrack(end + 1)
    #                 min_cuts = min(min_cuts, cuts)
    
    #         return min_cuts
    
    #     return backtrack(0) - 1  # subtract 1 because we count partitions, not cuts

    # Approach #2: Optimized - DP with palindrome pre-computation
    # Pre-compute palindrome table, then use DP for minimum cuts
    # dp[i] = minimum cuts needed for s[0:i]
    # is_palindrome[i][j] = True if s[i:j+1] is palindrome
    # N = length of string
    # Time: O(N^2) - palindrome table construction + DP
    # Space: O(N^2) - palindrome table + O(N) for DP array
    def minCut(self, s: str) -> int:
        n = len(s)

        # Build palindrome table using expand around center
        is_palindrome: List[List[bool]] = [[False] * n for _ in range(n)]

        # Every single character is a palindrome
        for i in range(n):
            is_palindrome[i][i] = True

        # Check for length 2 and more
        for length in range(2, n + 1):
            for start in range(n - length + 1):
                end = start + length - 1
                if s[start] == s[end]:
                    if length == 2:
                        is_palindrome[start][end] = True
                    else:
                        is_palindrome[start][end] = is_palindrome[start + 1][end - 1]

        # DP array: dp[i] = minimum cuts for s[0:i]
        dp: List[int] = [0] * n

        for i in range(n):
            if is_palindrome[0][i]:
                # If entire substring from start is palindrome, no cuts needed
                dp[i] = 0
            else:
                # Try all possible cuts
                min_cuts = i  # worst case: i cuts for i+1 characters
                for j in range(i):
                    # If s[j+1:i+1] is palindrome, we can cut at position j
                    if is_palindrome[j + 1][i]:
                        min_cuts = min(min_cuts, dp[j] + 1)
                dp[i] = min_cuts

        return dp[n - 1]

    def test(self):
        for s, expected in [
            ("aab", 1),
            ("a", 0),
            ("ab", 1),
            ("ababababababababababababcbabababababababababababa", 0),
            ("abcdefg", 6),
            ("racecar", 0),
            ("aabbcc", 2),
            ("aaabaa", 1),
        ]:
            output = self.minCut(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
