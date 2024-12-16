from typing import Optional
from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/arithmetic-slices-ii-subsequence/

# python3 -m unittest dp/0446-arithmetic-slices-ii-subsequence.py

"""
+-------------+---+---+---+---+-----+-------+
| diff | nums | 0 | 1 | 2 | 2 | 3   |     4 |
|-------------+---+---+---+---+-----+-------|
| 0           |   |   |   | 1 |     |       |
| 1           |   | 1 | 2 | 2 | 6   |     7 |
| 2           |   |   | 1 | 1 | 1   |     4 |
| 3           |   |   |   |   | 1   |     1 |
| 4           |   |   |   |   |     |     1 |
+-------------+---+---+---+---+-----+-------+
| count       |   | 0 | 1 | 1 | 2+2 | 6+1+1 |
| result      |   | 0 | 1 | 2 |  6  |  "14" |
+-------------+---+---+---+---+-----+-------+
"""

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(nnn), n=len(nums)
    # # Space: O(nd), d=max(nums)-min(nums)
    # def numberOfArithmeticSlices(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     dp = [defaultdict(int) for _ in range(n)]
    #     def dfs(prev: int, curr: int) -> int:
    #         difference = nums[curr] - nums[prev]
    #         if difference in dp[curr]:
    #             return dp[curr][difference]
    #         count = 0
    #         for next in range(curr+1, n):
    #             if nums[next] - nums[curr] == difference:
    #                 count += 1 + dfs(curr, next)
    #         dp[curr][difference] = count
    #         return count
    #     count = 0
    #     for prev in range(n):
    #         for curr in range(prev+1, n):
    #             count += dfs(prev, curr)
    #     return count

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(nn), n=len(nums)
    # Space: O(nd), d=max(nums)-min(nums)
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        n = len(nums)
        dp = [defaultdict(int) for _ in range(n)]
        count = 0
        for curr in range(n):
            for prev in range(curr):
                difference = nums[curr] - nums[prev]
                dp[curr][difference] += dp[prev][difference] + 1
                count += dp[prev][difference]
        return count

    def test(self):
        for nums, expected in [
            ([-104], 0),
		    ([7,7,7,7,7], 16),
            ([2,4,6,8,10], 7),
            ([0,2000000000,-294967296], 0),
            ([211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211,211], 16776915),
            ([-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362,-454,-908,-1362], 3244),
            ([1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536,131072,262144,524288,1048576,2097152,4194304,8388608,16777216,33554432,67108864,134217728,268435456,536870912,1,3,9,27,81,243,729,2187,6561,19683], 1),
            ([1073741824,536870912,268435456,134217728,67108864,33554432,16777216,8388608,4194304,2097152,1048576,524288,262144,131072,65536,32768,16384,8192,4096,2048,1024,512,256,128,64,32,16,8,4,2,1073740824,536870412,268435206,134217603,67108801,33554400,16777200,8388600,4194300,2097150,1048575,524287,262143,131071,65535,32767,16383,8191,4095,2047,1023,511,255,127,63,31,15,7,3,1], 0),
        ]:
            output = self.numberOfArithmeticSlices(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
