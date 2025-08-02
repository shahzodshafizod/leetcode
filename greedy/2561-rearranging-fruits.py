from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/rearranging-fruits/

# python3 -m unittest greedy/2561-rearranging-fruits.py


class Solution(unittest.TestCase):
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        freq: Counter[int] = Counter()
        m = 10**9
        for cost1 in basket1:
            freq[cost1] += 1
            m = min(m, cost1)
        for cost2 in basket2:
            freq[cost2] -= 1
            m = min(m, cost2)
        merge: List[int] = []
        for cost, cnt in freq.items():
            if cnt & 1 != 0:
                return -1
            merge.extend([cost] * (abs(cnt) // 2))
        if not merge:
            return 0
        merge.sort()
        return sum(min(2 * m, cost) for cost in merge[: len(merge) // 2])

    def test(self):
        for basket1, basket2, expected in [
            ([4, 2, 2, 2], [1, 4, 1, 2], 1),
            ([2, 3, 4, 1], [3, 2, 5, 1], -1),
        ]:
            output = self.minCost(basket1, basket2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
