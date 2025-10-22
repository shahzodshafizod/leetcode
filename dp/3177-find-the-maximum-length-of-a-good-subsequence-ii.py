from functools import lru_cache
from typing import List
import unittest

# https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-ii/

# python3 -m unittest dp/3177-find-the-maximum-length-of-a-good-subsequence-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (TLE)
    # # Time: O(nnk)
    # # Space: O(nk)
    # def maximumLength(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     memo = [[-1] * (k+1) for _ in range(n)]
    #     def dp(i: int, k: int) -> int:
    #         if i == n: return 0
    #         if memo[i][k] != -1:
    #             return memo[i][k]
    #         length = 1
    #         for j in range(i+1, n):
    #             if nums[j] == nums[i]:
    #                 length = max(length, 1 + dp(j, k))
    #             elif k >= 1:
    #                 length = max(length, 1 + dp(j, k-1))
    #         memo[i][k] = length
    #         return length
    #     return dp(0, k)

    # # Approach #2: Top-Down Dynamic Programming
    # # Time: O(nk)
    # # Space: O(nk)
    # def maximumLength(self, nums: List[int], k: int) -> int:
    #     from collections import defaultdict
    #     from bisect import bisect_right
    #     n = len(nums)
    #     indices = defaultdict(list)
    #     for i, num in enumerate(nums):
    #         indices[num].append(i)
    #     memo = {}
    #     def dp(i: int, k: int, same: bool) -> int:
    #         if i == n: return 0
    #         state = (i, k, same)
    #         if state in memo: return memo[state]
    #         v = indices[nums[i]]
    #         ind = bisect_right(v, i)
    #         take = 0
    #         # check if i was the last index of nums[i]
    #         if ind == len(v):
    #             take = 1
    #         else:
    #             take = 1 + dp(v[ind], k, True)
    #         if k > 0:
    #             take = max(take, 1 + dp(i+1, k-1, False))
    #         skip = 0 if same else dp(i+1, k, False)
    #         memo[state] = max(take, skip)
    #         return memo[state]
    #     return dp(0, k, False)

    # Approach #3: Bottom-Up Dynamic Programming (Space Optimized)
    # Time: O(nk)
    # Space: O(unique_nums * k)
    def maximumLength(self, nums: List[int], k: int) -> int:
        from collections import defaultdict
        dp = [defaultdict(int) for _ in range(k+1)]
        res = [0] * (k+1)
        for num in nums:
            for count in range(k, -1, -1):
                dp[count][num] += 1
                if count > 0:
                    dp[count][num] = max(dp[count][num], res[count-1] + 1)
                res[count] = max(res[count], dp[count][num])
        return res[k]

    def test(self):
        for nums, k, expected in [
            ([1,2,1,1,3], 2, 4),
            ([1,2,3,4,5,1], 0, 2),
        ]:
            output = self.maximumLength(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
