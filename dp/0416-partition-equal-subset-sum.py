from typing import List
import unittest

# https://leetcode.com/problems/partition-equal-subset-sum/

# python3 -m unittest dp/0416-partition-equal-subset-sum.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n*sum)
    # # Space: O(n*sum)
    # def canPartition(self, nums: List[int]) -> bool:
    #     total = sum(nums)
    #     if total&1:
    #         return False
    #     n = len(nums)
    #     half = total // 2
    #     memo = [[None] * (half+1) for _ in range(n)]
    #     def dfs(idx: int, total: int) -> bool:
    #         if idx == n or total <= 0:
    #             return total == 0
    #         if memo[idx][total] != None:
    #             return memo[idx][total]
    #         memo[idx][total] = (
    #             dfs(idx+1, total) or
    #             dfs(idx+1, total-nums[idx])
    #         )
    #         return memo[idx][total]
    #     return dfs(0, half)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n*sum)
    # Space: O(sum)
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total & 1:
            return False
        target = total // 2
        dp = set([0])
        for num in nums:
            next_dp = dp.copy()
            for t in dp:
                if t + num == target:
                    return True
                next_dp.add(t + num)
            dp = next_dp
        return target in dp

    def test(self):
        for nums, expected in [
            ([1, 5, 11, 5], True),
            ([1, 2, 3, 5], False),
        ]:
            output = self.canPartition(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
