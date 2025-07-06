from functools import lru_cache  # pylint: disable=unused-import
from typing import List
import unittest

# https://leetcode.com/problems/split-array-with-same-average/

# python3 -m unittest dp/0805-split-array-with-same-average.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
    # # Time: O(2^N)
    # # Space: O(2^N)
    # def splitArraySameAverage(self, nums: List[int]) -> bool:
    #     n = len(nums)
    #     total = sum(nums)
    #     @lru_cache(None)
    #     def dfs(idx: int, sum1: int, len1: int) -> bool:
    #         if idx == n:
    #             # sum1/len1 = sum2/len2
    # 		    # sum1*len2 = sum2*len1
    #             # len2 = n - len1
    #             # sum2 = total - sum1
    #             return 0 < len1 < n and sum1*(n-len1) == (total-sum1)*len1
    #         # decision to include
    #         if dfs(idx+1, sum1+nums[idx], len1+1):
    #             return True
    #         # decision NOT to include
    #         return dfs(idx+1, sum1, len1)
    #     return dfs(0, 0, 0)

    # # Approach #2: Top-Down Dynamic Programming
    # # Time: O(N^3)
    # # Space: O(N^2 * TotalSum)
    # def splitArraySameAverage(self, nums: List[int]) -> bool:
    #     total_sum, n = sum(nums), len(nums)
    #     @lru_cache(None)
    #     def dfs(start: int, size: int, target: int) -> bool:
    #         if size == 0 and target == 0:
    #             return True
    #         if start == n or size == 0 or target < 0:
    #             return False
    #         return dfs(start+1, size, target) or \
    #             dfs(start+1, size-1, target-nums[start])
    #     for size in range(1, n//2+1):
    #         if total_sum*size%n == 0:
    #             if dfs(0, size, total_sum*size//n):
    #                 return True
    #     return False

    # # Approach #3: Bottom-Up Dynamic Programming
    # # Time: O(N^2 * TotalSum)
    # # Space: O(N * TotalSum)
    # def splitArraySameAverage(self, nums: List[int]) -> bool:
    #     n = len(nums)
    #     total_sum = sum(nums)
    #     possible = False
    #     for size in range(1, n):
    #         if total_sum*size%n == 0:
    #             possible = True
    #             break
    #     if not possible:
    #         return False
    #     dp = [set() for _ in range(n//2+1)]
    #     dp[0].add(0)
    #     for num in nums:
    #         for size in range(n//2, 0, -1):
    #             for target_sum in dp[size-1]:
    #                 dp[size].add(target_sum + num)
    #     for size in range(1, n//2+1):
    #         if total_sum*size%n == 0:
    #             target_sum = total_sum*size//n
    #             if target_sum in dp[size]:
    #                 return True
    #     return False

    # Approach #4: Bit Manipulation
    # Time: O(N*TotalSum)
    # Space: O(TotalSum)
    def splitArraySameAverage(self, nums: List[int]) -> bool:
        n = len(nums)
        if n == 1:
            return False
        total = sum(nums)
        if total == 0:
            return True
        possible = False
        for size in range(1, n):
            if total * size % n == 0:
                possible = True
                break
        if not possible:
            return False
        dp = [0] * (total + 1)
        dp[0] = 1
        for num in nums:
            for target in range(total, -1, -1):
                if dp[target]:
                    dp[target + num] |= dp[target] << 1
        for target in range(1, total + 1):
            if target * n % total == 0:
                shift = target * n // total
                if shift and shift < n and (dp[target] & (1 << shift)):
                    return True
        return False

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4, 5, 6, 7, 8], True),
            ([3, 1], False),
            ([17, 5, 5, 1, 14, 10, 13, 1, 6], True),
            ([3, 4, 9, 4, 4, 3, 9, 8, 5, 3], True),
            ([5, 16, 4, 11, 4], True),
            ([2, 12, 18, 16, 19, 3], False),
            ([5, 3, 11, 19, 2], True),
            ([1, 6, 1], False),
            ([1], False),
            ([2], False),
            ([0, 0, 0, 0, 0], True),
            ([0], False),
            # ([3322,959,598,506,4442,3594,8382,7777,7002,7078,2278,4902,1276,1], False),
            # ([904,8738,6439,1889,138,5771,8899,5790,662,8402,3074,1844,5926,8720,7159,6793,7402,9466,1282,1748,434,842,22], False),
            # ([60,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30], False),
            # ([6795,3310,8624,475,7609,7858,7086,8934,6197,2431,3310,760,1432,7518,7068,7182,2681,2679,6461,9928,9651,3258,9346,1666,5400,8384,2751,1234,2183,3520], False),
        ]:
            output = self.splitArraySameAverage(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
