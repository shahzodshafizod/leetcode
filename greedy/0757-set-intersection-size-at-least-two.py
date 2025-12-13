from typing import List
import unittest

# from itertools import combinations

# https://leetcode.com/problems/set-intersection-size-at-least-two/

# python3 -m unittest greedy/0757-set-intersection-size-at-least-two.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # Time: O(C(max_val, k) * n) where k grows - exponential with large ranges
    # # Space: O(max_val) for combinations
    # def intersectionSizeTwo(self, intervals: List[List[int]]) -> int:
    #     # Find the range of all possible values
    #     max_val = max(end for _, end in intervals)
    #     min_val = min(start for start, _ in intervals)

    #     # Try all possible subsets of size k
    #     for size in range(2, max_val - min_val + 2):
    #         for combo in combinations(range(min_val, max_val + 1), size):
    #             combo_set = set(combo)
    #             valid = True

    #             for start, end in intervals:
    #                 count = sum(1 for x in combo_set if start <= x <= end)
    #                 if count < 2:
    #                     valid = False
    #                     break

    #             if valid:
    #                 return size

    #     return 0

    # Approach #2: Greedy with Sorting
    # Time: O(n log n) - dominated by sorting
    # Space: O(1) - constant extra space (excluding sort)
    def intersectionSizeTwo(self, intervals: List[List[int]]) -> int:
        # Sort by end point, then by start point descending
        intervals.sort(key=lambda x: (x[1], -x[0]))

        result = 0
        first, second = -1, -1

        for start, end in intervals:
            # Case 1: Both points already in interval
            if start <= first:
                continue

            # Case 2: Only second point in interval
            if start <= second:
                first = second
                second = end
                result += 1
            else:
                # Case 3: Neither point in interval
                first = end - 1
                second = end
                result += 2

        return result

    def test(self):
        for intervals, expected in [
            ([[1, 3], [1, 4], [2, 5], [3, 5]], 3),
            ([[1, 2], [2, 3], [2, 4], [4, 5]], 5),
            ([[1, 10], [2, 3], [4, 5], [6, 7], [8, 9]], 8),
        ]:
            output = self.intersectionSizeTwo(intervals)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
