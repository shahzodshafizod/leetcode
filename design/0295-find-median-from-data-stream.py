import heapq
import unittest

# https://leetcode.com/problems/find-median-from-data-stream/

# python3 -m unittest design/0295-find-median-from-data-stream.py

class MedianFinder():
    def __init__(self):
        self.max_heap = [] # less that or equal to median
        self.min_heap = [] # greate that median
    
    def addNum(self, num: int) -> None:
        heapq.heappush(self.max_heap, -num)
        heapq.heappush(self.min_heap, -heapq.heappop(self.max_heap))
        if len(self.min_heap) > len(self.max_heap):
            heapq.heappush(self.max_heap, -heapq.heappop(self.min_heap))

    def findMedian(self) -> float:
        if len(self.max_heap) > len(self.min_heap):
            return float(-self.max_heap[0])
        return (self.min_heap[0]-self.max_heap[0])/2

class TestMedianFinder(unittest.TestCase):
    def test(self) -> None:
        for commands, values, expected in [
            (
                ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"],
                [[],[1],[2],[],[3],[]],
                [None, None, None, 1.5, None, 2.0],
            ),
            (
                ["MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"],
                [[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]],
                [None, None, 6.00000, None, 8.00000, None, 6.00000, None, 6.00000, None, 6.00000, None, 5.50000, None, 6.00000, None, 5.50000, None, 5.00000, None, 4.00000, None, 3.00000],
            ),
            (
                ["MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"],
                [[],[1],[],[2],[],[3],[],[4],[],[5],[],[6],[],[7],[],[8],[],[9],[]],
                [None, None, 1.0, None, 1.5, None, 2.0, None, 2.5, None, 3.0, None, 3.5, None, 4.0, None, 4.5, None, 5.0],
            ),
            (
                ["MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"],
                [[],[1],[],[1],[],[1],[],[1],[],[1],[],[1],[],[1],[],[1],[],[1],[]],
                [None, None, 1.0, None, 1.0, None, 1.0, None, 1.0, None, 1.0, None, 1.0, None, 1.0, None, 1.0, None, 1.0],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "MedianFinder":
                        median = MedianFinder()
                    case "addNum":
                        median.addNum(values[idx][0])
                    case "findMedian":
                        output = median.findMedian()
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
