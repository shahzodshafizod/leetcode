from typing import List
import unittest

# https://leetcode.com/problems/find-building-where-alice-and-bob-can-meet/

# python3 -m unittest stacks/monotonic/2940-find-building-where-alice-and-bob-can-meet.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(Q*H), Q=len(queries), H=len(heights)
    # # Space: O(1)
    # def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
    #     hlen = len(heights)
    #     ans = [-1] * len(queries)
    #     for idx, (a, b) in enumerate(queries):
    #         if a > b:
    #             a, b = b, a
    #         if a == b or heights[a] < heights[b]:
    #             ans[idx] = b
    #         else:
    #             for next in range (b+1, hlen):
    #                 if heights[next] > heights[a]:
    #                     ans[idx] = next
    #                     break
    #     return ans

    # Approach #2: Monotonic Decreasing Stack
    # Time: O(QlogQ + QlogH), Q=len(queries), H=len(heights)
    # Space: O(Q+H)
    def leftmostBuildingQueries(self, heights: List[int], queries: List[List[int]]) -> List[int]:
        qlen = len(queries)
        ans = [-1] * qlen
        remained = []
        for idx, (a, b) in enumerate(queries):
            if a > b:
                a, b = b, a
            if a == b or heights[a] < heights[b]:
                ans[idx] = b
            else:
                remained.append((b, idx, heights[a]))
        hid = len(heights) - 1
        stack = []
        for qid, aid, height in sorted(remained, reverse=True):  # O(q)
            while hid > qid:
                while stack and heights[stack[-1]] <= heights[hid]:
                    stack.pop()
                stack.append(hid)
                hid -= 1
            left, right = 0, len(stack) - 1
            while left <= right:  # O(log h)
                mid = (left + right) // 2
                if heights[stack[mid]] > height:
                    left = mid + 1
                else:
                    right = mid - 1
            if right != -1:
                ans[aid] = stack[right]
        return ans

    def test(self):
        for heights, queries, expected in [
            ([6, 4, 8, 5, 2, 7], [[0, 1], [0, 3], [2, 4], [3, 4], [2, 2]], [2, 5, -1, 5, 2]),
            ([5, 3, 8, 2, 6, 1, 4, 6], [[0, 7], [3, 5], [5, 2], [3, 0], [1, 6]], [7, 6, -1, 4, 6]),
        ]:
            output = self.leftmostBuildingQueries(heights, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
