from typing import List
import unittest

# https://leetcode.com/problems/minimum-time-to-make-rope-colorful/

# python3 -m unittest greedy/1578-minimum-time-to-make-rope-colorful.py


class Solution(unittest.TestCase):
    # # Approach #1: Stack
    # # Time: O(n)
    # # Space: O(n)
    # def minCost(self, colors: str, neededTime: List[int]) -> int:
    #     stack: List[int] = []
    #     time = 0
    #     for i, c in enumerate(colors):
    #         if stack and colors[stack[-1]] == c:
    #             if neededTime[i] > neededTime[stack[-1]]:
    #                 time += neededTime[stack[-1]]
    #                 stack[-1] = i
    #             else:
    #                 time += neededTime[i]
    #         else:
    #             stack.append(i)
    #     return time

    # Approach #2: Greedy
    # Time: O(n)
    # Space: O(1)
    def minCost(self, colors: str, neededTime: List[int]) -> int:
        res = max_cost = 0
        for i in range(len(colors)):
            if i > 0 and colors[i] != colors[i - 1]:
                max_cost = 0
            res += min(max_cost, neededTime[i])
            max_cost = max(max_cost, neededTime[i])
        return res

    def test(self):
        for colors, neededTime, expected in [
            ("abaac", [1, 2, 3, 4, 5], 3),
            ("abc", [1, 2, 3], 0),
            ("aabaa", [1, 2, 3, 4, 1], 2),
        ]:
            output = self.minCost(colors, neededTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
