from typing import List
import unittest

# https://leetcode.com/problems/burst-balloons/

# python3 -m unittest dp/0312-burst-balloons.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(N^3)
    # # Space: O(N^2)
    # def maxCoins(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     dp = [[None] * n for _ in range(n)]
    #     def dfs(prev: int, left: int, right: int, next: int) -> int:
    #         if left > right:
    #             return 0
    #         if dp[left][right] != None:
    #             return dp[left][right]
    #         dp[left][right] = 0
    #         for idx in range(left, right+1):
    #             coins = prev*nums[idx]*next
    #             coins += dfs(prev, left, idx-1, nums[idx])
    #             coins += dfs(nums[idx], idx+1, right, next)
    #             dp[left][right] = max(dp[left][right], coins)
    #         return dp[left][right]
    #     return dfs(1, 0, n-1, 1)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(N^3)
    # Space: O(N^2)
    def maxCoins(self, nums: List[int]) -> int:
        nums = [1] + [num for num in nums if num != 0] + [1]
        n = len(nums)
        dp = [[0] * n for _ in range(n)]
        for left in range(n-2, 0, -1):
            for right in range(left, n-1):
                for idx in range(left, right+1):
                    dp[left][right] = max(
                        dp[left][right],
                        nums[left-1]*nums[idx]*nums[right+1]
                        + dp[left][idx-1]
                        + dp[idx+1][right]
                    )
        # [0] and [n-1] are not part of the nums
        return dp[1][n-2]

    def test(self):
        for nums, expected in [
            ([3,1,5,8], 167),
            ([1,5], 10),
            ([9, 76, 64, 21, 97, 60], 1086136),
        ]:
            output = self.maxCoins(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
