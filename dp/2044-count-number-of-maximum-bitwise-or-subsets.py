# from collections import defaultdict
# from copy import deepcopy
from typing import List
import unittest

# https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/

# python3 -m unittest dp/2044-count-number-of-maximum-bitwise-or-subsets.py

# The maximum bitwise-OR is the bitwise-OR of the whole array.


class Solution(unittest.TestCase):
    # Approach #1: Dynamic Programming (Top-Down)
    # Caching is inefficient, because the input length is much less:
    # w/ cache: n*131071 = 16*131071 = 2_097_136
    # w/o cache: 2^n = 2^16 = 65_536
    # Time: O(2^n)
    # Space: O(2^n)
    def countMaxOrSubsets(self, nums: List[int]) -> int:
        maxor = 0
        for num in nums:
            maxor |= num
        n = len(nums)

        # bin(1e5) = 11000011010100000
        # dec(11111111111111111) = 131071
        def dfs(idx: int, logicsum: int) -> int:
            if idx == n:
                nonlocal maxor
                return int(logicsum == maxor)
            return dfs(idx + 1, logicsum) + dfs(idx + 1, logicsum | nums[idx])

        return dfs(0, 0)

    # # Approach #2: Dynamic Programming (Bottom-Up)
    # # Time: O(n x maxor), maxor=2^n
    # # Space: O(maxor)
    # def countMaxOrSubsets(self, nums: List[int]) -> int:
    #     dp: dict[int, int] = defaultdict(int)
    #     dp[0] = 1
    #     maxor = 0
    #     for curr in nums:
    #         maxor |= curr
    #         new_dp = deepcopy(dp)
    #         for prev, cnt in dp.items():
    #             new_dp[curr | prev] += cnt
    #         dp = new_dp
    #     return dp[maxor]

    # # Approach #3: Bit Manipulation
    # # Time: O(n x maxor), maxor=2^n
    # # Space: O(1)
    # def countMaxOrSubsets(self, nums: List[int]) -> int:
    #     maxor = 0
    #     for num in nums:
    #         maxor |= num
    #     n = len(nums)
    #     count = 0
    #     for subset in range(1, 2**n):
    #         logicsum = 0
    #         for idx in range(n):
    #             if (1 << idx) & subset:
    #                 logicsum |= nums[idx]
    #         if logicsum == maxor:
    #             count += 1
    #     return count

    def test(self) -> None:
        for nums, expected in [
            ([3, 1], 2),
            ([2, 2, 2], 7),
            ([3, 2, 1, 5], 6),
            ([12007], 1),
            ([39317, 28735, 71230, 59313, 19954], 5),
            ([35569, 91997, 54930, 66672, 12363], 6),
            # ([32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30], 57083),
            # ([63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115], 3772),
            # ([6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890], 65535),
            # ([89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878], 52940),
            # ([90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649], 60072),
        ]:
            output = self.countMaxOrSubsets(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
