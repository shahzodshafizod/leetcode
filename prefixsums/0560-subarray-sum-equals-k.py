from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/subarray-sum-equals-k/

# python3 -m unittest prefixsums/0560-subarray-sum-equals-k.py

class Solution(unittest.TestCase):
    # # Approach #1: Prefix Sum
    # # Time: O(nn)
    # # Space: O(1)
    # def subarraySum(self, nums: List[int], k: int) -> int:
    #     count, presum = 0, 0
    #     for end, num in enumerate(nums):
    #         presum += num
    #         subsum = presum
    #         for start in range(end+1):
    #             if subsum == k:
    #                 count += 1
    #             subsum -= nums[start]
    #     return count

    # Approach #2: Prefix Sum + Hash Table
    # Time: O(n)
    # Space: O(n)
    def subarraySum(self, nums: List[int], k: int) -> int:
        counter = defaultdict(int)
        counter[0] = 1
        count, presum = 0, 0
        for num in nums:
            presum += num
            count += counter[presum - k]
            counter[presum] += 1
        return count

    def test(self):
        for nums, k, expected in [
            ([1,1,1], 2, 2),
		    ([1,2,3], 3, 2),
        ]:
            output = self.subarraySum(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
