from typing import List
import unittest

# https://leetcode.com/problems/divide-array-into-equal-pairs/

# python3 -m unittest hashes/2206-divide-array-into-equal-pairs.py


class Solution(unittest.TestCase):
    # # Approach #2: Sorting
    # # Time: O(nlogn)
    # # Space: O(n), Timsort algorithm has a space complexity of O(2n)=O(n)
    # def divideArray(self, nums: List[int]) -> bool:
    #     prev, count = -1, 0
    #     for num in sorted(nums):
    #         if prev == num:
    #             count += 1
    #         else:
    #             if count&1:
    #                 break
    #             count = 1
    #         prev = num
    #     return count&1 == 0

    # Approach #1: Hashes, Counting
    # Time: O(n)
    # Space: O(n)
    def divideArray(self, nums: List[int]) -> bool:
        count = {}
        for num in nums:
            count[num] = 1 + count.get(num, 0)
        return all(cnt & 1 == 0 for cnt in count.values())

    def test(self):
        for nums, expected in [
            ([3, 2, 3, 2, 2, 2], True),
            ([1, 2, 3, 4], False),
            ([1, 2, 4, 7], False),
        ]:
            output = self.divideArray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
