from typing import List
import unittest

# https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/

# python3 -m unittest binarysearch/2616-minimize-the-maximum-difference-of-pairs.py

class Solution(unittest.TestCase):
    # Approach: Greedy + Binary Search
    # Time: O(nlogn + nlogm), n=len(nums), m=max(nums)
    # Space: O(n), In Python, the sort() method sorts a list using
    #   the Timsort algorithm which is a combination of Merge Sort
    #   and Insertion Sort and has a space complexity of O(n)
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums.sort()
        left, right = 0, nums[-1]
        n = len(nums)
        while left < right:
            mid = (left + right) // 2

            cnt, idx = 0, 1
            while idx < n and cnt < p:
                if nums[idx] - nums[idx-1] <= mid:
                    cnt += 1
                    idx += 1
                idx += 1

            if cnt == p:
                right = mid
            else:
                left = mid+1
        return left

    def test(self):
        for nums, p, expected in [
            ([10,1,2,7,1,3], 2, 1),
            ([4,2,1,2], 1, 0),
            ([8,9,1,5,4,3,6,4,3,7], 4, 1),
        ]:
            output = self.minimizeMax(nums, p)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
