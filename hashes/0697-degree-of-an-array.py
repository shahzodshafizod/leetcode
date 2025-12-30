from typing import List, Dict
import unittest

# https://leetcode.com/problems/degree-of-an-array/

# python3 -m unittest hashes/0697-degree-of-an-array.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Try all subarrays
    # # For each subarray, calculate its degree and check if it matches array degree
    # # N = length of array
    # # Time: O(N^3) - N^2 subarrays, O(N) to calculate degree for each
    # # Space: O(N) - frequency map for each subarray
    # def findShortestSubArray(self, nums: List[int]) -> int:
    #     # First, find the degree of the full array
    #     from collections import Counter
    #     full_degree = max(Counter(nums).values())
    
    #     n = len(nums)
    #     min_length = n  # Worst case: entire array
    
    #     # Try all subarrays
    #     for start in range(n):
    #         freq: Dict[int, int] = {}
    #         for end in range(start, n):
    #             # Add nums[end] to current subarray
    #             freq[nums[end]] = freq.get(nums[end], 0) + 1
    
    #             # Check if this subarray has the same degree
    #             current_degree = max(freq.values())
    #             if current_degree == full_degree:
    #                 min_length = min(min_length, end - start + 1)
    
    #     return min_length

    # Approach #2: Optimized - Hash Table with first/last occurrence tracking
    # Track frequency, first occurrence, and last occurrence for each element
    # For elements with max frequency, find minimum (last - first + 1)
    # N = length of array
    # Time: O(N) - single pass through array
    # Space: O(N) - hash tables for tracking
    def findShortestSubArray(self, nums: List[int]) -> int:
        # Track frequency, first occurrence, last occurrence
        freq: Dict[int, int] = {}
        first: Dict[int, int] = {}
        last: Dict[int, int] = {}

        for i, num in enumerate(nums):
            freq[num] = freq.get(num, 0) + 1
            if num not in first:
                first[num] = i
            last[num] = i

        # Find the degree (max frequency)
        degree = max(freq.values())

        # Find minimum length subarray for elements with max frequency
        min_length = len(nums)
        for num, count in freq.items():
            if count == degree:
                length = last[num] - first[num] + 1
                min_length = min(min_length, length)

        return min_length

    def test(self):
        for nums, expected in [
            ([1, 2, 2, 3, 1], 2),
            ([1, 2, 2, 3, 1, 4, 2], 6),
            ([1], 1),
            ([1, 2, 3, 4, 5], 1),
            ([1, 1, 1, 1], 4),
            ([1, 2, 2, 3, 1], 2),
            ([1, 3, 2, 2, 3, 1], 2),
        ]:
            output = self.findShortestSubArray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
