from typing import List
import unittest

# https://leetcode.com/problems/minimum-index-of-a-valid-split/

# python3 -m unittest hashes/2780-minimum-index-of-a-valid-split.py


class Solution(unittest.TestCase):
    # # Approach #1: Hash Map
    # # Time: O(n)
    # # Space: O(n)
    # def minimumIndex(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     count = defaultdict(int)
    #     dominent = -1
    #     for idx in range(n):
    #         count[nums[idx]] += 1
    #         if count[nums[idx]]*2 > n:
    #             dominent = nums[idx]
    #     prefix, postfix = 0, count[dominent]
    #     for idx in range(n):
    #         if nums[idx] == dominent:
    #             prefix += 1
    #             postfix -= 1
    #         if prefix*2 > idx+1 and postfix*2 > (n-idx-1):
    #             return idx
    #     return -1

    # Approach #2: Boyer-Moore Majority Voting Algorithm
    # Time: O(n)
    # Space: O(1)
    def minimumIndex(self, nums: List[int]) -> int:
        dominent = nums[0]
        count = 0
        for num in nums:
            count += 1 if num == dominent else -1
            if count == 0:
                dominent = num
                count = 1
        prefix, postfix = 0, sum(1 for num in nums if num == dominent)
        n = len(nums)
        for idx in range(n):
            if nums[idx] == dominent:
                prefix += 1
                postfix -= 1
            if prefix * 2 > idx + 1 and postfix * 2 > n - idx - 1:
                return idx
        return -1

    def test(self):
        for nums, expected in [
            ([1, 2, 2, 2], 2),
            ([2, 1, 3, 1, 1, 1, 7, 1, 2, 1], 4),
            ([3, 3, 3, 3, 7, 2, 2], -1),
        ]:
            output = self.minimumIndex(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
