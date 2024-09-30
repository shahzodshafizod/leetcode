import heapq
from typing import List
import unittest

# https://leetcode.com/problems/maximum-performance-of-a-team/

# python3 -m unittest greedy/1383-maximum-performance-of-a-team.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(2^N)
    # def maxPerformance(self, n: int, speed: List[int], efficiency: List[int], k: int) -> int:
    #     def dp(idx: int, spd: int, eff: int, count: int) -> int:
    #         if idx == n or count == k:
    #             return spd * eff
    #         return max(
    #             # decision to include
    #             dp(idx+1, spd+speed[idx], min(eff, efficiency[idx]), count+1),
    #             # decision NOT to include
    #             dp(idx+1, spd, eff, count)
    #         )
    #     return dp(0, 0, 1e8, 0) % (10**9+7)

    # Time: O(n Log n)
    # Space: O(n)
    def maxPerformance(self, n: int, speed: List[int], efficiency: List[int], k: int) -> int:
        engineers = [[eff,spd] for eff,spd in zip(efficiency, speed)]
        engineers.sort(reverse=True)
        performance, total_speed = 0, 0
        min_heap = []
        for eff, spd in engineers:
            if len(min_heap) == k:
                total_speed -= heapq.heappop(min_heap)
            total_speed += spd
            heapq.heappush(min_heap, spd)
            performance = max(performance, total_speed * eff)
        return performance % (10**9+7)

    def test(self) -> None:
        for n, speed, efficiency, k, expected in [
            (6, [2,10,3,1,5,8], [5,4,3,9,7,2], 2, 60),
            (6, [2,10,3,1,5,8], [5,4,3,9,7,2], 3, 68),
            (6, [2,10,3,1,5,8], [5,4,3,9,7,2], 4, 72),
            (3, [2,8,2], [2,7,1], 2, 56),
        ]:
            output = self.maxPerformance(n, speed, efficiency, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
