from typing import List
import unittest

# https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/

# python3 -m unittest twopointers/2570-merge-two-2d-arrays-by-summing-values.py


class Solution(unittest.TestCase):
    def mergeArrays(self, nums1: List[List[int]], nums2: List[List[int]]) -> List[List[int]]:
        res = []
        n1, n2 = len(nums1), len(nums2)
        idx1, idx2 = 0, 0
        while idx1 < n1 and idx2 < n2:
            if nums1[idx1][0] < nums2[idx2][0]:
                res.append(nums1[idx1])
                idx1 += 1
            elif nums2[idx2][0] < nums1[idx1][0]:
                res.append(nums2[idx2])
                idx2 += 1
            else:
                res.append([nums1[idx1][0], nums1[idx1][1] + nums2[idx2][1]])
                idx1 += 1
                idx2 += 1
        while idx1 < n1:
            res.append(nums1[idx1])
            idx1 += 1
        while idx2 < n2:
            res.append(nums2[idx2])
            idx2 += 1
        return res

    def test(self):
        for nums1, nums2, expected in [
            ([[1, 2], [2, 3], [4, 5]], [[1, 4], [3, 2], [4, 1]], [[1, 6], [2, 3], [3, 2], [4, 6]]),
            ([[2, 4], [3, 6], [5, 5]], [[1, 3], [4, 3]], [[1, 3], [2, 4], [3, 6], [4, 3], [5, 5]]),
            ([[1000, 1000]], [[999, 999]], [[999, 999], [1000, 1000]]),
            ([[1, 3], [2, 2], [3, 1]], [[1, 1], [2, 2], [3, 3]], [[1, 4], [2, 4], [3, 4]]),
            ([[1, 555], [3, 777], [5, 999]], [[2, 666], [4, 888]], [[1, 555], [2, 666], [3, 777], [4, 888], [5, 999]]),
            (
                [[998, 100], [999, 1000]],
                [[997, 500], [999, 1000], [1000, 1000]],
                [[997, 500], [998, 100], [999, 2000], [1000, 1000]],
            ),
        ]:
            output = self.mergeArrays(nums1, nums2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
