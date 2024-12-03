from typing import List
import unittest

# https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/

# python3 -m unittest binarysearch/1964-find-the-longest-valid-obstacle-course-at-each-position.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def longestObstacleCourseAtEachPosition(self, obstacles: List[int]) -> List[int]:
    #     n = len(obstacles)
    #     ans = [0] * n
    #     for curr in range(n):
    #         prev = curr
    #         for j in range(curr-1,-1,-1):
    #             if obstacles[j] <= obstacles[curr] and ans[j] > ans[prev]:
    #                 prev = j
    #         if prev == curr:
    #             ans[curr] = 1
    #         else:
    #             ans[curr] = ans[prev]+1
    #     return ans

    # Approach: Binary Search + LIS (Longest Increasing (Non-Decreasing) Subsequence)
    # Time: O(n log n)
    # Space: O(n)
    def longestObstacleCourseAtEachPosition(self, obstacles: List[int]) -> List[int]:
        n = len(obstacles)
        lis, ans = [None] * n, [None] * n
        size = 0
        for idx in range(n):
            left, right = 0, size
            while left < right:
                mid = (left + right) // 2
                if lis[mid] > obstacles[idx]:
                    right = mid
                else:
                    left = mid + 1
            if right == size:
                lis[size] = obstacles[idx]
                size += 1
            else:
                lis[right] = obstacles[idx]
            ans[idx] = right+1
        return ans

    def test(self) -> None:
        for obstacles, expected in [
            ([1,2,3,2], [1,2,3,3]),
            ([2,2,1], [1,2,1]),
            ([3,1,5,6,4,2], [1,1,2,3,2,2]),
            ([5,1,5,5,1,3,4,5,1,4], [1,1,2,3,2,3,4,5,3,5]),
        ]:
            output = self.longestObstacleCourseAtEachPosition(obstacles)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
