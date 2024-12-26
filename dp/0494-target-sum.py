from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/target-sum/

# python3 -m unittest dp/0494-target-sum.py

class Solution(unittest.TestCase):
    # # Approach #1: Backtracking
    # # Time: O(2^n)
    # # Space: O(n), for recursion stack
    # def findTargetSumWays(self, nums: List[int], target: int) -> int:
    #     n = len(nums)
    #     def backtrack(idx: int, result: int) -> int:
    #         if idx == n: return int(result == target)
    #         return (
    #             backtrack(idx+1, result+nums[idx]) +
    #             backtrack(idx+1, result-nums[idx])
    #         )
    #     return backtrack(0, 0)

    # # Approach #2: Backtracking with Memoization
    # # Time: O(n*sum(nums))
    # # Space: O(n*sum(nums))
    # def findTargetSumWays(self, nums: List[int], target: int) -> int:
    #     n = len(nums)
    #     memo = [{} for _ in range(n)]
    #     def backtrack(idx: int, result: int) -> int:
    #         if idx == n: return int(result == target)
    #         if result in memo[idx]: return memo[idx][result]
    #         memo[idx][result] = (
    #             backtrack(idx+1, result+nums[idx]) +
    #             backtrack(idx+1, result-nums[idx])
    #         )
    #         return memo[idx][result]
    #     return backtrack(0, 0)

    # # Approach #3: Bottom-Up Dynamic Programming
    # # Time: O(n*sum(nums))
    # # Space: O(n*sum(nums))
    # def findTargetSumWays(self, nums: List[int], target: int) -> int:
    #     n = len(nums)
    #     dp = [defaultdict(int) for _ in range(n+1)]
    #     dp[0][0] = 1
    #     for idx in range(n):
    #         for curr_sum, count in dp[idx].items():
    #             dp[idx+1][curr_sum+nums[idx]] += count
    #             dp[idx+1][curr_sum-nums[idx]] += count
    #     return dp[n][target]

    # Approach #4: Bottom-Up Dynamic Programming (Space Optimized)
    # Time: O(n*sum(nums))
    # Space: O(sum(nums))
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        dp = defaultdict(int)
        dp[0] = 1
        for idx in range(len(nums)):
            prev = dp
            dp = defaultdict(int)
            for curr_sum, count in prev.items():
                dp[curr_sum+nums[idx]] += count
                dp[curr_sum-nums[idx]] += count
        return dp[target]

    def test(self):
        for nums, target, expected in [
            ([1], 1, 1),
            ([1], 2, 0),
            ([100, 100], -300, 0),
            ([1,1,1,1,1], 3, 5),
            # ([12,25,42,49,41,15,22,34,28,31], 35, 8),
            # ([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1], 1, 524288),
            # ([3,2,3,5,7,11,13,17,19,23,29,2,7,9,13,27,31,37,47,53], 107, 0),
            # ([0,5,22,39,30,5,40,23,47,43,19,36,10,28,46,14,11,5,50,41], 48, 5726),
            # ([30,29,48,14,29,4,31,12,40,13,30,1,5,32,16,17,13,20,41,38], 9, 6867),
            # ([3,2,3,5,7,11,13,17,19,23,29,2,107,109,113,127,131,137,47,53], 4, 2780),
        ]:
            output = self.findTargetSumWays(nums, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
