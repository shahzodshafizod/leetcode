from typing import List
import unittest

# https://leetcode.com/problems/find-in-mountain-array/

# python3 -m unittest binarysearch/1095-find-in-mountain-array.py

# """
# This is MountainArray's API interface.
# You should not implement it, or speculate about its implementation
# """
# class MountainArray:
#    def get(self, index: int) -> int:
#    def length(self) -> int:


class MountainArray:
    def __init__(self, arr: List[int]):
        self.arr = arr

    def get(self, index: int) -> int:
        if index < 0 or index >= self.length():
            return -1
        return self.arr[index]

    def length(self) -> int:
        return len(self.arr)


class Solution(unittest.TestCase):
    def findInMountainArray(self, target: int, ma: MountainArray) -> int:
        left, right = 1, ma.length() - 2
        peak = 1
        while left <= right:
            peak = (left + right) // 2
            if ma.get(peak - 1) < ma.get(peak):
                left = peak + 1
            else:
                right = peak - 1
        left, right = 0, peak
        while left <= right:
            mid = (left + right) // 2
            val = ma.get(mid)
            if val < target:
                left = mid + 1
            elif val > target:
                right = mid - 1
            else:
                return mid
        left, right = peak, ma.length() - 1
        while left <= right:
            mid = (left + right) // 2
            val = ma.get(mid)
            if val < target:
                right = mid - 1
            elif val > target:
                left = mid + 1
            else:
                return mid
        return -1

    def test(self):
        for mountainArr, target, expected in [
            ([6, 7, 2], 2, 2),
            ([0, 1, 2, 4, 2, 1], 3, -1),
            ([1, 2, 3, 4, 5, 3, 1], 3, 2),
            ([7, 17, 24, 28, 30, 33, 26, 20, 12, 2], 2, 9),
            # ([11,21,29,38,39,40,47,49,50,54,57,60,61,69,78,81,83,92,102,104,109,114,123,132,135,143,144,153,163,166,167,175,173,165,155,145,136,126,118,110,104,103,100,92,87,77,71,65,62,59,57,52,49,41,40,34,30,28,24,15,7,0], 92, 17), # pylint: disable=line-too-long
        ]:
            mountainArr = MountainArray(mountainArr)
            output = self.findInMountainArray(target, mountainArr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
