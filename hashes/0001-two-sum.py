import unittest
from typing import List

# https://leetcode.com/problems/two-sum/

# python3 -m unittest hashes/0001-two-sum.py

class Solution(unittest.TestCase):
    # # Approach: Two Pointers
    # # Time: O(n log n)
    # # Space: O(1)
    # def twoSum(self, nums: List[int], target: int) -> List[int]:
    #     indices = [idx for idx in range(len(nums))]
    #     indices.sort(key=lambda key: nums[key])
    #     lo, hi = 0, len(nums)-1
    #     while lo < hi:
    #         total = nums[indices[lo]]+nums[indices[hi]]
    #         if total > target:
    #             hi -= 1
    #         elif total < target:
    #             lo += 1
    #         else:
    #             return [indices[lo],indices[hi]]
    #     return []

    # Approach: Hash Table
    # Time: O(n)
    # Space: O(n)
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diffs = {}
        for idx in range(len(nums)):
            if nums[idx] in diffs:
                return [diffs[nums[idx]],idx]
            diffs[target-nums[idx]] = idx
        return []

    def test(self):
        for nums, target, expected in [
            ([1,3,7,9,2],11,[3,4]),
            ([1,3,7,9,2], 25, []),
            ([], 1, []),
            ([5], 5, []),
            ([1,6], 7, [0,1]),
            ([],1, []),
            ([2,7,11,15], 9, [0,1]),
            ([3,2,4], 6, [1,2]),
            ([3,3], 6, [0,1]),
        ]:
            output = self.twoSum(nums, target)
            output.sort()
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
