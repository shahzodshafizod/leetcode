from typing import List
import unittest

# https://leetcode.com/problems/delete-columns-to-make-sorted-iii/

# python3 -m unittest dp/0960-delete-columns-to-make-sorted-iii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Try all column subsets
    # # Generate all 2^m subsets of columns and check if each subset is valid
    # # A subset is valid if keeping only those columns makes each row sorted
    # # N = number of strings, M = length of each string
    # # Time: O(2^M * N * M) - 2^M subsets, O(N*M) to validate each
    # # Space: O(M) - recursion stack and subset tracking
    # def minDeletionSize(self, strs: List[str]) -> int:
    #     m = len(strs[0])
    #     min_deletions = m  # Worst case: delete all columns
    
    #     def is_valid(columns: List[int]) -> bool:
    #         """Check if keeping these columns makes all rows sorted"""
    #         if not columns:
    #             return True
    #         # Check each row is sorted with these columns
    #         for row in strs:
    #             for i in range(1, len(columns)):
    #                 if row[columns[i - 1]] > row[columns[i]]:
    #                     return False
    #         return True
    
    #     def backtrack(idx: int, kept: List[int]):
    #         nonlocal min_deletions
    #         # Base case: checked all columns
    #         if idx == m:
    #             if is_valid(kept):
    #                 deletions = m - len(kept)
    #                 min_deletions = min(min_deletions, deletions)
    #             return
    
    #         # Try keeping column idx
    #         kept.append(idx)
    #         backtrack(idx + 1, kept)
    #         kept.pop()
    
    #         # Try deleting column idx
    #         backtrack(idx + 1, kept)
    
    #     backtrack(0, [])
    #     return min_deletions

    # Approach #2: Optimized - Dynamic Programming (LIS variant)
    # Key insight: Find maximum columns we can KEEP (like LIS on columns)
    # dp[i] = max columns we can keep ending at column i
    # For each column i, check all previous columns j < i
    # If column j can come before i (all chars at j <= chars at i), update dp[i]
    # N = number of strings, M = length of each string
    # Time: O(M^2 * N) - M^2 column pairs, N to check compatibility
    # Space: O(M) - DP array
    def minDeletionSize(self, strs: List[str]) -> int:
        n, m = len(strs), len(strs[0])

        # dp[i] = maximum number of columns we can keep ending at column i
        dp = [1] * m  # Each column can at least keep itself

        # For each column i, check which previous columns j can come before it
        for i in range(1, m):
            for j in range(i):
                # Check if column j can come before column i
                # This means: for all rows, strs[row][j] <= strs[row][i]
                can_precede = True
                for row in range(n):
                    if strs[row][j] > strs[row][i]:
                        can_precede = False
                        break

                if can_precede:
                    dp[i] = max(dp[i], dp[j] + 1)

        # Maximum columns we can keep
        max_kept = max(dp)

        # Minimum columns to delete
        return m - max_kept

    def test(self):
        for strs, expected in [
            (["babca", "bbazb"], 3),
            (["edcba"], 4),
            (["ghi", "def", "abc"], 0),
            (["a", "b"], 0),
            (["zyx", "wvu", "tsr"], 2),
            (["abc", "bcd", "cde"], 0),
            (["ca", "bb", "ac"], 1),
        ]:
            output = self.minDeletionSize(strs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
