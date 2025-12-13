from typing import List
import unittest

# https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/

# python3 -m unittest arrays/3036-number-of-subarrays-that-match-a-pattern-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # For each position, check if the pattern matches
    # # Pattern: 1 = increase, 0 = equal, -1 = decrease
    # # Time: O(n * m) - too slow for large inputs due to checking pattern at each position
    # # This approach checks every possible starting position and validates the entire pattern,
    # # which becomes inefficient for large arrays
    # # Space: O(1)
    # def countMatchingSubarrays(self, nums: List[int], pattern: List[int]) -> int:
    #     n = len(nums)
    #     m = len(pattern)
    #     count = 0

    #     for i in range(n - m):
    #         match = True
    #         for j in range(m):
    #             if pattern[j] == 1:
    #                 if nums[i + j + 1] <= nums[i + j]:
    #                     match = False
    #                     break
    #             elif pattern[j] == 0:
    #                 if nums[i + j + 1] != nums[i + j]:
    #                     match = False
    #                     break
    #             else:  # pattern[j] == -1
    #                 if nums[i + j + 1] >= nums[i + j]:
    #                     match = False
    #                     break
    #         if match:
    #             count += 1

    #     return count

    # Approach #2: Optimized - Convert to pattern array and use KMP or Rolling Hash
    # Convert nums to pattern representation, then find pattern matches
    # Time: O(n + m) using KMP algorithm
    # Space: O(n + m) for converted array and KMP table
    def countMatchingSubarrays(self, nums: List[int], pattern: List[int]) -> int:
        # Convert nums to pattern representation
        n = len(nums)
        nums_pattern: List[int] = []

        for i in range(n - 1):
            if nums[i + 1] > nums[i]:
                nums_pattern.append(1)
            elif nums[i + 1] == nums[i]:
                nums_pattern.append(0)
            else:
                nums_pattern.append(-1)

        # Use KMP to find pattern matches
        def kmp_search(text: List[int], pattern: List[int]) -> int:
            if not pattern:
                return 0

            # Build LPS (Longest Prefix Suffix) array
            lps = [0] * len(pattern)
            length = 0
            i = 1

            while i < len(pattern):
                if pattern[i] == pattern[length]:
                    length += 1
                    lps[i] = length
                    i += 1
                else:
                    if length != 0:
                        length = lps[length - 1]
                    else:
                        lps[i] = 0
                        i += 1

            # Search for pattern in text
            count = 0
            i = 0  # text index
            j = 0  # pattern index

            while i < len(text):
                if text[i] == pattern[j]:
                    i += 1
                    j += 1

                if j == len(pattern):
                    count += 1
                    j = lps[j - 1]
                elif i < len(text) and text[i] != pattern[j]:
                    if j != 0:
                        j = lps[j - 1]
                    else:
                        i += 1

            return count

        return kmp_search(nums_pattern, pattern)

    def test(self):
        for nums, pattern, expected in [
            ([1, 2, 3, 4, 5, 6], [1, 1], 4),
            ([1, 4, 4, 1, 3, 5, 5, 3], [1, 0, -1], 2),
            ([1, 2, 3], [1], 2),
            ([5, 5, 5, 5], [0, 0], 2),
        ]:
            output = self.countMatchingSubarrays(nums, pattern)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
