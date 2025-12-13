from typing import List
import unittest

# https://leetcode.com/problems/find-the-sum-of-the-power-of-all-subsequences/

# python3 -m unittest dp/3082-find-the-sum-of-the-power-of-all-subsequences.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force with Backtracking
    # # Try all subsequences, count those with sum = k, power = 2^length
    # # Time: O(2^n) - exponential, TLE for n > 20
    # # Space: O(n) - recursion depth
    # def sumOfPower(self, nums: List[int], k: int) -> int:
    #     MOD = 10**9 + 7
    #     n = len(nums)

    #     def dfs(i: int, current_sum: int) -> int:
    #         if current_sum == k:
    #             return pow(2, n - i, MOD)
    #         if current_sum > k or i >= n:
    #             return 0
    #         # Take element: adds to sum, Skip element: doubles the count (can be in/out of outer)
    #         return (2 * dfs(i + 1, current_sum) + dfs(i + 1, current_sum + nums[i])) % MOD

    #     return dfs(0, 0)

    # Approach #2: Space-Optimized DP
    # Key insight: When we select elements summing to k, remaining elements can be in/out
    # dp[v] = sum of 2^(remaining elements) for all ways to achieve sum v
    # For each element: dp[v] = 2*dp[v] + dp[v-a] (double existing + add new ways)
    # Time: O(n * k), Space: O(k)
    def sumOfPower(self, nums: List[int], k: int) -> int:
        MOD = 10**9 + 7
        dp = [0] * (k + 1)
        dp[0] = 1

        for num in nums:
            # Iterate backwards to avoid using updated values
            for v in range(k, -1, -1):
                # 2 * dp[v]: existing ways, element not included (can be in/out of outer)
                # dp[v - num]: new ways by including this element
                dp[v] = (2 * dp[v] + (dp[v - num] if v >= num else 0)) % MOD

        return dp[k]

    def test(self):
        for nums, k, expected in [
            ([1, 2, 3], 3, 6),
            ([2, 3, 3], 5, 4),
            ([1, 2, 3], 7, 0),
            ([15, 7, 9, 8, 22, 15, 16, 21, 9, 12, 8, 24, 12, 6, 12, 18, 4, 25, 14, 17, 1, 15, 1, 5, 18], 15, 148897792),
            ([14, 19, 7, 24, 21, 4, 21, 5, 20, 2, 3, 17, 8, 3, 16, 3, 24, 7, 10, 24, 9, 11, 24, 11, 19], 24, 487587840),
        ]:
            output = self.sumOfPower(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
