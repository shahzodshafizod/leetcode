import heapq
from typing import List
import unittest

# https://leetcode.com/problems/rank-transform-of-an-array/

# python3 -m unittest arrays/1331-rank-transform-of-an-array.py

class Solution(unittest.TestCase):
    # # Approach: Heap (Priority Queue)
    # # Time: O(NLogN)
    # # Space: O(N)
    # def arrayRankTransform(self, arr: List[int]) -> List[int]:
    #     min_heap = []
    #     for idx, num in enumerate(arr):
    #         heapq.heappush(min_heap, (num, idx))
    #     rank = 0
    #     prev = None
    #     while min_heap:
    #         num, idx = heapq.heappop(min_heap)
    #         if rank == 0 or num != prev:
    #             rank += 1
    #         arr[idx] = rank
    #         prev = num
    #     return arr

    # Approach: Set + Hash Table + Sorting
    # Time: O(NLogN)
    # Space: O(N)
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        ranks = {num:idx+1 for idx,num in enumerate(sorted(set(arr)))}
        return [ranks[num] for num in arr]

    def test(self) -> None:
        for arr, expected in [
            ([40,10,20,30], [4,1,2,3]),
            ([100,100,100], [1,1,1]),
            ([37,12,28,9,100,56,80,5,12], [5,3,4,2,8,6,7,1,3]),
            ([40,10,10,20,30], [4,1,1,2,3]),
            ([], []),
            ([-1000000000,-1000000000,-1000000000,1000000000,1000000000,1000000000], [1,1,1,2,2,2]),
            ([0,0,0,-1], [2,2,2,1]),
        ]:
            output = self.arrayRankTransform(arr)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
