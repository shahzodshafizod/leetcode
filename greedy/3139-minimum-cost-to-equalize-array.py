from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-to-equalize-array/

# python3 -m unittest greedy/3139-minimum-cost-to-equalize-array.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force (Enumeration with Greedy Pairing)
    # # For each possible target value, calculate cost using greedy pairing strategy
    # # Time: O(n * max_val) - enumerate many targets, for each calculate cost
    # # Space: O(1)
    # def minCostToEqualizeArray(self, nums: List[int], cost1: int, cost2: int) -> int:
    #     MOD = 10**9 + 7
    #     n = len(nums)
    #
    #     if n == 1:
    #         return 0
    #
    #     max_val = max(nums)
    #     min_val = min(nums)
    #     min_cost = 10**18
    #
    #     # Try different target values
    #     # Need to check beyond max_val if pairing is beneficial
    #     upper_bound = 2 * max_val - min_val
    #     for target in range(max_val, upper_bound + 1):
    #         total = sum(target - num for num in nums)
    #         max_inc = max(target - num for num in nums)
    #
    #         if 2 * cost1 <= cost2 or n <= 2:
    #             # Always use cost1
    #             cost = total * cost1
    #         else:
    #             # Try to use cost2
    #             other_inc = total - max_inc
    #             if max_inc > other_inc:
    #                 # Can't pair everything for the max element
    #                 cost = other_inc * cost2 + (max_inc - other_inc) * cost1
    #             else:
    #                 # Can pair everything
    #                 cost = (total // 2) * cost2 + (total % 2) * cost1
    #
    #         min_cost = min(min_cost, cost)
    #
    #     return min_cost % MOD

    # Approach #2: Greedy with 4 Cases Analysis (Optimized)
    # Based on lee215's solution: analyze 4 distinct cases
    # Case 1: cost1*2 <= cost2 or n <= 2 → always use cost1
    # Case 2: Use cost2 as much as possible up to max_val
    # Case 3: Go beyond max_val to reduce op1 usage
    # Case 4: Go even further to eliminate op1 completely
    # Time: O(n), Space: O(1)
    def minCostToEqualizeArray(self, nums: List[int], cost1: int, cost2: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)

        if n == 1:
            return 0

        ma, mi = max(nums), min(nums)
        total = ma * n - sum(nums)  # Total increments needed to reach max

        # Case 1: cost1 is always better, or array too small for pairing
        if cost1 * 2 <= cost2 or n <= 2:
            return (total * cost1) % MOD

        # Case 2: Increment to max(nums), use cost2 as much as possible
        op1 = max(0, (ma - mi) * 2 - total)  # Increments that must use cost1
        op2 = total - op1  # Increments that can be paired
        res = (op1 + op2 % 2) * cost1 + op2 // 2 * cost2

        # Case 3: Increment beyond max to reduce op1
        # Each round increases total by n, but reduces op1 by (n-2)
        # Continue while (n-2)*cost1 > (n-1)*cost2
        if op1 > 0:
            rounds = op1 // (n - 2)
            total += rounds * n
            op1 %= n - 2
            op2 = total - op1
            res = min(res, (op1 + op2 % 2) * cost1 + op2 // 2 * cost2)

        # Case 4: Increment even further to make total even (eliminate all op1)
        for _ in range(2):
            total += n
            res = min(res, total % 2 * cost1 + total // 2 * cost2)

        return res % MOD

    def test(self):
        for nums, cost1, cost2, expected in [
            ([4, 1], 5, 2, 15),  # Make both 4: need 3 increments for 1->4, use pairs or singles
            ([2, 3, 3, 3, 5], 2, 1, 6),
            ([3, 5, 3], 1, 3, 4),
            ([1, 14, 14, 15], 2, 1, 20),
            ([1, 1000000], 1000000, 1000000, 998993007),
        ]:
            output = self.minCostToEqualizeArray(nums, cost1, cost2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
