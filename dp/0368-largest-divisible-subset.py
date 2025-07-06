from typing import List
import unittest

# https://leetcode.com/problems/largest-divisible-subset/

# python3 -m unittest dp/0368-largest-divisible-subset.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
    #     nums.sort()
    #     n = len(nums)
    #     memo = [None] * n
    #     def dp(idx: int) -> List[int]:
    #         if idx == n: return []
    #         if memo[idx] != None:
    #             return memo[idx]
    #         answer = []
    #         for j in range(idx+1, n):
    #             if nums[j] % nums[idx] == 0:
    #                 tmp = dp(j)
    #                 if len(tmp) > len(answer):
    #                     answer = tmp
    #         memo[idx] = [nums[idx]] + answer
    #         return memo[idx]
    #     answer = []
    #     for idx in range(n):
    #         tmp = dp(idx)
    #         if len(tmp) > len(answer):
    #             answer = tmp
    #     return answer

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(nn)
    # Space: O(nn)
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        nums.sort()
        n = len(nums)
        dp = [[num] for num in nums]
        answer = []
        for idx in range(n - 1, -1, -1):
            for j in range(idx + 1, n):
                if nums[j] % nums[idx] == 0 and len(dp[j]) + 1 > len(dp[idx]):
                    dp[idx] = [nums[idx]] + dp[j]
            answer = dp[idx] if len(dp[idx]) > len(answer) else answer
        return answer

    def test(self):
        for nums, expected in [
            ([1, 2, 3], [1, 2]),
            ([1, 2, 4, 8], [1, 2, 4, 8]),
        ]:
            output = self.largestDivisibleSubset(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
