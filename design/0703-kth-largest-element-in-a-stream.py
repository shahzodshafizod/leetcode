import heapq
from typing import List
import unittest

# https://leetcode.com/problems/kth-largest-element-in-a-stream/

# python3 -m unittest design/0703-kth-largest-element-in-a-stream.py

# Time: O(LogK)TimeForEachAddOperation
# Space: O(K)SpaceToStoreKElementsInTheHeap
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.nums = []
        for num in nums:
            self.add(num)

    def add(self, val: int) -> int:
        if len(self.nums) < self.k:
            heapq.heappush(self.nums, val)
        elif val > self.nums[0]:
            heapq.heappushpop(self.nums, val)
        return self.nums[0]

# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)

class TestKthLargest(unittest.TestCase):
    def test(self) -> None:
        for commands, values, expected in [
            (
                ["KthLargest", "add", "add", "add", "add", "add"],
                [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]],
                [None, 4, 5, 5, 8, 8]
            ),
            (
                ["KthLargest", "add", "add", "add", "add", "add"],
                [[3, [5, -1]], [2], [1], [-1], [3], [4]],
                [None, -1, 1, 1, 2, 3],
            ),
            (
                ["KthLargest", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add"],
                [[7, [-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1]], [3], [2], [3], [1], [2], [4], [5], [5], [6], [7], [7], [8], [2], [3], [1], [1], [1], [10], [11], [5], [6], [2], [4], [7], [8], [5], [6]],
                [None, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7],
            ),
            (
                ["KthLargest", "add", "add", "add", "add", "add"],
                [[1, []], [-3], [-2], [-4], [0], [4]],
                [None, -3, -2, -2, 0, 4],
            )
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "KthLargest":
                        heap = KthLargest(values[idx][0], values[idx][1])
                    case "add":
                        output = heap.add(values[idx][0])
                self.assertEqual(output, expected[idx])
