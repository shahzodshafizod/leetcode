from collections import Counter
from typing import List
import bisect
import unittest

# https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/

# python3 -m unittest binarysearch/3347-maximum-frequency-of-an-element-after-performing-operations-ii.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # For each possible target value, count how many elements can be transformed to it
    # # Time: O(n * range), where range is max(nums) - min(nums) + 2*k
    # # Space: O(1)
    # def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
    #     # Try all possible target values
    #     min_val = min(nums) - k
    #     max_val = max(nums) + k
    #     freq = 0
    #     for target in range(min_val, max_val + 1):
    #         # Count elements already at target (no operation needed)
    #         at_target = 0
    #         # Count elements that can be transformed to target (need operation)
    #         transformable = 0

    #         for num in nums:
    #             if num == target:
    #                 at_target += 1
    #             elif abs(num - target) <= k:
    #                 transformable += 1

    #         # We can use up to numOperations to transform elements
    #         transformable = min(transformable, numOperations)

    #         # Total frequency = already at target + transformed to target
    #         freq = max(freq, at_target + transformable)

    #     return freq

    # Approach 2: Optimized with Sorting + Binary Search + Sliding Window
    # Time: O(n log n)
    # Space: O(n)
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        nums.sort()

        candidates = set(nums)
        for num in nums:
            candidates.add(num - k)
            candidates.add(num + k)

        counter = Counter(nums)
        freq = 0
        for target in sorted(candidates):
            at_target = counter.get(target, 0)

            lid = bisect.bisect_left(nums, target - k)
            rid = bisect.bisect_right(nums, target + k)

            total_in_range = rid - lid

            transformable = total_in_range - at_target
            transformable = min(transformable, numOperations)

            freq = max(freq, at_target + transformable)

        return freq

    def test(self):
        for nums, k, numOperations, expected in [
            ([1, 4, 5], 1, 2, 2),
            ([5, 11, 20, 20], 5, 1, 2),
        ]:
            output = self.maxFrequency(nums, k, numOperations)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
