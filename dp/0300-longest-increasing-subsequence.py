from typing import List
import unittest

# https://leetcode.com/problems/longest-increasing-subsequence/

# python3 -m unittest dp/0300-longest-increasing-subsequence.py

class Solution(unittest.TestCase):
    # # Approach: Dynamic Programming (Bottom-Up)
    # # Time: O(nn)
    # # Space: O(n)
    # def lengthOfLIS(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     dp = [1] * n
    #     length = 1
    #     for curr in range(n):
    #         for prev in range(curr):
    #             if nums[prev] < nums[curr]:
    #                 dp[curr] = max(dp[curr], 1+dp[prev])
    #         length = max(length, dp[curr])
    #     return length

    # Approach: Binary Search
    # Time: O(nlogn)
    # Space: O(n)
    def lengthOfLIS(self, nums: List[int]) -> int:
        lis = []
        for num in nums:
            if not lis or num > lis[-1]:
                lis.append(num)
                continue
            left, right = 0, len(lis)
            while left < right:
                mid = left + (right-left) // 2
                if lis[mid] < num:
                    left = mid + 1
                else:
                    right = mid
            # the current element (num) is either equal to or smaller than
            # the existing element (lis[left]), and this allows for potentially
            # more elements to be added to the subsequence in the future.
            lis[left] = num
        return len(lis)
    
    def test(self) -> None:
        for nums, expected in [
            ([10,9,2,5,3,7,101,18], 4),
            ([0,1,0,3,2,3], 4),
            ([7,7,7,7,7,7,7], 1),
            ([3,1,8,2,5], 3),
            ([5,2,8,6,3,6,9,5], 4),
            ([4,10,4,3,8,9], 3),
            ([0,1,0,3,2,3], 4),
            ([3,5,6,2,5,4,19,5,6,7,12], 6),
            ([1,4,2,3], 3),
        ]:
            output = self.lengthOfLIS(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
