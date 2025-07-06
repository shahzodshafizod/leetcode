from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/

# python3 -m unittest greedy/2948-make-lexicographically-smallest-array-by-swapping-elements.py


class Solution(unittest.TestCase):
    def lexicographicallySmallestArray(self, nums: List[int], limit: int) -> List[int]:
        groups = []
        group_indices = {}
        for num in sorted(nums):
            if not groups or abs(num - groups[-1][-1]) > limit:
                groups.append(deque())
            groups[-1].append(num)
            group_indices[num] = len(groups) - 1
        result = []
        for num in nums:
            result.append(groups[group_indices[num]].popleft())
        return result

    def test(self):
        for nums, limit, expected in [
            ([1, 5, 3, 9, 8], 2, [1, 3, 5, 8, 9]),
            ([1, 7, 6, 18, 2, 1], 3, [1, 6, 7, 18, 1, 2]),
            ([1, 7, 28, 19, 10], 3, [1, 7, 28, 19, 10]),
        ]:
            output = self.lexicographicallySmallestArray(nums, limit)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
