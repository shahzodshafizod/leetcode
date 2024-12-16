import unittest

# https://leetcode.com/problems/maximum-frequency-stack/

# python3 -m unittest design/0895-maximum-frequency-stack.py

# Heap (Priority Queue) will not work,
# because it cannot keep track of order:
# "If there is a tie for the most frequent element,
# the element closest to the stack's top is removed and returned."

class FreqStack:

    def __init__(self):
        self.counts = {}
        self.buckets = {}
        self.max_cnt = 0

    def push(self, val: int) -> None:
        cnt = self.counts.get(val, 0) + 1
        self.counts[val] = cnt
        if cnt not in self.buckets:
            self.buckets[cnt] = []
        self.buckets[cnt].append(val)
        self.max_cnt = max(self.max_cnt, cnt)

    def pop(self) -> int:
        val = self.buckets[self.max_cnt].pop()
        self.counts[val] -= 1
        if not self.buckets[self.max_cnt]:
            self.max_cnt -= 1
        return val

class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
                [[],[5],[7],[5],[7],[4],[5],[],[],[],[]],
                [None,None,None,None,None,None,None,5,7,5,4],
            ),
            (
                ["FreqStack","push","push","push","push","push","push","pop","push","pop","push","pop","push","pop","push","pop","pop","pop","pop","pop","pop"],
                [[],[4],[0],[9],[3],[4],[2],[],[6],[],[1],[],[1],[],[4],[],[],[],[],[],[]],
                [None,None,None,None,None,None,None,4,None,6,None,1,None,1,None,4,2,3,9,0,4],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "FreqStack":
                        f = FreqStack()
                    case "push":
                        f.push(values[idx][0])
                    case "pop":
                        output = f.pop()
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()
