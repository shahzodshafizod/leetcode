import unittest
from typing import List

# https://leetcode.com/problems/move-zeroes/

# python3 -m unittest twopointers/0283-move-zeroes.py


class Solution(unittest.TestCase):
    def moveZeroes(self, nums: List[int]) -> None:
        slow = 0
        for fast in range(len(nums)):
            if nums[fast] != 0:
                nums[slow], nums[fast] = nums[fast], nums[slow]
                slow += 1

    def testMoveZeroes(self) -> None:
        for nums, expected in [
            ([0], [0]),
            ([1, 0, 0], [1, 0, 0]),
            ([0, 1, 0, 3, 12], [1, 3, 12, 0, 0]),
            ([0, 0, 0, 0, 0, 0, 1], [1, 0, 0, 0, 0, 0, 0]),
            ([1, 2, 3, 0, 0, 0, 4, 5], [1, 2, 3, 4, 5, 0, 0, 0]),
            ([4, 2, 4, 0, 0, 3, 0, 5, 1, 0], [4, 2, 4, 3, 5, 1, 0, 0, 0, 0]),
            (
                [45192, 0, -659, -52359, -99225, -75991, 0, -15155, 27382, 59818, 0, -30645, -17025, 81209, 887, 64648],
                [45192, -659, -52359, -99225, -75991, -15155, 27382, 59818, -30645, -17025, 81209, 887, 64648, 0, 0, 0],
            ),
            # ([16601,78714,11653,-45162,0,-22859,0,36007,19143,-91750,0,-45361,-10715,46528,-91518,-36985,59578,76628,-87592,89803,0,-41430,44290,59831,41824,-30916,-6521,61614,46035,39346,0,0,32417], [16601,78714,11653,-45162,-22859,36007,19143,-91750,-45361,-10715,46528,-91518,-36985,59578,76628,-87592,89803,-41430,44290,59831,41824,-30916,-6521,61614,46035,39346,32417,0,0,0,0,0,0]),
        ]:
            self.moveZeroes(nums)
            self.assertEqual(expected, nums, f"expected: {expected}, moved: {nums}")
