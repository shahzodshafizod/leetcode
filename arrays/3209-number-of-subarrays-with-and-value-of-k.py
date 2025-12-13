from typing import List
import unittest

# https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/

# python3 -m unittest arrays/3209-number-of-subarrays-with-and-value-of-k.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # Time: O(n^2) - check all n*(n+1)/2 subarrays, TLE for large n
    # # Space: O(1) - constant space
    # def countSubarrays(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     count = 0

    #     for i in range(n):
    #         and_val = nums[i]
    #         for j in range(i, n):
    #             if i == j:
    #                 and_val = nums[i]
    #             else:
    #                 and_val &= nums[j]

    #             if and_val == k:
    #                 count += 1

    #     return count

    # Approach #2: Optimized with AND Property
    # Time: O(n * log(max_val)) - for each position, update ~30 distinct values
    # Space: O(log(max_val)) - store distinct AND values
    def countSubarrays(self, nums: List[int], k: int) -> int:
        count = 0
        # prev_ands stores (and_value, count_of_subarrays_with_this_value) pairs
        prev_ands = {}

        for num in nums:
            # Start new subarray at current position
            curr_ands = {num: 1}

            # Extend all previous subarrays
            for and_val, cnt in prev_ands.items():
                new_and = and_val & num
                curr_ands[new_and] = curr_ands.get(new_and, 0) + cnt

            # Count subarrays with AND = k
            count += curr_ands.get(k, 0)

            prev_ands = curr_ands

        return count

    def test(self):
        for nums, k, expected in [
            ([1, 1, 1], 1, 6),  # All subarrays: [1], [1], [1], [1,1], [1,1], [1,1,1]
            ([1, 1, 2], 1, 3),  # [1], [1], [1,1]
            ([1, 2, 3], 2, 2),  # [2], [1,2]
            ([4, 5, 6], 4, 4),  # [4], [4,5], [5,6], [4,5,6]
            ([7, 7, 7], 7, 6),  # All subarrays
        ]:
            output = self.countSubarrays(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
