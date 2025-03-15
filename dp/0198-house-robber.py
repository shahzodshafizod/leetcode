from typing import List
import unittest

# https://leetcode.com/problems/house-robber/

# python3 -m unittest dp/0198-house-robber.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(nn)
    # # Space: O(n)
    # def rob(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     memo = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx >= n:
    #             return 0
    #         if memo[idx] != None:
    #             return memo[idx]
    #         memo[idx] = max(
    #             nums[idx] + dp(idx+2),
    #             dp(idx+1),
    #         )
    #         return memo[idx]
    #     return dp(0)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n)
    # Space: O(1)
    def rob(self, nums: List[int]) -> int:
        prev, curr = 0, 0
        for num in nums:
            prev, curr = curr, max(prev + num, curr)
        return curr

    def test(self):
        for nums, expected in [
			([1,2,3,1], 4),
			([2,7,9,3,1], 12),
			([2,1,1,2], 4),
			([2,1,7,9], 11),
			([100,10,1,10,100], 201),
			([1,3,5,9], 12),
			([2,4,6,8], 12),
			([40,2,4,10], 50),
			([350], 350),
			# ([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], 0),
		]:
            output = self.rob(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
