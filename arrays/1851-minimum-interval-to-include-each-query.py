import heapq
from typing import List
import unittest

# https://leetcode.com/problems/minimum-interval-to-include-each-query/

# python3 -m unittest arrays/1851-minimum-interval-to-include-each-query.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force (Time Limit Exceeded)
    # # Time: O(NxQ), N=len(intervals), Q=len(queries)
    # # Space: O(1)
    # def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:
    #     for idx, query in enumerate(queries):
    #         queries[idx] = -1
    #         for interval in intervals:
    #             left, right = interval
    #             if left <= query <= right:
    #                 length = right - left + 1
    #                 if queries[idx] == -1 or length < queries[idx]:
    #                     queries[idx] = length
    #     return queries

    # # Approach: Sweep Line + Monotonic Increasing Stack (Time Limit Exceeded)
    # # Time: O((N+Q)Log(N+Q)), N=len(intervals), Q=len(queries)
    # # Space: O(N+Q)
    # def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:
    #     INTERVAL, QUERY = 1, 2
    #     points = [] # (left,kind,right) or (query,kind,index)
    #     for interval in intervals:
    #         left, right = interval
    #         heapq.heappush(points, (left, INTERVAL, right))
    #     for idx, query in enumerate(queries):
    #         heapq.heappush(points, (query, QUERY, idx))
    #     stack = [] # (length,right)
    #     while points:
    #         point = heapq.heappop(points)
    #         if point[1] == INTERVAL:
    #             left, right = point[0], point[2]
    #             length = right - left + 1
    #             stack.append((length,right))
    #             idx = len(stack)-1
    #             while idx > 0 and stack[idx-1][0] < stack[idx][0]:
    #                 stack[idx], stack[idx-1] = stack[idx-1], stack[idx]
    #                 idx -= 1
    #         else:
    #             query, idx = point[0], point[2]
    #             while stack and stack[-1][1] < query:
    #                 stack.pop()
    #             queries[idx] = stack[-1][0] if stack else -1
    #     return queries

    # Approach: Heap (Priority Queue)
    # Time: O(NLogN + QLogQ), N=len(intervals), Q=len(queries)
    # Space: O(N+Q)
    def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:
        intervals.sort()
        qmap = {}
        idx, n = 0, len(intervals)
        inbound = [] # (length, right)
        for query in sorted(queries):
            while idx < n and intervals[idx][0] <= query:
                left, right = intervals[idx][0], intervals[idx][1]
                heapq.heappush(inbound, (right-left+1, right))
                idx += 1
            while inbound and inbound[0][1] < query:
                heapq.heappop(inbound)
            qmap[query] = inbound[0][0] if inbound else -1
        for idx, query in enumerate(queries):
            queries[idx] = qmap[query]
        return queries

    def test(self) -> None:
        for intervals, queries, expected in [
            ([[1,4],[2,4],[3,6],[4,4]], [2,3,4,5], [3,3,1,4]),
            ([[2,3],[2,5],[1,8],[20,25]], [2,19,5,22], [2,-1,4,6]),
        ]:
            output = self.minInterval(intervals, queries)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
