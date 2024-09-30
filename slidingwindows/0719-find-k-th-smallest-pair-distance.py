from typing import List
import unittest

# python3 -m unittest slidingwindows/0719-find-k-th-smallest-pair-distance.py

class Solution(unittest.TestCase):
    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        def countPairDistances(distance: int) -> int:
            left, count = 0, 0
            for right in range(len(nums)):
                while nums[right] - nums[left] > distance:
                    left += 1
                count += right - left
            return count

        nums.sort()
        left, right = 0, nums[-1]
        while left < right:
            distance = left + (right - left) // 2
            if countPairDistances(distance) >= k:
                right = distance
            else:
                left = distance + 1
        return right

    def test(self) -> None:
        for nums, k, expected in [
            ([1,3,1], 1, 0),
            ([1,1,1], 2, 0),
            ([1,6,1], 3, 5),
        ]:
            output = self.smallestDistancePair(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
