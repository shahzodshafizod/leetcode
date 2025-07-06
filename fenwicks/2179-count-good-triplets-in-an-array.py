from typing import List
import unittest

# https://leetcode.com/problems/count-good-triplets-in-an-array/

# python3 -m unittest fenwicks/2179-count-good-triplets-in-an-array.py


class Solution(unittest.TestCase):
    # Approach: Fenwick Tree (Binary Indexed Tree)
    # Time: O(nlogn)
    # Space: O(n)
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        tree = [0] * (n + 2)  # 1-indexed

        def update(num: int):
            while num < len(tree):
                tree[num] += 1
                num += num & -num

        def query(num: int) -> int:
            total = 0
            while num > 0:
                total += tree[num]
                num -= num & -num
            return total

        indices2 = [-1] * n
        for idx, num in enumerate(nums2):
            indices2[num] = idx
        total = 0
        for idx1, num in enumerate(nums1):
            idx2 = indices2[num]
            left_common = query(idx2 + 1)
            right_common = (n - idx2 - 1) - (idx1 - left_common)
            total += left_common * right_common
            update(idx2 + 1)
        return total

    def test(self):
        for nums1, nums2, expected in [
            ([2, 0, 1, 3], [0, 1, 2, 3], 1),
            ([4, 0, 1, 3, 2], [4, 1, 0, 2, 3], 4),
        ]:
            output = self.goodTriplets(nums1, nums2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
