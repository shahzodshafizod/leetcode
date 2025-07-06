from typing import List
import unittest
import heapq

# https://leetcode.com/problems/zero-array-transformation-iii/

# python3 -m unittest greedy/3362-zero-array-transformation-iii.py


class Solution(unittest.TestCase):
    # Approach: Greedy Line Sweep
    # Time: O(n + mlogm); n=len(nums), m=len(queries)
    # Space: O(n+m)
    def maxRemoval(self, nums: List[int], queries: List[List[int]]) -> int:
        line = [0] * (len(nums) + 1)
        ops, max_heap = 0, []
        queries.sort()
        qid, n = 0, len(queries)
        for idx, num in enumerate(nums):
            ops += line[idx]
            while qid < n and queries[qid][0] == idx:
                heapq.heappush(max_heap, -queries[qid][1])
                qid += 1
            while ops < num and max_heap and -max_heap[0] >= idx:
                ops += 1
                line[-heapq.heappop(max_heap) + 1] -= 1
            if ops < num:
                return -1
        return len(max_heap)

    def test(self):
        for nums, queries, expected in [
            ([1, 2, 3, 4], [[0, 3]], -1),
            ([2, 0, 2], [[0, 2], [0, 2], [1, 1]], 1),
            ([0, 0, 1, 1, 0], [[3, 4], [0, 2], [2, 3]], 2),
            ([0, 0, 1, 1, 0, 0], [[2, 3], [0, 2], [3, 5]], 2),
            ([0, 0, 3], [[0, 2], [1, 1], [0, 0], [0, 0]], -1),
            ([1, 1, 1, 1], [[1, 3], [0, 2], [1, 3], [1, 2]], 2),
            ([1, 1, 1, 1], [[1, 3], [0, 2], [1, 3], [0, 3]], 3),
            ([0, 3], [[0, 1], [0, 0], [0, 1], [0, 1], [0, 0]], 2),
            ([1, 2], [[1, 1], [0, 0], [1, 1], [1, 1], [0, 1], [0, 0]], 4),
            ([0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0], [[0, 6], [7, 14], [6, 7]], 2),
        ]:
            output = self.maxRemoval(nums, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
