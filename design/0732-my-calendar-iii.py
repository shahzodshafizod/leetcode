from sortedcontainers import SortedList
import unittest

# https://leetcode.com/problems/my-calendar-iii/

# python3 -m unittest design/0732-my-calendar-iii.py

# Approach: Line Sweep
class MyCalendarThree:
    def __init__(self):
        self.count = SortedList()

    def book(self, startTime: int, endTime: int) -> int:
        self.count.add((startTime, 1))
        self.count.add((endTime, -1))
        overlap, max_overlap = 0, 0
        for _, cnt in self.count:
            overlap += cnt
            max_overlap = max(max_overlap, overlap)
        return max_overlap

class TestMyCalendarThree(unittest.TestCase):
    def test(self) -> None:
        for commands, values, expected in [
            (
                ["MyCalendarThree","book","book","book","book","book","book"],
                [[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]],
                [None,1,1,2,3,3,3],
            ),
            (
                ["MyCalendarThree","book"],
                [[],[24,40],[43,50],[27,43],[5,21],[30,40],[14,29],[3,19],[3,14],[25,39],[6,19]],
                [None,1,1,2,2,3,3,3,3,4,4],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "MyCalendarThree":
                        calendar = MyCalendarThree()
                    case "book":
                        startTime, endTime = values[idx]
                        output = calendar.book(startTime, endTime)
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your MyCalendarThree object will be instantiated and called as such:
# obj = MyCalendarThree()
# param_1 = obj.book(startTime,endTime)