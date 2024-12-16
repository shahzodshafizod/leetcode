from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/subarray-sums-divisible-by-k/

# python3 -m unittest hashes/0974-subarray-sums-divisible-by-k.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(N^2)
    # # Space: O(1)
    # def subarraysDivByK(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     count = 0
    #     for left in range(n):
    #         curr_sum = 0
    #         for right in range(left, n):
    #             curr_sum = (curr_sum + nums[right]) % k
    #             if curr_sum == 0:
    #                 count += 1
    #     return count

    # Approach: Hash Table
    # Time: O(N)
    # Space: O(K)
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        count, prefmod = 0, 0
        mods = defaultdict(int)
        mods[0] = 1 # 0 is divisable by k
        for num in nums:
            prefmod = (prefmod + num + k) % k
            count += mods[prefmod]
            mods[prefmod] += 1
        return count

    def test(self):
        for nums, k, expected in [
            ([4,5,0,-2,-3,1], 5, 7),
            ([5], 9, 0),
            ([23,2,6,4,7], 6, 3),
            ([23,2,6,4,7], 13, 0),
            ([4,5,0,-2,-3,1], 5, 7),
            ([5], 9, 0),
            ([-5], 5, 1),
            ([0], 2, 1),
        ]:
            output = self.subarraysDivByK(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
