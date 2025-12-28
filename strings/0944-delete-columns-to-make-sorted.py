from typing import List
import unittest

# https://leetcode.com/problems/delete-columns-to-make-sorted/

# python3 -m unittest strings/0944-delete-columns-to-make-sorted.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check each column individually
    # # For each column, iterate through all rows and check if sorted
    # # Count columns that are not sorted
    # # N = number of strings, M = length of each string
    # # Time: O(N * M) - check each character in the grid
    # # Space: O(1) - only counter variable
    # def minDeletionSize(self, strs: List[str]) -> int:
    #     if not strs:
    #         return 0
    
    #     rows = len(strs)
    #     cols = len(strs[0])
    #     delete_count = 0
    
    #     for col in range(cols):
    #         # Check if this column is sorted
    #         for row in range(1, rows):
    #             if strs[row][col] < strs[row - 1][col]:
    #                 # Found unsorted pair, delete this column
    #                 delete_count += 1
    #                 break
    
    #     return delete_count

    # Approach #2: Optimized - Single pass column check
    # Same logic but cleaner - iterate columns, check if sorted
    # N = number of strings, M = length of each string
    # Time: O(N * M) - check each character once
    # Space: O(1) - only counter variable
    def minDeletionSize(self, strs: List[str]) -> int:
        rows = len(strs)
        cols = len(strs[0]) if strs else 0
        delete_count = 0

        for col in range(cols):
            # Check if column is sorted
            for row in range(1, rows):
                if strs[row][col] < strs[row - 1][col]:
                    delete_count += 1
                    break

        return delete_count

    def test(self):
        for strs, expected in [
            (["cba", "daf", "ghi"], 1),
            (["a", "b"], 0),
            (["zyx", "wvu", "tsr"], 3),
            (["abc", "bce", "cae"], 1),
            (["a"], 0),
            (["abc"], 0),
            (["rrjk", "furt", "guzm"], 2),
        ]:
            output = self.minDeletionSize(strs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
