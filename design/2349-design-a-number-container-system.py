from sortedcontainers import SortedSet
from collections import defaultdict
import heapq
import unittest

# https://leetcode.com/problems/design-a-number-container-system/

# python3 -m unittest design/2349-design-a-number-container-system.py

# Approach: Heap (Priority Queue)
# Time: O(n log n)
# Space: O(n)
class NumberContainers:
    def __init__(self):
        self.idx2num = {} # index -> number
        self.num2idx = defaultdict(list) # {number: [indices]}

    def change(self, index: int, number: int) -> None:
        self.idx2num[index] = number
        heapq.heappush(self.num2idx[number], index)

    def find(self, number: int) -> int:
        while self.num2idx[number] and self.idx2num[self.num2idx[number][0]] != number:
            heapq.heappop(self.num2idx[number])
        return self.num2idx[number][0] if self.num2idx[number] else -1

# # Approach: Sorted Set
# # Time: O(n log n)
# # Space: O(n)
# class NumberContainers:
#     def __init__(self):
#         self.idx2num = {} # index -> number
#         self.num2idx = defaultdict(SortedSet) # {number: [indices]}
# 
#     def change(self, index: int, number: int) -> None:
#         if index in self.idx2num:
#             self.num2idx[self.idx2num[index]].remove(index)
#         self.idx2num[index] = number
#         self.num2idx[number].add(index)
# 
#     def find(self, number: int) -> int:
#         return self.num2idx[number][0] if self.num2idx[number] else -1

class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["NumberContainers","find","change","change","change","change","find","change","find"],
                [[],[10],[2,10],[1,10],[3,10],[5,10],[10],[1,20],[10]],
                [None,-1,None,None,None,None,1,None,2],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "NumberContainers":
                        container = NumberContainers()
                    case "change":
                        container.change(values[idx][0], values[idx][1])
                    case "find":
                        output = container.find(values[idx][0])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
