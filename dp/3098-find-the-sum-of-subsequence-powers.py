from typing import List
import unittest

# https://leetcode.com/problems/find-the-sum-of-subsequence-powers/

# python3 -m unittest dp/3098-find-the-sum-of-subsequence-powers.py


class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming
    # Time: O(nnn k d), n=len(nums), d=diff
    # Space: O(nkd)
    def sumOfPowers(self, nums: List[int], k: int) -> int:
        nums.sort()
        memo = {}
        MOD = 10**9 + 7
        n = len(nums)
        def dp(idx: int, k: int, diff: int) -> int:
            if k == 0:
                return diff
            key = (idx, k, diff)
            if key in memo:
                return memo[key]
            total = 0
            for nxt in range(idx+1, n):
                total = (total + dp(nxt, k-1, min(diff, nums[nxt] - nums[idx]))) % MOD
            memo[key] = total
            return total
        return sum(dp(i, k-1, (1 << 32) - 1) for i in range(n)) % MOD

    def test(self):
        for nums, k, expected in [
            ([1,2,3,4], 3, 4),
            ([2,2], 2, 0),
            ([4,3,-1], 2, 10),
            # ([-24, -921, 119, -291, -65, -628, 372, 274, 962, -592, -10, 67, -977, 85, -294, 349, -119, -846, -959, -79, -877, 833, 857, 44, 826, -295, -855, 554, -999, 759, -653, -423, -599, -928], 19, 990202285),
        ]:
            output = self.sumOfPowers(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
