from typing import List, Dict
import unittest
from math import gcd
from functools import reduce


# https://leetcode.com/problems/minimum-increments-for-target-multiples-in-an-array/

# python3 -m unittest dp/3444-minimum-increments-for-target-multiples-in-an-array.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force with Bitmask DP
    # # Time: O(n * 3^m * MAX_COST) - TLE due to MAX_COST loop
    # # Space: O(2^m)
    # def minOperations(self, nums: List[int], target: List[int]) -> int:
    #     m = len(target)
    #     INF = 1 << 31
    #     MAX_COST = 10001

    #     size = 1 << m
    #     dp = [INF] * size
    #     dp[0] = 0
    #     nums.sort()

    #     for num in nums:
    #         new_dp = dp[:]
    #         for mask in range(size):
    #             if dp[mask] == INF:
    #                 continue

    #             uncovered = (size - 1) ^ mask
    #             subset = uncovered
    #             while subset > 0:
    #                 # Try all costs from 0 to MAX_COST (THIS CAUSES TLE)
    #                 min_cost = INF
    #                 for cost in range(MAX_COST):
    #                     val = num + cost
    #                     if all((subset & (1 << i) == 0) or (val % target[i] == 0) for i in range(m)):
    #                         min_cost = cost
    #                         break

    #                 if min_cost != INF:
    #                     new_mask = mask | subset
    #                     new_dp[new_mask] = min(new_dp[new_mask], dp[mask] + min_cost)

    #                 subset = (subset - 1) & uncovered

    #         dp = new_dp

    #     return dp[size - 1]

    # Approach #2: Optimized DP with LCM Calculation
    # Time: O(n * 3^m * m) where 3^m comes from subset iteration
    # Space: O(2^m)
    def minOperations(self, nums: List[int], target: List[int]) -> int:
        m = len(target)
        INF = 1 << 31

        # Precompute LCM for each non-empty subset of targets
        lcm_map: Dict[int, int] = {}
        for mask in range(1, 1 << m):
            targets_subset = [target[i] for i in range(m) if mask & (1 << i)]
            lcm_val = reduce(lambda a, b: a * b // gcd(a, b), targets_subset)
            lcm_map[mask] = lcm_val

        size = 1 << m
        dp = [INF] * size
        dp[0] = 0
        nums.sort()

        for num in nums:
            new_dp = dp[:]
            for mask in range(1 << m):
                if dp[mask] == INF:
                    continue

                uncovered = ((1 << m) - 1) ^ mask
                subset = uncovered
                while subset > 0:
                    # Calculate exact cost using LCM
                    lcm_val = lcm_map[subset]
                    remainder = num % lcm_val
                    cost = 0 if remainder == 0 else lcm_val - remainder

                    new_mask = mask | subset
                    new_dp[new_mask] = min(new_dp[new_mask], dp[mask] + cost)

                    subset = (subset - 1) & uncovered

            dp = new_dp

        return dp[(1 << m) - 1] if dp[(1 << m) - 1] != INF else -1

    def test(self):
        for nums, target, expected in [
            ([1, 2, 3], [4], 1),
            ([8, 4], [10, 5], 2),
            ([7, 9, 10], [7], 0),
            ([1, 2, 8], [5, 10], 2),
            ([3, 5, 7], [6], 1),
        ]:
            output = self.minOperations(nums, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
