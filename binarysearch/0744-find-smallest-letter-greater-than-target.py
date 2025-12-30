from typing import List
import unittest

# https://leetcode.com/problems/find-smallest-letter-greater-than-target/

# python3 -m unittest binarysearch/0744-find-smallest-letter-greater-than-target.py


class Solution(unittest.TestCase):
    # # Approach 1: Linear Search
    # # Simply iterate through the array and return the first character greater than target
    # # If no such character exists, return the first character (wrap around)
    # # Time: O(n) where n is the length of letters
    # # Space: O(1)
    # def nextGreatestLetter(self, letters: List[str], target: str) -> str:
    #     for letter in letters:
    #         if letter > target:
    #             return letter
    #     # If no character is greater than target, wrap around to first character
    #     return letters[0]

    # Approach 2: Binary Search (Optimized)
    # Since the array is sorted, we can use binary search to find the smallest character
    # greater than target in O(log n) time
    # The key insight is that we're looking for the leftmost position where letter > target
    # Time: O(log n) where n is the length of letters
    # Space: O(1)
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters) - 1

        # Binary search to find the smallest character > target
        while left <= right:
            mid = (left + right) // 2

            if letters[mid] > target:
                # This could be our answer, but there might be a smaller one to the left
                right = mid - 1
            else:
                # letters[mid] <= target, so we need to search right
                left = mid + 1

        # If left goes beyond the array, wrap around to the first character
        # Otherwise, left points to the smallest character > target
        return letters[left % len(letters)]

    def test(self):
        for letters, target, expected in [
            (["c", "f", "j"], "a", "c"),
            (["c", "f", "j"], "c", "f"),
            (["x", "x", "y", "y"], "z", "x"),
            (["e", "e", "e", "e", "e", "e", "n", "n", "n", "n"], "e", "n"),
            (["e", "e", "e", "e", "e", "e", "n", "n", "n", "n"], "n", "e"),
            (["c", "f", "j"], "d", "f"),
            (["c", "f", "j"], "g", "j"),
            (["c", "f", "j"], "j", "c"),
        ]:
            output = self.nextGreatestLetter(letters, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
