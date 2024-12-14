from math import sqrt
from typing import List
import heapq
import unittest

# https://leetcode.com/problems/take-gifts-from-the-richest-pile/

# python3 -m unittest heaps/2558-take-gifts-from-the-richest-pile.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(nxk)
    # # Space: O(1)
    # def pickGifts(self, gifts: List[int], k: int) -> int:
    #     n = len(gifts)
    #     for _ in range(k):
    #         maxid = 0
    #         for idx in range(n):
    #             if gifts[idx] > gifts[maxid]:
    #                 maxid = idx
    #         gifts[maxid] = int(sqrt(gifts[maxid]))
    #     return sum(gifts)

    # Approach: Heap (Priority Queue)
    # Time: O(n+klogn)
    # Space: O(n)
    def pickGifts(self, gifts: List[int], k: int) -> int:
        max_heap = [-gift for gift in gifts]
        heapq.heapify(max_heap)
        for _ in range(k):
            heapq.heapreplace(max_heap, -int(sqrt(-max_heap[0])))
        return -sum(max_heap)

    def test(self) -> None:
        for gifts, k, expected in [
            ([1,1,1,1], 4, 4),
            ([25,64,9,4,100], 4, 29),
            ([1,2,2,2,2,2,1], 1000, 7),
            ([11,41,66,63,14,39,70,58,12,71], 20, 17),
            # ([64,43,32,42,9,25,73,29,56,41,27,71,62,57,67,34,8,71,2,12,52,1,31], 51, 33),
            # ([12,34,9,29,10,39,45,56,24,8,65,48,15,5,3,25,24,16,62,27,8,3,70,55,13,60], 3, 591),
            # ([70,51,68,36,67,31,28,54,41,74,31,38,24,25,24,5,34,61,9,12,17,20,5,11,75], 18, 191),
            # ([8,21,21,44,68,33,16,57,23,2,61,53,73,66,40,46,50,33,20,72,2,59,11,43,6,70,13,51,26,34,46,61,73,22,27,36,18,31,62], 24, 396),
            # ([5,52,73,54,6,22,58,9,34,21,58,68,63,72,1,5,64,42,37,46,17,40,50,54,11,1,25,43,21,31,29,58,49,73,54,40,60,7,54,25], 37, 214),
        ]:
            output = self.pickGifts(gifts, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
