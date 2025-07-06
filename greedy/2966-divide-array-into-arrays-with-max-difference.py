from typing import List
import unittest

# https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/

# python3 -m unittest greedy/2966-divide-array-into-arrays-with-max-difference.py


class Solution(unittest.TestCase):
    # Approach: Sorting
    # Time: O(nlogn)
    # Space: O(n)
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        nums.sort()
        parts = []
        for idx in range(2, len(nums), 3):
            if nums[idx] - nums[idx - 2] > k:
                return []
            parts.append(nums[idx - 2 : idx + 1])
        return parts

    def test(self):
        for nums, k, expected in [
            ([1, 3, 4, 8, 7, 9, 3, 5, 1], 2, [[1, 1, 3], [3, 4, 5], [7, 8, 9]]),
            ([2, 4, 2, 2, 5, 2], 2, []),
            (
                [4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11],
                14,
                [[2, 2, 2], [4, 5, 5], [5, 5, 7], [7, 8, 8], [9, 9, 10], [11, 12, 12]],
            ),
        ]:
            output = self.divideArray(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
