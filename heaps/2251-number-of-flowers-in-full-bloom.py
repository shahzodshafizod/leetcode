import heapq
from typing import List
import unittest

# https://leetcode.com/problems/number-of-flowers-in-full-bloom/

# python3 -m unittest heaps/2251-number-of-flowers-in-full-bloom.py

class Solution(unittest.TestCase):
    # Approach: Heap (Priority Queue)
    # Time: O(nlogn + mlogm), n=len(flowers), m=len(people)
    # Space: O(n + m)
    def fullBloomFlowers(self, flowers: List[List[int]], people: List[int]) -> List[int]:
        flowers.sort()
        ends = []
        n, m = len(flowers), len(people)
        answer = [0] * m
        fid = 0
        for person, pid in sorted([(p,i) for i, p in enumerate(people)]):
            while fid < n and flowers[fid][0] <= person:
                heapq.heappush(ends, flowers[fid][1])
                fid += 1
            while ends and ends[0] < person:
                heapq.heappop(ends)
            answer[pid] = len(ends)
        return answer

    def test(self) -> None:
        for flowers, people, expected in [
            ([[1,10],[3,3]], [3,3,2], [2,2,1]),
            ([[1,6],[3,7],[9,12],[4,13]], [2,3,7,11], [1,2,2,2]),
        ]:
            output = self.fullBloomFlowers(flowers, people)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
