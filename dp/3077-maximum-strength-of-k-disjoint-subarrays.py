from typing import List
import unittest

# https://leetcode.com/problems/maximum-strength-of-k-disjoint-subarrays/

# python3 -m unittest dp/3077-maximum-strength-of-k-disjoint-subarrays.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # Try all possible ways to select k disjoint subarrays via backtracking.
    # # Generate all combinations of k disjoint subarrays and compute their strength.
    # # Time: O(n^(2k)) - exponential in k, TLE for k > 3
    # # Space: O(k) - recursion depth
    # def maximumStrength(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     max_strength = float('-inf')

    #     def calc_strength(selections: List[tuple[int, int, int]]) -> int:
    #         total = 0
    #         for subarray_num, start, end in selections:
    #             multiplier = (k - subarray_num + 1) * (1 if subarray_num % 2 == 1 else -1)
    #             subarray_sum = sum(nums[idx] for idx in range(start, end + 1))
    #             total += multiplier * subarray_sum
    #         return total

    #     def backtrack(pos: int, selected: List[tuple[int, int, int]], count: int):
    #         nonlocal max_strength

    #         if count == k:
    #             max_strength = max(max_strength, calc_strength(selected))
    #             return

    #         if pos >= n:
    #             return

    #         # Try all possible subarrays starting from pos
    #         for start in range(pos, n):
    #             for end in range(start, n):
    #                 backtrack(end + 1, selected + [(count + 1, start, end)], count + 1)

    #     backtrack(0, [], 0)
    #     return int(max_strength) if max_strength != float('-inf') else 0

    # Approach #2: Dynamic Programming (Kadane-like for each level)
    # dp[i][j] = max strength selecting i subarrays from first j elements
    # For each subarray level i, use Kadane's algorithm to find max weighted subarray sum
    # multiplier for i-th subarray = (k - i + 1) if i is odd, -(k - i + 1) if i is even
    # Time: O(n * k), Space: O(n * k)
    def maximumStrength(self, nums: List[int], k: int) -> int:
        n = len(nums)
        INF = float('-inf')

        # dp[i][j] = max strength selecting i subarrays from first j elements
        dp = [[INF] * (n + 1) for _ in range(k + 1)]
        for j in range(n + 1):
            dp[0][j] = 0

        for i in range(1, k + 1):
            # multiplier for i-th subarray: (k - i + 1) if i is odd, else -(k - i + 1)
            multiplier = (k - i + 1) if i % 2 == 1 else -(k - i + 1)
            max_sum = INF
            current = INF

            for j in range(i - 1, n):
                # current = max sum ending at position j for subarray level i
                # Either extend previous subarray or start new from dp[i-1][j]
                current = max(current + nums[j] * multiplier, dp[i - 1][j] + nums[j] * multiplier)
                max_sum = max(max_sum, current)
                dp[i][j + 1] = max_sum

        return int(dp[k][n]) if dp[k][n] > INF else 0

    def test(self):
        for nums, k, expected in [
            ([1, 2, 3, -1, 2], 3, 22),
            ([12, -2, -2, -2, -2], 5, 64),
            ([-1, -2, -3], 1, -1),
            ([-99, 85], 1, 85),
        ]:
            output = self.maximumStrength(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
