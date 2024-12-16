from collections import defaultdict
from typing import List
import unittest
from sortedcontainers import SortedDict

# https://leetcode.com/problems/data-stream-as-disjoint-intervals/

# python3 -m unittest design/0352-data-stream-as-disjoint-intervals.py

"""
[1,2,3,4,6,8,9] -> [[1,4], [6,6], [8,9]]
[1,2,3,4,6,7,8,9] -> [[1,4],[6,9]]
[1,2,3,4,5,6,7,8,9] -> [[1,9]]
"""

# # Approach #1: Union-Find
# # Time: O(N^2)
# # Space: O(N)
# class SummaryRanges:
#     def __init__(self):
#         self.MAX = int(1e4)
#         self.parent = [None] * (self.MAX+2)
#     def addNum(self, value: int) -> None:
#         if value == 0 or self.parent[value-1] == None:
#             # 1.1. create a new interval
#             self.parent[value] = value
#         else:
#             # 1.2. connect with the left-side interval
#             self.parent[value] = self.parent[value-1]
#         # 2. connect with the right-side interval
#         if self.parent[value+1] != None:
#             self.parent[value+1] = self.parent[value]
#     def find(self, x: int) -> int:
#         if self.parent[x] != x:
#             self.parent[x] = self.find(self.parent[x])
#         return self.parent[x]
#     def getIntervals(self) -> List[List[int]]:
#         intervals = []
#         idx = self.MAX+1
#         while idx-1 >= 0:
#             idx -= 1
#             if self.parent[idx] == None:
#                 continue
#             end = idx
#             idx = self.find(idx)
#             intervals.insert(0, [idx, end])
#         return intervals

# # Approach #2: Binary Search
# # Time: O(NLogN)
# # Space: O(N)
# class SummaryRanges:
#     def __init__(self):
#         self.ranges = []
#     def addNum(self, value: int) -> None:
#         n = len(self.ranges)
#         left, right = 0, n-1
#         while left <= right:
#             mid = left + ((right-left)>>1)
#             start, end = self.ranges[mid]
#             # 1. check if in a valid range
#             if start <= value and value <= end:
#                 return
#             if start > value:
#                 right = mid-1
#             else:
#                 left = mid+1
#         prev, next = right, right+1
#         # 2. check if can be merged
#         merged = False
#         if prev >= 0 and self.ranges[prev][1]+1 == value:
#             self.ranges[prev][1] += 1
#             merged = True
#         if next < n and self.ranges[next][0]-1 == value:
#             self.ranges[next][0] -= 1
#             merged = True
#         if merged:
#             if prev >= 0 and next < n and self.ranges[prev][1] == self.ranges[next][0]:
#                 self.ranges[prev][1] = self.ranges[next][1]
#                 self.ranges = self.ranges[:next] + self.ranges[next+1:]
#             return
#         # 3. add a new range
#         self.ranges.insert(next, [value, value])
#     def getIntervals(self) -> List[List[int]]:
#         return self.ranges

# Approach #1: SortedDict
# Time: O(NLogN)
# Space: O(N)
class SummaryRanges:
    def __init__(self):
        self.treeMap = SortedDict()
    def addNum(self, value: int) -> None:
        self.treeMap[value] = True
    def getIntervals(self) -> List[List[int]]:
        res = []
        for n in self.treeMap:
            if res and res[-1][1] + 1 == n:
                res[-1][1] = n
            else:
                res.append([n, n])
        return res

class Test(unittest.TestCase):
    def test(self) -> None:
        for commands, values, expected in [
            (
                ["SummaryRanges","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals"],
                [[],[1],[],[3],[],[7],[],[2],[],[6],[]],
                [None,None,[[1,1]],None,[[1,1],[3,3]],None,[[1,1],[3,3],[7,7]],None,[[1,3],[7,7]],None,[[1,3],[6,7]]],
            ),
            (
                ["SummaryRanges","addNum","getIntervals","addNum","getIntervals"],
                [[],[1],[],[0],[]],
                [None,None,[[1,1]],None,[[0,1]]],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "SummaryRanges":
                        sg = SummaryRanges()
                    case "addNum":
                        sg.addNum(values[idx][0])
                    case "getIntervals":
                        output = sg.getIntervals()
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
