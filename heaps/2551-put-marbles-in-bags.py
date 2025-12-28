from typing import List
import unittest

# https://leetcode.com/problems/put-marbles-in-bags/

# python3 -m unittest heaps/2551-put-marbles-in-bags.py


class Solution(unittest.TestCase):
    # # Approach #1: Heap (Priority Queue)
    # # Time: O(nlogn)
    # # Space: O(2n) = O(n)
    # def putMarbles(self, weights: List[int], k: int) -> int:
    #     min_heap, max_heap = [], []
    #     for idx in range(len(weights)-1):
    #         sum = weights[idx]+weights[idx+1]
    #         heapq.heappush(min_heap, sum)
    #         heapq.heappush(max_heap, -sum)
    #     min, max = 0, 0
    #     for _ in range(k-1):
    #         min += heapq.heappop(min_heap)
    #         max -= heapq.heappop(max_heap)
    #     return max - min

    # # Approach #2: Sorting
    # # Time: O(nlogn)
    # # Space: O(n)
    # def putMarbles(self, weights: List[int], k: int) -> int:
    #     borders = sorted([weights[idx] + weights[idx+1] for idx in range(len(weights)-1)])
    #     return sum(borders[1-k:]) - sum(borders[:k-1])

    # Approach #2: Sorting (Space-Optimized)
    # Time: O(nlogn)
    # Space: O(1)
    def putMarbles(self, weights: List[int], k: int) -> int:
        n = len(weights)
        for idx in range(n - 1):
            weights[idx] += weights[idx + 1]
        weights = weights[: n - 1]
        weights.sort()
        return sum(weights[1 - k :]) - sum(weights[: k - 1])

    def test(self):
        for weights, k, expected in [
            ([1, 3, 5, 1], 2, 4),
            ([1, 3], 2, 0),
            (
                [54, 6, 34, 66, 63, 52, 39, 62, 46, 75, 28, 65, 18, 37, 18, 13, 33, 69, 19, 40, 13, 10, 43, 61, 72],
                4,
                289,
            ),
            ([1, 4, 2, 5, 2], 3, 3),
            ([1, 9, 8, 7, 1, 1, 1], 3, 28),
        ]:
            output = self.putMarbles(weights, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
