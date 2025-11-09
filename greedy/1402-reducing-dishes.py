from typing import List
import unittest

# https://leetcode.com/problems/reducing-dishes/

# python3 -m unittest greedy/1402-reducing-dishes.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def maxSatisfaction(self, satisfaction: List[int]) -> int:
    #     satisfaction.sort()
    #     max_sum = 0
    #     n = len(satisfaction)
    #     for start in range(n):
    #         cur_sum = 0
    #         time = 1
    #         for i in range(start, n):
    #             cur_sum += time * satisfaction[i]
    #             time += 1
    #         max_sum = max(max_sum, cur_sum)

    #     return max_sum

    # Approach #2: Greedy from the end
    # Time: O(n log n)
    # Space: O(1)
    def maxSatisfaction(self, satisfaction: List[int]) -> int:
        satisfaction.sort(reverse=True)
        total, presum = 0, 0
        for dish in satisfaction:
            presum += dish
            if presum < 0:
                break
            total += presum
        return total

    def test(self):
        for satisfaction, expected in [
            ([-1, -8, 0, 5, -9], 14),
            ([4, 3, 2], 20),
            ([-1, -4, -5], 0),
        ]:
            output = self.maxSatisfaction(satisfaction.copy())
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
