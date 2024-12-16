import heapq
from typing import List
import unittest

# https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/

# python3 -m unittest heaps/3264-final-array-state-after-k-multiplication-operations-i.py

class Solution(unittest.TestCase):
    # Approach: Heap (Priority Queue)
    # Time: O(k log n)
    # Space: O(n)
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        min_heap = [(num,idx) for idx,num in enumerate(nums)]
        heapq.heapify(min_heap)
        for _ in range(k):
            _, idx = min_heap[0]
            nums[idx] *= multiplier
            heapq.heapreplace(min_heap, (nums[idx], idx))
        return nums

    def test(self):
        for nums, k, multiplier, expected in [
            ([2,1,3,5,6], 5, 2, [8,4,6,5,6]),
            ([1,2], 3, 4, [16,8]),
            ([2,1], 1, 2, [2,2]),
        ]:
            output = self.getFinalState(nums, k, multiplier)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
